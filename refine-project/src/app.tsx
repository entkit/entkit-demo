import React, { useState } from "react";
import { AuthBindings, Authenticated, Refine } from "@refinedev/core";
import { ErrorComponent, Layout, notificationProvider } from "@refinedev/antd";
import "@refinedev/antd/dist/reset.css";
import { GraphQLClient } from "graphql-request";
import dataProvider from "./data-provider";
import routerProvider, {
    NavigateToResource,
    CatchAllNavigate,
    UnsavedChangesNotifier,
} from "@refinedev/react-router-v6";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { RoutesBundle } from "./routes";
import * as environment from "./environment";
import * as AntdIcons from "@ant-design/icons";
import { useKeycloak } from "@react-keycloak/web";
import { getPermissions } from "./auth";
import { usePermissions } from "@refinedev/core";

const client = new GraphQLClient(environment.graphqlUrl);

function App() {
    const [permissions, setPermissions] = useState<any | undefined>();
    const { keycloak, initialized } = useKeycloak();
    if (!initialized) {
        return <div>Loading...</div>;
    }
    const authProvider: AuthBindings = {
        login: async () => {
            const urlSearchParams = new URLSearchParams(window.location.search);
            const { to } = Object.fromEntries(urlSearchParams.entries());
            await keycloak.login({
                redirectUri: `${window.location.origin}${to || ""}`,
            });

            return {
                success: false,
                error: new Error("Login failed"),
            };
        },
        logout: async () => {
            try {
                await keycloak.logout({
                    redirectUri: window.location.origin,
                });
                return {
                    success: true,
                    redirectTo: "/login",
                };
            } catch (error) {
                return {
                    success: false,
                    error: new Error("Logout failed"),
                };
            }
        },
        onError: async (error) => {
            console.error(error);
            return { error };
        },
        check: async () => {
            try {
                const { token } = keycloak;
                if (token) {
                    client.setHeaders({ Authorization: `Bearer ${token}` });

                    await new Promise((resolve, reject) => {
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
                                            "Content-Type":
                                                "application/x-www-form-urlencoded",
                                            Authorization:
                                                "Bearer " + keycloak.token,
                                        },
                                        body: `grant_type=urn:ietf:params:oauth:grant-type:uma-ticket&response_include_resource_name=true&response_mode=permissions&audience=${environment.keycloakBackendClientId}`,
                                    },
                                )
                                    .then((response) => {
                                        if (!response.ok) {
                                            throw new Error(
                                                "Failed to get permissions",
                                            );
                                        }
                                        return response.json();
                                    })
                                    .then((data) => {
                                        const res: Record<string, string[]> =
                                            {};
                                        data.forEach(
                                            (p: {
                                                rsname: string;
                                                scopes: string[];
                                            }) => {
                                                res[
                                                    p.rsname.replace(/^QWE/, "")
                                                ] = p.scopes.map((s) =>
                                                    s.replace(/^QWE/, ""),
                                                );
                                            },
                                        );
                                        setPermissions(res);
                                    })
                                    .catch((error) => reject(error));
                            })
                            .catch((error) =>
                                reject(
                                    "Failed to update token:" + error.message,
                                ),
                            );
                    });

                    return {
                        authenticated: true,
                    };
                } else {
                    return {
                        authenticated: false,
                        logout: true,
                        redirectTo: "/login",
                        error: new Error("Token not found"),
                    };
                }
            } catch (error) {
                return {
                    authenticated: false,
                    logout: true,
                    redirectTo: "/login",
                    error: new Error("Token not found"),
                };
            }
        },
        getPermissions: async () => getPermissions(keycloak),
        getIdentity: async () => {
            if (keycloak?.tokenParsed) {
                return {
                    name: keycloak.tokenParsed.family_name,
                };
            }
            return null;
        },
    };

    return (
        <BrowserRouter>
            <Refine
                authProvider={authProvider}
                routerProvider={routerProvider}
                dataProvider={dataProvider(client)}
                notificationProvider={notificationProvider}
                catchAll={<ErrorComponent />}
                resources={[
                    {
                        name: "company",
                        list: "/company",
                        show: "/company/show/:id",
                        edit: "/company/edit/:id",
                        meta: {
                            icon: <AntdIcons.ShopOutlined />,
                            hide: !permissions?.Company?.includes("Read"),
                        },
                    },
                    {
                        name: "country",
                        list: "/country",
                        show: "/country/show/:id",
                        edit: "/country/edit/:id",
                        meta: {
                            icon: <AntdIcons.GlobalOutlined />,
                            hide: !permissions?.Country?.includes("Read"),
                        },
                    },
                    {
                        name: "email",
                        list: "/email",
                        show: "/email/show/:id",
                        edit: "/email/edit/:id",
                        meta: {
                            icon: <AntdIcons.MailOutlined />,
                            hide: !permissions?.Email?.includes("Read"),
                        },
                    },
                    {
                        name: "image",
                        list: "/image",
                        show: "/image/show/:id",
                        edit: "/image/edit/:id",
                        meta: {
                            icon: <AntdIcons.CameraOutlined />,
                            hide: !permissions?.Image?.includes("Read"),
                        },
                    },
                    {
                        name: "location",
                        list: "/location",
                        show: "/location/show/:id",
                        edit: "/location/edit/:id",
                        meta: {
                            icon: <AntdIcons.PushpinOutlined />,
                            hide: !permissions?.Location?.includes("Read"),
                        },
                    },
                    {
                        name: "phone",
                        list: "/phone",
                        show: "/phone/show/:id",
                        edit: "/phone/edit/:id",
                        meta: {
                            icon: <AntdIcons.PhoneOutlined />,
                            hide: !permissions?.Phone?.includes("Read"),
                        },
                    },
                    {
                        name: "product",
                        list: "/product",
                        show: "/product/show/:id",
                        edit: "/product/edit/:id",
                        meta: {
                            icon: <AntdIcons.FileOutlined />,
                            hide: !permissions?.Product?.includes("Read"),
                        },
                    },
                    {
                        name: "vendor",
                        list: "/vendor",
                        show: "/vendor/show/:id",
                        edit: "/vendor/edit/:id",
                        meta: {
                            icon: <AntdIcons.StarOutlined />,
                            hide: !permissions?.Vendor?.includes("Read"),
                        },
                    },
                    {
                        name: "warehouse",
                        list: "/warehouse",
                        show: "/warehouse/show/:id",
                        edit: "/warehouse/edit/:id",
                        meta: {
                            icon: <AntdIcons.OrderedListOutlined />,
                            hide: !permissions?.Warehouse?.includes("Read"),
                        },
                    },
                    {
                        name: "website",
                        list: "/website",
                        show: "/website/show/:id",
                        edit: "/website/edit/:id",
                        meta: {
                            icon: <AntdIcons.LinkOutlined />,
                            hide: !permissions?.Website?.includes("Read"),
                        },
                    },
                ]}
            >
                <RoutesBundle />
                <UnsavedChangesNotifier />
            </Refine>
        </BrowserRouter>
    );
}

export default App;
