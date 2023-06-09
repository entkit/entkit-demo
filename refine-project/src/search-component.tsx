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

import { useState, useEffect } from "react";
import { Typography, Avatar, AutoComplete, Input } from "antd";
import { useList, useLink } from "@refinedev/core";
import * as Interfaces from "./typedefs";
import debounce from "lodash/debounce";
import { FileImageOutlined, SearchOutlined } from "@ant-design/icons";

interface IOptionGroup {
    value: string;
    label: string | React.ReactNode;
}

interface IOptions {
    label: string | React.ReactNode;
    options: IOptionGroup[];
}

const { Text } = Typography;

export const SearchComponent: React.FC = () => {
    const [options, setOptions] = useState<IOptions[]>([]);
    const [value, setValue] = useState<string>("");
    const Link = useLink();

    const renderTitle = (title: string) => (
        <Typography.Title>
            <Typography.Text style={{ fontSize: "16px" }}>
                {title}
            </Typography.Text>
        </Typography.Title>
    );

    const renderItem = (
        title: string,
        imageUrl: string | null,
        link: string,
    ) => ({
        value: title,
        key: link,
        label: (
            <Link
                key={title + link}
                to={link}
                style={{ display: "flex", alignItems: "center" }}
            >
                {imageUrl ? (
                    <Avatar
                        size={48}
                        src={imageUrl}
                        style={{ minWidth: "48px" }}
                    />
                ) : (
                    <FileImageOutlined style={{ fontSize: "48px" }} />
                )}
                <Text style={{ marginLeft: "16px" }}>{title}</Text>
            </Link>
        ),
    });
    const { refetch: refetchCompany } =
        useList<Interfaces.DemoCompanyInterface>({
            resource: "company",
            metaData: {
                fields: ["id", "name"],
                searchQuery: value,
            },
            queryOptions: {
                enabled: false,
                onSuccess: (data) => {
                    const storeOptionGroup = data.data.map((item) =>
                        renderItem(
                            String(item.name),
                            null,
                            window.environment.appPath +
                                "com/show/:id".replace(":id", String(item.id)),
                        ),
                    );
                    if (storeOptionGroup.length > 0) {
                        setOptions((prevOptions) => [
                            ...prevOptions,
                            {
                                label: renderTitle("Company"),
                                options: storeOptionGroup,
                            },
                        ]);
                    }
                },
            },
        });
    const { refetch: refetchCountry } =
        useList<Interfaces.DemoCountryInterface>({
            resource: "country",
            metaData: {
                fields: ["id", "name"],
                searchQuery: value,
            },
            queryOptions: {
                enabled: false,
                onSuccess: (data) => {
                    const storeOptionGroup = data.data.map((item) =>
                        renderItem(
                            String(item.name),
                            null,
                            window.environment.appPath +
                                "country/show/:id".replace(
                                    ":id",
                                    String(item.id),
                                ),
                        ),
                    );
                    if (storeOptionGroup.length > 0) {
                        setOptions((prevOptions) => [
                            ...prevOptions,
                            {
                                label: renderTitle("Country"),
                                options: storeOptionGroup,
                            },
                        ]);
                    }
                },
            },
        });
    const { refetch: refetchEmail } = useList<Interfaces.DemoEmailInterface>({
        resource: "email",
        metaData: {
            fields: ["id", "title"],
            searchQuery: value,
        },
        queryOptions: {
            enabled: false,
            onSuccess: (data) => {
                const storeOptionGroup = data.data.map((item) =>
                    renderItem(
                        String(item.title),
                        null,
                        window.environment.appPath +
                            "email/show/:id".replace(":id", String(item.id)),
                    ),
                );
                if (storeOptionGroup.length > 0) {
                    setOptions((prevOptions) => [
                        ...prevOptions,
                        {
                            label: renderTitle("Email"),
                            options: storeOptionGroup,
                        },
                    ]);
                }
            },
        },
    });
    const { refetch: refetchImage } = useList<Interfaces.DemoImageInterface>({
        resource: "image",
        metaData: {
            fields: ["id", "title", "originalURL"],
            searchQuery: value,
        },
        queryOptions: {
            enabled: false,
            onSuccess: (data) => {
                const storeOptionGroup = data.data.map((item) =>
                    renderItem(
                        String(item.title),
                        `${item.originalURL}`,
                        window.environment.appPath +
                            "image/show/:id".replace(":id", String(item.id)),
                    ),
                );
                if (storeOptionGroup.length > 0) {
                    setOptions((prevOptions) => [
                        ...prevOptions,
                        {
                            label: renderTitle("Image"),
                            options: storeOptionGroup,
                        },
                    ]);
                }
            },
        },
    });
    const { refetch: refetchLocation } =
        useList<Interfaces.DemoLocationInterface>({
            resource: "location",
            metaData: {
                fields: ["id", "title"],
                searchQuery: value,
            },
            queryOptions: {
                enabled: false,
                onSuccess: (data) => {
                    const storeOptionGroup = data.data.map((item) =>
                        renderItem(
                            String(item.title),
                            null,
                            window.environment.appPath +
                                "location/show/:id".replace(
                                    ":id",
                                    String(item.id),
                                ),
                        ),
                    );
                    if (storeOptionGroup.length > 0) {
                        setOptions((prevOptions) => [
                            ...prevOptions,
                            {
                                label: renderTitle("Location"),
                                options: storeOptionGroup,
                            },
                        ]);
                    }
                },
            },
        });
    const { refetch: refetchPhone } = useList<Interfaces.DemoPhoneInterface>({
        resource: "phone",
        metaData: {
            fields: ["id", "title"],
            searchQuery: value,
        },
        queryOptions: {
            enabled: false,
            onSuccess: (data) => {
                const storeOptionGroup = data.data.map((item) =>
                    renderItem(
                        String(item.title),
                        null,
                        window.environment.appPath +
                            "phone/show/:id".replace(":id", String(item.id)),
                    ),
                );
                if (storeOptionGroup.length > 0) {
                    setOptions((prevOptions) => [
                        ...prevOptions,
                        {
                            label: renderTitle("Phone"),
                            options: storeOptionGroup,
                        },
                    ]);
                }
            },
        },
    });
    const { refetch: refetchProduct } =
        useList<Interfaces.DemoProductInterface>({
            resource: "product",
            metaData: {
                fields: ["id", "name", "image"],
                searchQuery: value,
            },
            queryOptions: {
                enabled: false,
                onSuccess: (data) => {
                    const storeOptionGroup = data.data.map((item) =>
                        renderItem(
                            String(item.name),
                            `${item.image}`,
                            window.environment.appPath +
                                "product/show/:id".replace(
                                    ":id",
                                    String(item.id),
                                ),
                        ),
                    );
                    if (storeOptionGroup.length > 0) {
                        setOptions((prevOptions) => [
                            ...prevOptions,
                            {
                                label: renderTitle("Product"),
                                options: storeOptionGroup,
                            },
                        ]);
                    }
                },
            },
        });
    const { refetch: refetchVendor } = useList<Interfaces.DemoVendorInterface>({
        resource: "vendor",
        metaData: {
            fields: ["id", "name"],
            searchQuery: value,
        },
        queryOptions: {
            enabled: false,
            onSuccess: (data) => {
                const storeOptionGroup = data.data.map((item) =>
                    renderItem(
                        String(item.name),
                        null,
                        window.environment.appPath +
                            "vendor/show/:id".replace(":id", String(item.id)),
                    ),
                );
                if (storeOptionGroup.length > 0) {
                    setOptions((prevOptions) => [
                        ...prevOptions,
                        {
                            label: renderTitle("Vendor"),
                            options: storeOptionGroup,
                        },
                    ]);
                }
            },
        },
    });
    const { refetch: refetchWarehouse } =
        useList<Interfaces.DemoWarehouseInterface>({
            resource: "warehouse",
            metaData: {
                fields: ["id", "name"],
                searchQuery: value,
            },
            queryOptions: {
                enabled: false,
                onSuccess: (data) => {
                    const storeOptionGroup = data.data.map((item) =>
                        renderItem(
                            String(item.name),
                            null,
                            window.environment.appPath +
                                "warehouse/show/:id".replace(
                                    ":id",
                                    String(item.id),
                                ),
                        ),
                    );
                    if (storeOptionGroup.length > 0) {
                        setOptions((prevOptions) => [
                            ...prevOptions,
                            {
                                label: renderTitle("Warehouse"),
                                options: storeOptionGroup,
                            },
                        ]);
                    }
                },
            },
        });
    const { refetch: refetchWebsite } =
        useList<Interfaces.DemoWebsiteInterface>({
            resource: "website",
            metaData: {
                fields: ["id", "title"],
                searchQuery: value,
            },
            queryOptions: {
                enabled: false,
                onSuccess: (data) => {
                    const storeOptionGroup = data.data.map((item) =>
                        renderItem(
                            String(item.title),
                            null,
                            window.environment.appPath +
                                "website/show/:id".replace(
                                    ":id",
                                    String(item.id),
                                ),
                        ),
                    );
                    if (storeOptionGroup.length > 0) {
                        setOptions((prevOptions) => [
                            ...prevOptions,
                            {
                                label: renderTitle("Website"),
                                options: storeOptionGroup,
                            },
                        ]);
                    }
                },
            },
        });

    useEffect(() => {
        setOptions([]);
        if (value.length < 3) {
            return;
        }
        refetchCompany();
        refetchCountry();
        refetchEmail();
        refetchImage();
        refetchLocation();
        refetchPhone();
        refetchProduct();
        refetchVendor();
        refetchWarehouse();
        refetchWebsite();
    }, [value]);

    return (
        <AutoComplete
            style={{
                width: "100%",
                maxWidth: "550px",
            }}
            options={options}
            filterOption={false}
            onSearch={debounce((value: string) => setValue(value), 300)}
        >
            <Input
                size="large"
                placeholder="Search"
                suffix={<SearchOutlined />}
            />
        </AutoComplete>
    );
};
