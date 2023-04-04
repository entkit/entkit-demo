import React from "react";
import { createRoot } from "react-dom/client";

import App from "./app";
import * as environment from "./environment";
import Keycloak from "keycloak-js";
import { ReactKeycloakProvider } from "@react-keycloak/web";

const keycloak = new Keycloak(environment.keycloak);
const container = document.getElementById("root") as HTMLElement;
const root = createRoot(container);

root.render(
    <ReactKeycloakProvider authClient={keycloak}>
        <React.StrictMode>
            <App />
        </React.StrictMode>
    </ReactKeycloakProvider>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
import { ReportHandler } from "web-vitals";

const reportWebVitals = (onPerfEntry?: ReportHandler) => {
    if (onPerfEntry && onPerfEntry instanceof Function) {
        import("web-vitals").then(
            ({ getCLS, getFID, getFCP, getLCP, getTTFB }) => {
                getCLS(onPerfEntry);
                getFID(onPerfEntry);
                getFCP(onPerfEntry);
                getLCP(onPerfEntry);
                getTTFB(onPerfEntry);
            },
        );
    }
};
reportWebVitals();
