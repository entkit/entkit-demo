import { keycloakBackendClientId } from "./environment";
import Keycloak from "keycloak-js";

export function getPermissions(keycloak: Keycloak) {
    return new Promise((resolve, reject) => {
        keycloak
            .updateToken(30)
            .then(() => {
                fetch(
                    keycloak.authServerUrl +
                        "/realms/" +
                        keycloak.realm +
                        "/protocol/openid-connect/token",
                    {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/x-www-form-urlencoded",
                            Authorization: "Bearer " + keycloak.token,
                        },
                        body: `grant_type=urn:ietf:params:oauth:grant-type:uma-ticket&response_include_resource_name=true&response_mode=permissions&audience=${keycloakBackendClientId}`,
                    },
                )
                    .then((response) => {
                        if (!response.ok) {
                            throw new Error("Failed to get permissions");
                        }
                        return response.json();
                    })
                    .then((data) => {
                        const res: Record<string, string[]> = {};
                        data.forEach(
                            (p: { rsname: string; scopes: string[] }) => {
                                res[p.rsname.replace(/^QWE/, "")] =
                                    p.scopes.map((s) => s.replace(/^QWE/, ""));
                            },
                        );
                        return res;
                    })
                    .then((data) => resolve(data))
                    .catch((error) => reject(error));
            })
            .catch((error) => {
                reject("Failed to update token:" + error.message);
            });
    });
}
