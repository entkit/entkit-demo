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
export const QWEDateViewOnShow: React.FC<ViewProps<Date>> = ({ value }) => {
    return value ? (
        <RA.DateField format="LLL" value={value} />
    ) : (
        <RA.TextField value="Never" />
    );
};
export const QWEDateViewOnList: React.FC<ViewProps<Date>> = QWEDateViewOnShow;
export const QWEDateViewOnForm: React.FC = (props) => {
    return <Antd.DatePicker {...props} showTime={true} />;
};
// endregion Date

// region Boolean
export const QWEBooleanViewOnShow: React.FC<ViewProps<Type.QWEBoolean>> = ({
    value,
}) => {
    return <RA.TextField value={value ? "Yes" : "No"} />;
};
export const QWEBooleanViewOnList = QWEBooleanViewOnShow;
export const QWEBooleanViewOnForm: React.FC = (props) => {
    return (
        <Antd.Radio.Group {...props}>
            <Antd.Radio value={true}>Yes</Antd.Radio>
            <Antd.Radio value={false}>No</Antd.Radio>
        </Antd.Radio.Group>
    );
};
// endregion Boolean

// region String
export const QWEStringViewOnShow: React.FC<ViewProps<Type.QWEString>> = ({
    value,
}) => {
    return <Antd.Typography.Text copyable={true}>{value}</Antd.Typography.Text>;
};
export const QWEStringViewOnList: React.FC<ViewProps<Type.QWEString>> = ({
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
export const QWEStringViewOnForm: React.FC = (props) => {
    return <Antd.Input {...props} />;
};
// endregion String

// region Number
export const QWENumberViewOnShow: React.FC<ViewProps<Type.QWENumber>> = ({
    value,
    ...props
}) => {
    return <QWEStringViewOnShow value={String(value)} {...props} />;
};
export const QWENumberViewOnList: React.FC<ViewProps<Type.QWENumber>> = ({
    value,
    ...props
}) => {
    return <QWENumberViewOnShow value={value} {...props} />;
};
export const QWENumberViewOnForm: React.FC = (props) => {
    return <Antd.InputNumber {...props} />;
};
// endregion Number

// region String List
export const QWEStringListViewOnShow: React.FC<
    ViewProps<Type.QWEStringList>
> = ({ value }) => {
    return (
        <>
            {value?.map((v, i) => (
                <QWEStringViewOnShow key={i} value={String(i + 1) + ". " + v} />
            ))}
        </>
    );
};
export const QWEStringListViewOnList: React.FC<
    ViewProps<Type.QWEStringList>
> = ({ value }) => {
    return (
        <Antd.Typography.Text>
            {String(value?.length || 0) + " items"}
        </Antd.Typography.Text>
    );
};
export const QWEStringListViewOnForm: React.FC = (props) => {
    return <Antd.Select {...props} mode="tags" />;
};
// endregion

// region Image
export const QWEImageViewOnShow: React.FC<ViewProps<Type.QWEImage>> = ({
    value,
}) => {
    return (
        <Antd.Image style={{ width: "100%", maxWidth: "256px" }} src={value} />
    );
};
export const QWEImageViewOnList: React.FC<ViewProps<Type.QWEImage>> = ({
    value,
}) => {
    return <Antd.Image width={48} src={value} />;
};
export const QWEImageViewOnForm: React.FC = QWEStringViewOnForm;
// endregion Image

// region UUID
export const QWEUUIDViewOnShow: React.FC<ViewProps<Type.QWEUUID>> =
    QWEStringViewOnShow;
export const QWEUUIDViewOnList: React.FC<ViewProps<Type.QWEUUID>> =
    QWEStringViewOnList;
export const QWEUUIDViewOnForm: React.FC = QWEStringListViewOnForm;
// endregion UUID

// region Code
export const QWECodeViewOnForm: React.FC<any> = (props) => {
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
export const QWECodeViewOnShow: React.FC<ViewProps<Type.QWECode>> = ({
    value,
}) => {
    return <QWECodeViewOnForm value={value} readOnly={true} />;
};
export const QWECodeViewOnList: React.FC<ViewProps<Type.QWECode>> = ({
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
export const QWERichTextViewOnShow: React.FC<ViewProps<Type.QWERichText>> = ({
    value,
}) => {
    return <div dangerouslySetInnerHTML={{ __html: value || "" }}></div>;
};
export const QWERichTextViewOnList: React.FC<ViewProps<Type.QWERichText>> =
    QWEStringViewOnList;
export const QWERichTextViewOnForm: React.FC = (props) => (
    <ReactQuill {...props} theme="snow" />
);
// endregion Rich Text

// region URL
export const QWEURLViewOnShow: React.FC<ViewProps<Type.QWEURL>> = ({
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
export const QWEURLViewOnList: React.FC<ViewProps<Type.QWEURL>> = ({
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
export const QWEURLViewOnForm: React.FC = QWEStringViewOnForm;
// endregion URL

// region Enums

export const QWEEnumsProcessStatusViewOnShow = QWEStringViewOnShow;
export const QWEEnumsProcessStatusViewOnList = QWEStringViewOnList;
export const QWEEnumsProcessStatusViewOnForm: React.FC = (props) => {
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

export const CompanyBadge: React.FC<Partial<Type.QWECompanyInterface>> = (
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

export const CountryBadge: React.FC<Partial<Type.QWECountryInterface>> = (
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

export const EmailBadge: React.FC<Partial<Type.QWEEmailInterface>> = (
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

export const ImageBadge: React.FC<Partial<Type.QWEImageInterface>> = (
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

export const LocationBadge: React.FC<Partial<Type.QWELocationInterface>> = (
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

export const PhoneBadge: React.FC<Partial<Type.QWEPhoneInterface>> = (
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

export const ProductBadge: React.FC<Partial<Type.QWEProductInterface>> = (
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

export const VendorBadge: React.FC<Partial<Type.QWEVendorInterface>> = (
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

export const WarehouseBadge: React.FC<Partial<Type.QWEWarehouseInterface>> = (
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

export const WebsiteBadge: React.FC<Partial<Type.QWEWebsiteInterface>> = (
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
