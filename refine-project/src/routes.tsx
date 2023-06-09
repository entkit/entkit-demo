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

import React from "react";

// Components imports
import * as List from "./list";
import * as Show from "./show";
import * as Create from "./create";
import * as Edit from "./edit";
import * as EdgesDiagram from "./edges-diagram";

import { Login } from "./login";
import { Header } from "./header";
import { Authenticated } from "@refinedev/core";
import { ErrorComponent, ThemedLayoutV2, ThemedTitleV2 } from "@refinedev/antd";
import { Route, Routes, Outlet } from "react-router-dom";
import {
    NavigateToResource,
    CatchAllNavigate,
} from "@refinedev/react-router-v6";

export const RoutesBundle: React.FC = () => {
    const loginUrl = window.environment.appPath + "login";
    return (
        <Routes>
            <Route
                path={window.environment.appPath}
                element={
                    <Authenticated
                        fallback={<CatchAllNavigate to={loginUrl} />}
                    >
                        <ThemedLayoutV2
                            Header={Header}
                            Title={({ collapsed }) => (
                                <ThemedTitleV2
                                    // collapsed is a boolean value that indicates whether the <Sidebar> is collapsed or not
                                    collapsed={collapsed}
                                    // icon={collapsed ? <MySmallIcon /> : <MyLargeIcon />}
                                    text="EntKit"
                                />
                            )}
                        >
                            <Outlet />
                        </ThemedLayoutV2>
                    </Authenticated>
                }
            >
                <Route
                    index
                    element={<NavigateToResource resource="company" />}
                />

                <Route path="com">
                    <Route path="show/:id" element={<Show.CompanyMainShow />} />
                    <Route index path="" element={<List.CompanyList />} />
                    <Route path="create" element={<Create.CompanyCreate />} />
                    <Route path="edit/:id" element={<Edit.CompanyEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.CompanyEdgesDiagram />}
                    />
                </Route>

                <Route path="country">
                    <Route path="show/:id" element={<Show.CountryMainShow />} />
                    <Route index path="" element={<List.CountryList />} />
                    <Route path="create" element={<Create.CountryCreate />} />
                    <Route path="edit/:id" element={<Edit.CountryEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.CountryEdgesDiagram />}
                    />
                </Route>

                <Route path="email">
                    <Route path="show/:id" element={<Show.EmailMainShow />} />
                    <Route index path="" element={<List.EmailList />} />
                    <Route path="create" element={<Create.EmailCreate />} />
                    <Route path="edit/:id" element={<Edit.EmailEdit />} />
                </Route>

                <Route path="image">
                    <Route path="show/:id" element={<Show.ImageMainShow />} />
                    <Route index path="" element={<List.ImageList />} />
                    <Route path="create" element={<Create.ImageCreate />} />
                    <Route path="edit/:id" element={<Edit.ImageEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.ImageEdgesDiagram />}
                    />
                </Route>

                <Route path="location">
                    <Route
                        path="show/:id"
                        element={<Show.LocationMainShow />}
                    />
                    <Route index path="" element={<List.LocationList />} />
                    <Route path="create" element={<Create.LocationCreate />} />
                    <Route path="edit/:id" element={<Edit.LocationEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.LocationEdgesDiagram />}
                    />
                </Route>

                <Route path="phone">
                    <Route path="show/:id" element={<Show.PhoneMainShow />} />
                    <Route index path="" element={<List.PhoneList />} />
                    <Route path="create" element={<Create.PhoneCreate />} />
                    <Route path="edit/:id" element={<Edit.PhoneEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.PhoneEdgesDiagram />}
                    />
                </Route>

                <Route path="product">
                    <Route path="show/:id" element={<Show.ProductMainShow />} />
                    <Route index path="" element={<List.ProductList />} />
                    <Route path="create" element={<Create.ProductCreate />} />
                    <Route path="edit/:id" element={<Edit.ProductEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.ProductEdgesDiagram />}
                    />
                </Route>

                <Route path="vendor">
                    <Route path="show/:id" element={<Show.VendorMainShow />} />
                    <Route index path="" element={<List.VendorList />} />
                    <Route path="create" element={<Create.VendorCreate />} />
                    <Route path="edit/:id" element={<Edit.VendorEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.VendorEdgesDiagram />}
                    />
                </Route>

                <Route path="warehouse">
                    <Route
                        path="show/:id"
                        element={<Show.WarehouseMainShow />}
                    />
                    <Route index path="" element={<List.WarehouseList />} />
                    <Route path="create" element={<Create.WarehouseCreate />} />
                    <Route path="edit/:id" element={<Edit.WarehouseEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.WarehouseEdgesDiagram />}
                    />
                </Route>

                <Route path="website">
                    <Route path="show/:id" element={<Show.WebsiteMainShow />} />
                    <Route index path="" element={<List.WebsiteList />} />
                    <Route path="create" element={<Create.WebsiteCreate />} />
                    <Route path="edit/:id" element={<Edit.WebsiteEdit />} />
                    <Route
                        path="edges/:id"
                        element={<EdgesDiagram.WebsiteEdgesDiagram />}
                    />
                </Route>
            </Route>

            <Route
                path={window.environment.appPath}
                element={
                    <Authenticated fallback={<Outlet />}>
                        <NavigateToResource resource="company" />
                    </Authenticated>
                }
            >
                <Route path="login" element={<Login />} />
            </Route>

            <Route
                path={window.environment.appPath}
                element={
                    <Authenticated>
                        <ThemedLayoutV2>
                            <Outlet />
                        </ThemedLayoutV2>
                    </Authenticated>
                }
            >
                <Route path="*" element={<ErrorComponent />} />
            </Route>
        </Routes>
    );
};
