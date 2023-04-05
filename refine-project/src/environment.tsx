import { KeycloakConfig } from "keycloak-js";

export const meta = {};

export const keycloak: KeycloakConfig = {
    url: "http://localhost:8080",
    realm: "entkit-demo-2",
    clientId: "frontend",
};

export const keycloakBackendClientId = "backend";

export const graphqlUrl = "http://localhost/query";
