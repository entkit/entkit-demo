import React from "react";
import * as RA from "@refinedev/antd";
import * as Antd from "antd";
import * as AntdIcons from "@ant-design/icons";
import { useLink } from "@refinedev/core";
import * as Show from "./show";
import * as Type from "./typedefs";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import CodeEditor from "@uiw/react-textarea-code-editor";

export type ViewProps<T> = Omit<React.HTMLProps<HTMLElement>, "value"> & {
    value: T | undefined;
};

// region Date
export const DemoDateViewOnShow: React.FC<ViewProps<Date>> = ({ value }) => {
    return value ? (
        <RA.DateField format="LLL" value={value} />
    ) : (
        <RA.TextField value="Never" />
    );
};
export const DemoDateViewOnList: React.FC<ViewProps<Date>> = DemoDateViewOnShow;
export const DemoDateViewOnForm: React.FC = (props) => {
    return <Antd.DatePicker {...props} showTime={true} />;
};
// endregion Date

// region Boolean
export const DemoBooleanViewOnShow: React.FC<ViewProps<Type.DemoBoolean>> = ({
    value,
}) => {
    return <RA.TextField value={value ? "Yes" : "No"} />;
};
export const DemoBooleanViewOnList = DemoBooleanViewOnShow;
export const DemoBooleanViewOnForm: React.FC = (props) => {
    return (
        <Antd.Radio.Group {...props}>
            <Antd.Radio value={true}>Yes</Antd.Radio>
            <Antd.Radio value={false}>No</Antd.Radio>
        </Antd.Radio.Group>
    );
};
// endregion Boolean

// region String
export const DemoStringViewOnShow: React.FC<ViewProps<Type.DemoString>> = ({
    value,
}) => {
    return <Antd.Typography.Text copyable={true}>{value}</Antd.Typography.Text>;
};
export const DemoStringViewOnList: React.FC<ViewProps<Type.DemoString>> = ({
    value,
}) => {
    return (
        <Antd.Tooltip title={value}>
            <RA.TextField
                value={value}
                ellipsis={true}
                style={{ width: "100px" }}
            />
        </Antd.Tooltip>
    );
};
export const DemoStringViewOnForm: React.FC = (props) => {
    return <Antd.Input {...props} />;
};
// endregion String

// region Number
export const DemoNumberViewOnShow: React.FC<ViewProps<Type.DemoNumber>> = ({
    value,
    ...props
}) => {
    return <DemoStringViewOnShow value={String(value)} {...props} />;
};
export const DemoNumberViewOnList: React.FC<ViewProps<Type.DemoNumber>> = ({
    value,
    ...props
}) => {
    return <DemoNumberViewOnShow value={value} {...props} />;
};
export const DemoNumberViewOnForm: React.FC = (props) => {
    return <Antd.InputNumber {...props} />;
};
// endregion Number

// region String List
export const DemoStringListViewOnShow: React.FC<
    ViewProps<Type.DemoStringList>
> = ({ value }) => {
    return (
        <>
            {value?.map((v, i) => (
                <DemoStringViewOnShow
                    key={i}
                    value={String(i + 1) + ". " + v}
                />
            ))}
        </>
    );
};
export const DemoStringListViewOnList: React.FC<
    ViewProps<Type.DemoStringList>
> = ({ value }) => {
    return (
        <Antd.Typography.Text>
            {String(value?.length || 0) + " items"}
        </Antd.Typography.Text>
    );
};
export const DemoStringListViewOnForm: React.FC = (props) => {
    return <Antd.Select {...props} mode="tags" />;
};
// endregion

// region Image
export const DemoImageViewOnShow: React.FC<ViewProps<Type.DemoImage>> = ({
    value,
}) => {
    return (
        <Antd.Image style={{ width: "100%", maxWidth: "256px" }} src={value} />
    );
};
export const DemoImageViewOnList: React.FC<ViewProps<Type.DemoImage>> = ({
    value,
}) => {
    return <Antd.Image width={48} src={value} />;
};
export const DemoImageViewOnForm: React.FC = DemoStringViewOnForm;
// endregion Image

// region UUID
export const DemoUUIDViewOnShow: React.FC<ViewProps<Type.DemoUUID>> =
    DemoStringViewOnShow;
export const DemoUUIDViewOnList: React.FC<ViewProps<Type.DemoUUID>> =
    DemoStringViewOnList;
export const DemoUUIDViewOnForm: React.FC = DemoStringListViewOnForm;
// endregion UUID

// region Code
export const DemoCodeViewOnForm: React.FC<any> = (props) => {
    return (
        <CodeEditor
            {...props}
            language={props.Language || "js"}
            padding={15}
            style={{
                overflow: "auto",
                maxHeight: "80vh",
                minHeight: "400px",
                fontSize: 12,
                backgroundColor: "#222",
                fontFamily:
                    "ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace",
            }}
        />
    );
};
export const DemoCodeViewOnShow: React.FC<ViewProps<Type.DemoCode>> = ({
    value,
}) => {
    return <DemoCodeViewOnForm value={value} readOnly={true} />;
};
export const DemoCodeViewOnList: React.FC<ViewProps<Type.DemoCode>> = ({
    value,
}) => {
    return (
        <RA.TextField
            value={value}
            code={true}
            ellipsis={true}
            style={{ width: "100px" }}
        />
    );
};
// endregion Code

// region Rich Text
export const DemoRichTextViewOnShow: React.FC<ViewProps<Type.DemoRichText>> = ({
    value,
}) => {
    return <div dangerouslySetInnerHTML={{ __html: value || "" }}></div>;
};
export const DemoRichTextViewOnList: React.FC<ViewProps<Type.DemoRichText>> =
    DemoStringViewOnList;
export const DemoRichTextViewOnForm: React.FC = (props) => (
    <ReactQuill {...props} theme="snow" />
);
// endregion Rich Text

// region URL
export const DemoURLViewOnShow: React.FC<ViewProps<Type.DemoURL>> = ({
    value,
}) => {
    return (
        <Antd.Button
            href={value}
            target="_blank"
            icon={<AntdIcons.LinkOutlined />}
        >
            <RA.TextField
                value={value}
                ellipsis={true}
                style={{ width: "100%", maxWidth: "350px" }}
                type={"secondary"}
            />
        </Antd.Button>
    );
};
export const DemoURLViewOnList: React.FC<ViewProps<Type.DemoURL>> = ({
    value,
}) => {
    return (
        <Antd.Button
            type="primary"
            href={value}
            target="_blank"
            icon={<AntdIcons.LinkOutlined />}
        >
            Open
        </Antd.Button>
    );
};
export const DemoURLViewOnForm: React.FC = DemoStringViewOnForm;
// endregion URL

// region Enums

export const DemoEnumsProcessStatusViewOnShow = DemoStringViewOnShow;
export const DemoEnumsProcessStatusViewOnList = DemoStringViewOnList;
export const DemoEnumsProcessStatusViewOnForm: React.FC = (props) => {
    return (
        <Antd.Select
            {...props}
            options={[
                {
                    value: "none",
                    label: "None",
                },
                {
                    value: "done",
                    label: "Done",
                },
                {
                    value: "enqueued",
                    label: "Enqueued",
                },
                {
                    value: "in_progress",
                    label: "In Progress",
                },
                {
                    value: "failed",
                    label: "Failed",
                },
            ]}
        />
    );
};
// endregion Enums

// region Entity Badges

export const CompanyBadge: React.FC<Partial<Type.DemoCompanyInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.CompanyShow
                    breadcrumb={false}
                    title={props.name}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/company/show/" + props.id}>{props.name}</Link>
        </Antd.Popover>
    );
};

export const CountryBadge: React.FC<Partial<Type.DemoCountryInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.CountryShow
                    breadcrumb={false}
                    title={props.name}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/country/show/" + props.id}>{props.name}</Link>
        </Antd.Popover>
    );
};

export const EmailBadge: React.FC<Partial<Type.DemoEmailInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.EmailShow
                    breadcrumb={false}
                    title={props.title}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/email/show/" + props.id}>{props.title}</Link>
        </Antd.Popover>
    );
};

export const ImageBadge: React.FC<Partial<Type.DemoImageInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.ImageShow
                    breadcrumb={false}
                    title={props.title}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/image/show/" + props.id}>
                {props.originalURL ? (
                    <Antd.Image
                        width={48}
                        preview={false}
                        src={props.originalURL}
                    />
                ) : null}
            </Link>
        </Antd.Popover>
    );
};

export const LocationBadge: React.FC<Partial<Type.DemoLocationInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.LocationShow
                    breadcrumb={false}
                    title={props.title}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/location/show/" + props.id}>{props.title}</Link>
        </Antd.Popover>
    );
};

export const PhoneBadge: React.FC<Partial<Type.DemoPhoneInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.PhoneShow
                    breadcrumb={false}
                    title={props.title}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/phone/show/" + props.id}>{props.title}</Link>
        </Antd.Popover>
    );
};

export const ProductBadge: React.FC<Partial<Type.DemoProductInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.ProductShow
                    breadcrumb={false}
                    title={props.name}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/product/show/" + props.id}>
                {props.image ? (
                    <Antd.Image width={48} preview={false} src={props.image} />
                ) : null}
            </Link>
        </Antd.Popover>
    );
};

export const VendorBadge: React.FC<Partial<Type.DemoVendorInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.VendorShow
                    breadcrumb={false}
                    title={props.name}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/vendor/show/" + props.id}>{props.name}</Link>
        </Antd.Popover>
    );
};

export const WarehouseBadge: React.FC<Partial<Type.DemoWarehouseInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.WarehouseShow
                    breadcrumb={false}
                    title={props.name}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/warehouse/show/" + props.id}>{props.name}</Link>
        </Antd.Popover>
    );
};

export const WebsiteBadge: React.FC<Partial<Type.DemoWebsiteInterface>> = (
    props,
) => {
    const Link = useLink();
    return (
        <Antd.Popover
            overlayInnerStyle={{
                width: "50vw",
                height: "40vh",
                maxHeight: "400px",
                maxWidth: "500px",
                overflow: "auto",
            }}
            content={
                <Show.WebsiteShow
                    breadcrumb={false}
                    title={props.title}
                    id={props.id}
                    withEdges={false}
                    headerButtons={[]}
                />
            }
        >
            <Link to={"/website/show/" + props.id}>{props.title}</Link>
        </Antd.Popover>
    );
};

// endregion
