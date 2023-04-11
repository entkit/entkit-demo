// ---------------------------------------------------------
// THIS FILE HAS BEEN GENERATED BY EntKit. DO NOT EDIT.
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

export type DemoString = string
export type DemoNumber = number
export type DemoBoolean = boolean
export type DemoDate = Date
export type DemoUUID = string
export type DemoImage = string
export type DemoCode = string
export type DemoURL = string
export type DemoRichText = string
export type DemoStringList = DemoString[]
export type DemoNumberList = DemoNumber[]
export type DemoID = DemoUUID | DemoNumber


export enum DemoEnumsProcessStatus{
    "none" = "none",
    "done" = "done",
    "enqueued" = "enqueued",
    "in_progress" = "in_progress",
    "failed" = "failed",
}

interface gqlField<T> {
    edges?: {
        nodes?: Array<T>
    },
    totalCount?: number,
}


export interface DemoCompanyInterface {
    id: DemoUUID, // uuid.UUID
    name: DemoString, // string
    description: DemoRichText, // string
    countries?: DemoCountryInterface[],
    _countries?: gqlField<DemoCountryInterface>,
    phones?: DemoPhoneInterface[],
    _phones?: gqlField<DemoPhoneInterface>,
    emails?: DemoEmailInterface[],
    _emails?: gqlField<DemoEmailInterface>,
    websites?: DemoWebsiteInterface[],
    _websites?: gqlField<DemoWebsiteInterface>,
    locations?: DemoLocationInterface[],
    _locations?: gqlField<DemoLocationInterface>,
    logoImage?: DemoImageInterface,
    coverImage?: DemoImageInterface,
    galleryImages?: DemoImageInterface[],
    _galleryImages?: gqlField<DemoImageInterface>,
}
export interface DemoCountryInterface {
    id: DemoUUID, // uuid.UUID
    name: DemoString, // string
    code: DemoString, // string
    companies?: DemoCompanyInterface[],
    _companies?: gqlField<DemoCompanyInterface>,
    phones?: DemoPhoneInterface[],
    _phones?: gqlField<DemoPhoneInterface>,
    emails?: DemoEmailInterface[],
    _emails?: gqlField<DemoEmailInterface>,
    websites?: DemoWebsiteInterface[],
    _websites?: gqlField<DemoWebsiteInterface>,
    locations?: DemoLocationInterface[],
    _locations?: gqlField<DemoLocationInterface>,
}
export interface DemoEmailInterface {
    id: DemoUUID, // uuid.UUID
    title: DemoString, // string
    description: DemoString, // string
    address: DemoString, // string
    company?: DemoCompanyInterface,
    country?: DemoCountryInterface,
}
export interface DemoImageInterface {
    id: DemoUUID, // uuid.UUID
    title: DemoString, // string
    originalURL: DemoImage, // string
    galleryCompany?: DemoCompanyInterface,
    logoCompany?: DemoCompanyInterface,
    coverCompany?: DemoCompanyInterface,
}
export interface DemoLocationInterface {
    id: DemoUUID, // uuid.UUID
    title: DemoString, // string
    description: DemoString, // string
    latitude: DemoString, // float64
    longitude: DemoString, // float64
    address: DemoString, // string
    postcode: DemoString, // string
    type: DemoString, // string
    state: DemoString, // string
    suburb: DemoString, // string
    streetType: DemoString, // string
    streetName: DemoString, // string
    company?: DemoCompanyInterface,
    country?: DemoCountryInterface,
}
export interface DemoPhoneInterface {
    id: DemoUUID, // uuid.UUID
    title: DemoString, // string
    description: DemoString, // string
    number: DemoString, // string
    type: DemoString, // string
    company?: DemoCompanyInterface,
    country?: DemoCountryInterface,
}
export interface DemoProductInterface {
    id: DemoUUID, // uuid.UUID
    name: DemoString, // string
    description: DemoRichText, // string
    image: DemoImage, // string
    url: DemoURL, // string
    lastSell: DemoDate, // time.Time
    createdAt: DemoDate, // time.Time
    status: DemoEnumsProcessStatus, // enums.ProcessStatus
    buildStatus: DemoEnumsProcessStatus, // enums.ProcessStatus
    warehouse?: DemoWarehouseInterface,
    vendor?: DemoVendorInterface,
}
export interface DemoVendorInterface {
    id: DemoUUID, // uuid.UUID
    name: DemoString, // string
    schema: DemoCode, // string
    warehouses?: DemoWarehouseInterface[],
    _warehouses?: gqlField<DemoWarehouseInterface>,
    products?: DemoProductInterface[],
    _products?: gqlField<DemoProductInterface>,
}
export interface DemoWarehouseInterface {
    id: DemoUUID, // uuid.UUID
    name: DemoString, // string
    lastUpdate: DemoDate, // time.Time
    originalData: DemoCode, // string
    enabled: DemoBoolean, // bool
    filters: DemoStringList, // []string
    products?: DemoProductInterface[],
    _products?: gqlField<DemoProductInterface>,
    vendor?: DemoVendorInterface,
}
export interface DemoWebsiteInterface {
    id: DemoUUID, // uuid.UUID
    title: DemoString, // string
    description: DemoString, // string
    url: DemoURL, // string
    company?: DemoCompanyInterface,
    country?: DemoCountryInterface,
}