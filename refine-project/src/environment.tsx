import { KeycloakConfig } from "keycloak-js";

export const meta = {};

export const keycloak: KeycloakConfig = {
    url: "http://localhost:8080",
    realm: "entkit",
    clientId: "xcontain-frontend",
};

export const keycloakBackendClientId = "xcontain-backend";

export const graphqlUrl = "http://localhost/query";
