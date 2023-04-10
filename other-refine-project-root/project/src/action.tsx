
import React from "react";
import * as Antd from "antd";
import * as AntdIcons from "@ant-design/icons";
import {
    RefineButtonSingleProps,
    RefineButtonLinkingProps,
    RefineButtonCommonProps
} from "@refinedev/ui-types";
import {ButtonProps, Popconfirm} from "antd";
import {
    useCustomMutation,
    useNotification,
    useInvalidate,
    useLink,
} from "@refinedev/core";
import * as Custom from "./custom";
import * as Type from "./typedefs";
import { usePermissions } from "@refinedev/core";

export type CompanyListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CompanyListAction: React.FC<CompanyListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Company?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/com".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}

export type CompanyShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CompanyShowAction: React.FC<CompanyShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Company?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/com/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}

export type CompanyDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CompanyDeleteAction: React.FC<CompanyDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Company?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoCompanyInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "CompanyWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "company",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}

export type CompanyEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CompanyEditAction: React.FC<CompanyEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Company?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/com/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type CountryListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CountryListAction: React.FC<CountryListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Country?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/country".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type CountryShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CountryShowAction: React.FC<CountryShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Country?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/country/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type CountryDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CountryDeleteAction: React.FC<CountryDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Country?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoCountryInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "CountryWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "country",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type CountryEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const CountryEditAction: React.FC<CountryEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Country?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/country/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type EmailListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const EmailListAction: React.FC<EmailListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Email?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/email".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type EmailShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const EmailShowAction: React.FC<EmailShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Email?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/email/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type EmailDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const EmailDeleteAction: React.FC<EmailDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Email?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoEmailInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "EmailWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "email",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type EmailEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const EmailEditAction: React.FC<EmailEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Email?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/email/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type ImageListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ImageListAction: React.FC<ImageListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Image?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/image".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type ImageShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ImageShowAction: React.FC<ImageShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Image?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/image/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type ImageDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ImageDeleteAction: React.FC<ImageDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Image?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoImageInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "ImageWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "image",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type ImageEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ImageEditAction: React.FC<ImageEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Image?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/image/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type LocationListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const LocationListAction: React.FC<LocationListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Location?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/location".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type LocationShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const LocationShowAction: React.FC<LocationShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Location?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/location/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type LocationDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const LocationDeleteAction: React.FC<LocationDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Location?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoLocationInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "LocationWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "location",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type LocationEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const LocationEditAction: React.FC<LocationEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Location?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/location/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type PhoneListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const PhoneListAction: React.FC<PhoneListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Phone?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/phone".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type PhoneShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const PhoneShowAction: React.FC<PhoneShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Phone?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/phone/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type PhoneDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const PhoneDeleteAction: React.FC<PhoneDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Phone?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoPhoneInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "PhoneWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "phone",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type PhoneEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const PhoneEditAction: React.FC<PhoneEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Phone?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/phone/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type ProductListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ProductListAction: React.FC<ProductListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Product?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/product".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type ProductShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ProductShowAction: React.FC<ProductShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Product?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/product/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type ProductEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ProductEditAction: React.FC<ProductEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Product?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/product/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type ProductDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ProductDeleteAction: React.FC<ProductDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Product?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoProductInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "ProductWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "product",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type ProductMyCustomActionButtonActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const ProductMyCustomActionButtonAction: React.FC<ProductMyCustomActionButtonActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Product?.includes("MyCustomActionButton"));
    const additionalProps = null || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoProductInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="MyCustomActionButton"
        okText="MyCustomActionButton"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "MyCustomActionButton",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "ProductWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "product",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.FileOutlined/> } {...additionalProps} {...props} >
            {hideText || "MyCustomActionButton"}
        </Antd.Button>
    </Popconfirm> : null
}


export type VendorListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const VendorListAction: React.FC<VendorListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Vendor?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/vendor".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type VendorShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const VendorShowAction: React.FC<VendorShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Vendor?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/vendor/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type VendorDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const VendorDeleteAction: React.FC<VendorDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Vendor?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoVendorInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "VendorWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "vendor",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type VendorEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const VendorEditAction: React.FC<VendorEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Vendor?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/vendor/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type WarehouseListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WarehouseListAction: React.FC<WarehouseListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Warehouse?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/warehouse".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type WarehouseShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WarehouseShowAction: React.FC<WarehouseShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Warehouse?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/warehouse/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type WarehouseDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WarehouseDeleteAction: React.FC<WarehouseDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Warehouse?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoWarehouseInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "WarehouseWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "warehouse",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type WarehouseEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WarehouseEditAction: React.FC<WarehouseEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Warehouse?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/warehouse/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}


export type WebsiteListActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WebsiteListAction: React.FC<WebsiteListActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Website?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/website".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.UnorderedListOutlined/> } {...additionalProps} {...props} >
            {hideText || "List"}
        </Antd.Button>
    </Link> : null
}


export type WebsiteShowActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WebsiteShowAction: React.FC<WebsiteShowActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Website?.includes("Read"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/website/show/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EyeOutlined/> } {...additionalProps} {...props} >
            {hideText || "Show"}
        </Antd.Button>
    </Link> : null
}


export type WebsiteDeleteActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WebsiteDeleteAction: React.FC<WebsiteDeleteActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Website?.includes("Delete"));
    const additionalProps = {"danger":true} || {};

    const  notification = useNotification();
    const { mutate, isLoading } = useCustomMutation();
    //const { mutate, isLoading } = useCustomMutation<Type.DemoWebsiteInterface>();
    const invalidate = useInvalidate();

    return can ? <Popconfirm
        key="delete"
        okText="Delete"
        cancelText="Cancel"
        okType="primary"
        title="Are you sure?"
        okButtonProps={ { disabled: isLoading } }
        onConfirm={(): void => {
            mutate(
                {
                    method: "post",
                    url: "",
                    values: {},
                    metaData: {
                        operation: "delete",
                        variables: {
                            where: {
                                value: {
                                    idIn: recordItemIDs
                                },
                                type: "WebsiteWhereInput",
                                required: true
                            },
                        },
                        fields: null || undefined,
                    },
                },
                {
                    onSuccess: (resp) => {
                        recordItemIDs.forEach((id: Type.DemoID)=>{
                            invalidate({
                                resource: "website",
                                invalidates: ["resourceAll"],
                                id,
                            });
                        })
                        notification.open?.({
                            type: "success",
                            message: `Successfully`,
                        })
                        !onSuccess || onSuccess(resp);
                    },
                    onError: (error) => {
                        notification.open?.({
                            type: "error",
                            message: error.message,
                        })
                    },
                },
            );
        }}
    >
        <Antd.Button icon={ <AntdIcons.DeleteOutlined/> } {...additionalProps} {...props} >
            {hideText || "Delete"}
        </Antd.Button>
    </Popconfirm> : null
}


export type WebsiteEditActionProps = ButtonProps &
    RefineButtonCommonProps &
    RefineButtonSingleProps &
    RefineButtonLinkingProps & {
    recordItemIDs: Type.DemoID[],
    onSuccess?: (data: any)=>void
}

export const WebsiteEditAction: React.FC<WebsiteEditActionProps> = ({recordItemIDs, hideText, onSuccess, ...props}) => {
    const { data: permissions } = usePermissions<Record<string, string[]>>();
    const can = Boolean(permissions?.Website?.includes("Update"));
    const additionalProps = null || {};
    const Link = useLink();

    return can ? <Link to={ "/website/edit/:id".replace(":id", String(recordItemIDs[0])) }>
        <Antd.Button icon={ <AntdIcons.EditOutlined/> } {...additionalProps} {...props} >
            {hideText || "Edit"}
        </Antd.Button>
    </Link> : null
}