// Code generated by EntKit. DO NOT EDIT.
// ---------------------------------------------------------
//
// Copyright (C) 2023 EntKit. All Rights Reserved.
//
// This code is part of the EntKit library and is generated
// automatically to ensure optimal functionality and maintainability.
// Any changes made directly to this file may be overwritten
// by future code generation, leading to unexpected behavior.
//
// Please refer to the EntKit documentation for instructions on
// how to modify the library, extend its functionality or contribute
// to the project: https://entkit.com
// ---------------------------------------------------------
import React, {useState} from "react";
import {AuthBindings,Authenticated, Refine} from "@refinedev/core";
import {
    ErrorComponent,
    notificationProvider,
    RefineThemes,
} from "@refinedev/antd";
import { ConfigProvider } from "antd";
import "@refinedev/antd/dist/reset.css";
import {GraphQLClient} from "graphql-request";
import dataProvider from "./data-provider";
import routerProvider, {
    NavigateToResource,
    CatchAllNavigate,
    UnsavedChangesNotifier,
} from "@refinedev/react-router-v6";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { RoutesBundle } from "./routes";
import * as AntdIcons from "@ant-design/icons";
import {useKeycloak} from "@react-keycloak/web";

function App() {
    const client = new GraphQLClient(window.environment.graphqlUrl);
    const [permissions, setPermissions] = useState<any|undefined>();
    const { keycloak, initialized } = useKeycloak();
    if (!initialized) {
        return <div>Loading...</div>;
    }
    const loginUrl = window.environment.appPath + "login";
    const appUrl = window.location.origin + window.environment.appPath;

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
                    redirectUri: appUrl,
                });
                return {
                    success: true,
                    redirectTo: loginUrl,
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
                    client.setHeaders({Authorization: `Bearer ${token}`})
                    return {
                        authenticated: true,
                    };
                } else {
                    return {
                        authenticated: false,
                        logout: true,
                        redirectTo: loginUrl,
                        error: new Error("Token not found"),
                    };
                }
            } catch (error) {
                return {
                    authenticated: false,
                    logout: true,
                    redirectTo: loginUrl,
                    error: new Error("Token not found"),
                };
            }
        },
        getPermissions: async ()=> {
            return fetch(
                `${keycloak.authServerUrl}/realms/${keycloak.realm}/protocol/openid-connect/token`,
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                        'Authorization': `Bearer ${keycloak.token}`,
                    },
                    body: `grant_type=urn:ietf:params:oauth:grant-type:uma-ticket&response_include_resource_name=true&response_mode=permissions&audience=${window.environment.auth.keycloak.backendClientId}`
                })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Failed to get permissions');
                    }
                    return response.json();
                })
                .then(data=>{
                    const res: Record<string, string[]> = {}
                    data.forEach((p: { rsname: string, scopes: string[] })=>{
                        res[p.rsname.replace(/^Demo/, "")] = p.scopes.map(s=>s.replace(/^Demo/, ""));
                    })
                    setPermissions(res)
                    return res
                })
                .catch(error => console.log(error));
        },
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
            <ConfigProvider theme={RefineThemes.Blue}>
                <Refine
                    authProvider={authProvider}
                    routerProvider={routerProvider}
                    dataProvider={dataProvider(client)}
                    notificationProvider={notificationProvider}
                    catchAll={<ErrorComponent/>}
                    resources={
                        [
                            {
                                name: "company",
                                list: window.environment.appPath + "com",
                                show: window.environment.appPath + "com/show/:id",
                                create: window.environment.appPath + "com/create",
                                edit: window.environment.appPath + "com/edit/:id",
                                meta: {
                                    icon: <AntdIcons.ShopOutlined/>,
                                    hide: !permissions?.Company?.includes("Read")
                                }
                            },
                            {
                                name: "country",
                                list: window.environment.appPath + "country",
                                show: window.environment.appPath + "country/show/:id",
                                create: window.environment.appPath + "country/create",
                                edit: window.environment.appPath + "country/edit/:id",
                                meta: {
                                    icon: <AntdIcons.GlobalOutlined/>,
                                    hide: !permissions?.Country?.includes("Read")
                                }
                            },
                            {
                                name: "email",
                                list: window.environment.appPath + "email",
                                show: window.environment.appPath + "email/show/:id",
                                create: window.environment.appPath + "email/create",
                                edit: window.environment.appPath + "email/edit/:id",
                                meta: {
                                    icon: <AntdIcons.MailOutlined/>,
                                    hide: !permissions?.Email?.includes("Read")
                                }
                            },
                            {
                                name: "image",
                                list: window.environment.appPath + "image",
                                show: window.environment.appPath + "image/show/:id",
                                create: window.environment.appPath + "image/create",
                                edit: window.environment.appPath + "image/edit/:id",
                                meta: {
                                    icon: <AntdIcons.CameraOutlined/>,
                                    hide: !permissions?.Image?.includes("Read")
                                }
                            },
                            {
                                name: "location",
                                list: window.environment.appPath + "location",
                                show: window.environment.appPath + "location/show/:id",
                                create: window.environment.appPath + "location/create",
                                edit: window.environment.appPath + "location/edit/:id",
                                meta: {
                                    icon: <AntdIcons.PushpinOutlined/>,
                                    hide: !permissions?.Location?.includes("Read")
                                }
                            },
                            {
                                name: "phone",
                                list: window.environment.appPath + "phone",
                                show: window.environment.appPath + "phone/show/:id",
                                create: window.environment.appPath + "phone/create",
                                edit: window.environment.appPath + "phone/edit/:id",
                                meta: {
                                    icon: <AntdIcons.PhoneOutlined/>,
                                    hide: !permissions?.Phone?.includes("Read")
                                }
                            },
                            {
                                name: "product",
                                list: window.environment.appPath + "product",
                                show: window.environment.appPath + "product/show/:id",
                                create: window.environment.appPath + "product/create",
                                edit: window.environment.appPath + "product/edit/:id",
                                meta: {
                                    icon: <AntdIcons.FileOutlined/>,
                                    hide: !permissions?.Product?.includes("Read")
                                }
                            },
                            {
                                name: "vendor",
                                list: window.environment.appPath + "vendor",
                                show: window.environment.appPath + "vendor/show/:id",
                                create: window.environment.appPath + "vendor/create",
                                edit: window.environment.appPath + "vendor/edit/:id",
                                meta: {
                                    icon: <AntdIcons.StarOutlined/>,
                                    hide: !permissions?.Vendor?.includes("Read")
                                }
                            },
                            {
                                name: "warehouse",
                                list: window.environment.appPath + "warehouse",
                                show: window.environment.appPath + "warehouse/show/:id",
                                create: window.environment.appPath + "warehouse/create",
                                edit: window.environment.appPath + "warehouse/edit/:id",
                                meta: {
                                    icon: <AntdIcons.OrderedListOutlined/>,
                                    hide: !permissions?.Warehouse?.includes("Read")
                                }
                            },
                            {
                                name: "website",
                                list: window.environment.appPath + "website",
                                show: window.environment.appPath + "website/show/:id",
                                create: window.environment.appPath + "website/create",
                                edit: window.environment.appPath + "website/edit/:id",
                                meta: {
                                    icon: <AntdIcons.LinkOutlined/>,
                                    hide: !permissions?.Website?.includes("Read")
                                }
                            },
                        ]
                    }
                >
                    <RoutesBundle/>
                    <UnsavedChangesNotifier />
                </Refine>
            </ConfigProvider>
        </BrowserRouter>
    );
}

export default App;