// Custom types
export type EntString = string
export type EntNumber = number
export type EntBoolean = boolean
export type EntDate = Date
export type EntUUID = string
export type EntImage = string
export type EntCode = string
export type EntURL = string
export type EntRichText = string
export type EntStringList = EntString[]
export type EntNumberList = EntNumber[]
export type EntID = EntUUID | EntNumber


export enum EntEnumsProcessStatus{
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


export interface EntCompanyInterface {
    id: EntUUID, // uuid.UUID
    name: EntString, // string
    description: EntRichText, // string
    countries?: EntCountryInterface[],
    _countries?: gqlField<EntCountryInterface>,
    phones?: EntPhoneInterface[],
    _phones?: gqlField<EntPhoneInterface>,
    emails?: EntEmailInterface[],
    _emails?: gqlField<EntEmailInterface>,
    websites?: EntWebsiteInterface[],
    _websites?: gqlField<EntWebsiteInterface>,
    locations?: EntLocationInterface[],
    _locations?: gqlField<EntLocationInterface>,
    logoImage?: EntImageInterface,
    coverImage?: EntImageInterface,
    galleryImages?: EntImageInterface[],
    _galleryImages?: gqlField<EntImageInterface>,
}
export interface EntCountryInterface {
    id: EntUUID, // uuid.UUID
    name: EntString, // string
    code: EntString, // string
    companies?: EntCompanyInterface[],
    _companies?: gqlField<EntCompanyInterface>,
    phones?: EntPhoneInterface[],
    _phones?: gqlField<EntPhoneInterface>,
    emails?: EntEmailInterface[],
    _emails?: gqlField<EntEmailInterface>,
    websites?: EntWebsiteInterface[],
    _websites?: gqlField<EntWebsiteInterface>,
    locations?: EntLocationInterface[],
    _locations?: gqlField<EntLocationInterface>,
}
export interface EntEmailInterface {
    id: EntUUID, // uuid.UUID
    title: EntString, // string
    description: EntString, // string
    address: EntString, // string
    company?: EntCompanyInterface,
    country?: EntCountryInterface,
}
export interface EntImageInterface {
    id: EntUUID, // uuid.UUID
    title: EntString, // string
    originalURL: EntImage, // string
    galleryCompany?: EntCompanyInterface,
    logoCompany?: EntCompanyInterface,
    coverCompany?: EntCompanyInterface,
}
export interface EntLocationInterface {
    id: EntUUID, // uuid.UUID
    title: EntString, // string
    description: EntString, // string
    latitude: EntString, // float64
    longitude: EntString, // float64
    address: EntString, // string
    postcode: EntString, // string
    type: EntString, // string
    state: EntString, // string
    suburb: EntString, // string
    streetType: EntString, // string
    streetName: EntString, // string
    company?: EntCompanyInterface,
    country?: EntCountryInterface,
}
export interface EntPhoneInterface {
    id: EntUUID, // uuid.UUID
    title: EntString, // string
    description: EntString, // string
    number: EntString, // string
    type: EntString, // string
    company?: EntCompanyInterface,
    country?: EntCountryInterface,
}
export interface EntProductInterface {
    id: EntUUID, // uuid.UUID
    name: EntString, // string
    description: EntRichText, // string
    image: EntImage, // string
    url: EntURL, // string
    lastSell: EntDate, // time.Time
    createdAt: EntDate, // time.Time
    status: EntEnumsProcessStatus, // enums.ProcessStatus
    buildStatus: EntEnumsProcessStatus, // enums.ProcessStatus
    warehouse?: EntWarehouseInterface,
    vendor?: EntVendorInterface,
}
export interface EntVendorInterface {
    id: EntUUID, // uuid.UUID
    name: EntString, // string
    schema: EntCode, // string
    warehouses?: EntWarehouseInterface[],
    _warehouses?: gqlField<EntWarehouseInterface>,
    products?: EntProductInterface[],
    _products?: gqlField<EntProductInterface>,
}
export interface EntWarehouseInterface {
    id: EntUUID, // uuid.UUID
    name: EntString, // string
    lastUpdate: EntDate, // time.Time
    originalData: EntCode, // string
    enabled: EntBoolean, // bool
    filters: EntStringList, // []string
    products?: EntProductInterface[],
    _products?: gqlField<EntProductInterface>,
    vendor?: EntVendorInterface,
}
export interface EntWebsiteInterface {
    id: EntUUID, // uuid.UUID
    title: EntString, // string
    description: EntString, // string
    url: EntURL, // string
    company?: EntCompanyInterface,
    country?: EntCountryInterface,
}