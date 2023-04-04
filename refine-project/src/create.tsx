import { useState } from "react";
import * as RA from "@refinedev/antd";
import * as Antd from "antd";
import * as Interfaces from "./typedefs";
import { Cursors } from "./data-provider";
import dayjs from "dayjs";
import CodeEditor from "@uiw/react-textarea-code-editor";
import * as View from "./view";
import * as Custom from "./custom";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
export const CompanyCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWECompanyInterface>();

    const [countriesCursors, setCountriesCursors] = useState<Cursors>({});
    const { selectProps: countriesSelectProps } =
        RA.useSelect<Interfaces.QWECountryInterface>({
            resource: "Country",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: countriesCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [phonesCursors, setPhonesCursors] = useState<Cursors>({});
    const { selectProps: phonesSelectProps } =
        RA.useSelect<Interfaces.QWEPhoneInterface>({
            resource: "Phone",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: phonesCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [emailsCursors, setEmailsCursors] = useState<Cursors>({});
    const { selectProps: emailsSelectProps } =
        RA.useSelect<Interfaces.QWEEmailInterface>({
            resource: "Email",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: emailsCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [websitesCursors, setWebsitesCursors] = useState<Cursors>({});
    const { selectProps: websitesSelectProps } =
        RA.useSelect<Interfaces.QWEWebsiteInterface>({
            resource: "Website",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: websitesCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [locationsCursors, setLocationsCursors] = useState<Cursors>({});
    const { selectProps: locationsSelectProps } =
        RA.useSelect<Interfaces.QWELocationInterface>({
            resource: "Location",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: locationsCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [logo_imageCursors, setLogoImageCursors] = useState<Cursors>({});
    const { selectProps: logo_imageSelectProps } =
        RA.useSelect<Interfaces.QWEImageInterface>({
            resource: "Image",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: logo_imageCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [cover_imageCursors, setCoverImageCursors] = useState<Cursors>({});
    const { selectProps: cover_imageSelectProps } =
        RA.useSelect<Interfaces.QWEImageInterface>({
            resource: "Image",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: cover_imageCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [gallery_imagesCursors, setGalleryImagesCursors] = useState<Cursors>(
        {},
    );
    const { selectProps: gallery_imagesSelectProps } =
        RA.useSelect<Interfaces.QWEImageInterface>({
            resource: "Image",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: gallery_imagesCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="name"
                    label="Name"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="description"
                    label="Description"
                    rules={[{ required: true }]}
                >
                    <View.QWERichTextViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Countries"
                    name={["countryIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...countriesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Phones"
                    name={["phoneIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...phonesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Emails"
                    name={["emailIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...emailsSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Websites"
                    name={["websiteIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...websitesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Locations"
                    name={["locationIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...locationsSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Logo Image"
                    name="logoImageID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...logo_imageSelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Cover Image"
                    name="coverImageID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...cover_imageSelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Gallery Images"
                    name={["galleryImageIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select
                        {...gallery_imagesSelectProps}
                        mode={"multiple"}
                    />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const CountryCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWECountryInterface>();

    const [companiesCursors, setCompaniesCursors] = useState<Cursors>({});
    const { selectProps: companiesSelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: companiesCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [phonesCursors, setPhonesCursors] = useState<Cursors>({});
    const { selectProps: phonesSelectProps } =
        RA.useSelect<Interfaces.QWEPhoneInterface>({
            resource: "Phone",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: phonesCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [emailsCursors, setEmailsCursors] = useState<Cursors>({});
    const { selectProps: emailsSelectProps } =
        RA.useSelect<Interfaces.QWEEmailInterface>({
            resource: "Email",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: emailsCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [websitesCursors, setWebsitesCursors] = useState<Cursors>({});
    const { selectProps: websitesSelectProps } =
        RA.useSelect<Interfaces.QWEWebsiteInterface>({
            resource: "Website",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: websitesCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [locationsCursors, setLocationsCursors] = useState<Cursors>({});
    const { selectProps: locationsSelectProps } =
        RA.useSelect<Interfaces.QWELocationInterface>({
            resource: "Location",
            optionLabel: "title",
            optionValue: "id",
            metaData: {
                cursors: locationsCursors,
                fields: ["id", "title"],
            },
            onSearch: (value) => [
                {
                    field: "title",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="name"
                    label="Name"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="code"
                    label="Code"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Companies"
                    name={["companyIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...companiesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Phones"
                    name={["phoneIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...phonesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Emails"
                    name={["emailIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...emailsSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Websites"
                    name={["websiteIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...websitesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Locations"
                    name={["locationIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...locationsSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const EmailCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWEEmailInterface>();

    const [companyCursors, setCompanyCursors] = useState<Cursors>({});
    const { selectProps: companySelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: companyCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [countryCursors, setCountryCursors] = useState<Cursors>({});
    const { selectProps: countrySelectProps } =
        RA.useSelect<Interfaces.QWECountryInterface>({
            resource: "Country",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: countryCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="title"
                    label="Title"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="description"
                    label="Description"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="address"
                    label="Address"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Company"
                    name="companyID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...companySelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Country"
                    name="countryID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...countrySelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const ImageCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWEImageInterface>();

    const [gallery_companyCursors, setGalleryCompanyCursors] =
        useState<Cursors>({});
    const { selectProps: gallery_companySelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: gallery_companyCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [logo_companyCursors, setLogoCompanyCursors] = useState<Cursors>({});
    const { selectProps: logo_companySelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: logo_companyCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [cover_companyCursors, setCoverCompanyCursors] = useState<Cursors>(
        {},
    );
    const { selectProps: cover_companySelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: cover_companyCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="title"
                    label="Title"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="originalURL"
                    label="Original Url"
                    rules={[{ required: true }]}
                >
                    <View.QWEImageViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Gallery Company"
                    name="galleryCompanyID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select
                        {...gallery_companySelectProps}
                        mode={undefined}
                    />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Logo Company"
                    name="logoCompanyID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select
                        {...logo_companySelectProps}
                        mode={undefined}
                    />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Cover Company"
                    name="coverCompanyID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select
                        {...cover_companySelectProps}
                        mode={undefined}
                    />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const LocationCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWELocationInterface>();

    const [companyCursors, setCompanyCursors] = useState<Cursors>({});
    const { selectProps: companySelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: companyCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [countryCursors, setCountryCursors] = useState<Cursors>({});
    const { selectProps: countrySelectProps } =
        RA.useSelect<Interfaces.QWECountryInterface>({
            resource: "Country",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: countryCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="title"
                    label="Title"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="description"
                    label="Description"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="latitude"
                    label="Latitude"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="longitude"
                    label="Longitude"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="address"
                    label="Address"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="postcode"
                    label="Postcode"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="type"
                    label="Type"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="state"
                    label="State"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="suburb"
                    label="Suburb"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="streetType"
                    label="Street Type"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="streetName"
                    label="Street Name"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Company"
                    name="companyID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...companySelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Country"
                    name="countryID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...countrySelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const PhoneCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWEPhoneInterface>();

    const [companyCursors, setCompanyCursors] = useState<Cursors>({});
    const { selectProps: companySelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: companyCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [countryCursors, setCountryCursors] = useState<Cursors>({});
    const { selectProps: countrySelectProps } =
        RA.useSelect<Interfaces.QWECountryInterface>({
            resource: "Country",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: countryCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="title"
                    label="Title"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="description"
                    label="Description"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="number"
                    label="Number"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="type"
                    label="Type"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Company"
                    name="companyID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...companySelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Country"
                    name="countryID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...countrySelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const ProductCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWEProductInterface>();

    const [warehouseCursors, setWarehouseCursors] = useState<Cursors>({});
    const { selectProps: warehouseSelectProps } =
        RA.useSelect<Interfaces.QWEWarehouseInterface>({
            resource: "Warehouse",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: warehouseCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [vendorCursors, setVendorCursors] = useState<Cursors>({});
    const { selectProps: vendorSelectProps } =
        RA.useSelect<Interfaces.QWEVendorInterface>({
            resource: "Vendor",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: vendorCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="name"
                    label="Name"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="description"
                    label="Description"
                    rules={[{ required: true }]}
                >
                    <View.QWERichTextViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="image"
                    label="Image"
                    rules={[{ required: true }]}
                >
                    <View.QWEImageViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="url"
                    label="Url"
                    rules={[{ required: true }]}
                >
                    <View.QWEURLViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="status"
                    label="Status"
                    rules={[{ required: true }]}
                >
                    <View.QWEEnumsProcessStatusViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="buildStatus"
                    label="Build Status"
                    rules={[{ required: true }]}
                >
                    <View.QWEEnumsProcessStatusViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Warehouse"
                    name="warehouseID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...warehouseSelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Vendor"
                    name="vendorID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...vendorSelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const VendorCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWEVendorInterface>();

    const [warehousesCursors, setWarehousesCursors] = useState<Cursors>({});
    const { selectProps: warehousesSelectProps } =
        RA.useSelect<Interfaces.QWEWarehouseInterface>({
            resource: "Warehouse",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: warehousesCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [productsCursors, setProductsCursors] = useState<Cursors>({});
    const { selectProps: productsSelectProps } =
        RA.useSelect<Interfaces.QWEProductInterface>({
            resource: "Product",
            optionLabel: "url",
            optionValue: "id",
            metaData: {
                cursors: productsCursors,
                fields: ["id", "url"],
            },
            onSearch: (value) => [
                {
                    field: "url",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="name"
                    label="Name"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="schema"
                    label="Schema"
                    rules={[{ required: true }]}
                >
                    <View.QWECodeViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Warehouses"
                    name={["warehouseIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...warehousesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Products"
                    name={["productIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...productsSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const WarehouseCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWEWarehouseInterface>();

    const [productsCursors, setProductsCursors] = useState<Cursors>({});
    const { selectProps: productsSelectProps } =
        RA.useSelect<Interfaces.QWEProductInterface>({
            resource: "Product",
            optionLabel: "url",
            optionValue: "id",
            metaData: {
                cursors: productsCursors,
                fields: ["id", "url"],
            },
            onSearch: (value) => [
                {
                    field: "url",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [vendorCursors, setVendorCursors] = useState<Cursors>({});
    const { selectProps: vendorSelectProps } =
        RA.useSelect<Interfaces.QWEVendorInterface>({
            resource: "Vendor",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: vendorCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="name"
                    label="Name"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="originalData"
                    label="Original Data"
                    rules={[{ required: false }]}
                >
                    <View.QWECodeViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="enabled"
                    label="Enabled"
                    rules={[{ required: true }]}
                >
                    <View.QWEBooleanViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="filters"
                    label="Filters"
                    rules={[{ required: false }]}
                >
                    <View.QWEStringListViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Products"
                    name={["productIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...productsSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Vendor"
                    name="vendorID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...vendorSelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const WebsiteCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.QWEWebsiteInterface>();

    const [companyCursors, setCompanyCursors] = useState<Cursors>({});
    const { selectProps: companySelectProps } =
        RA.useSelect<Interfaces.QWECompanyInterface>({
            resource: "Company",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: companyCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [countryCursors, setCountryCursors] = useState<Cursors>({});
    const { selectProps: countrySelectProps } =
        RA.useSelect<Interfaces.QWECountryInterface>({
            resource: "Country",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: countryCursors,
                fields: ["id", "name"],
            },
            onSearch: (value) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });

    return (
        <RA.Create saveButtonProps={saveButtonProps}>
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="title"
                    label="Title"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="description"
                    label="Description"
                    rules={[{ required: true }]}
                >
                    <View.QWEStringViewOnForm />
                </Antd.Form.Item>
                <Antd.Form.Item
                    name="url"
                    label="Url"
                    rules={[{ required: true }]}
                >
                    <View.QWEURLViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Company"
                    name="companyID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...companySelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Country"
                    name="countryID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...countrySelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
