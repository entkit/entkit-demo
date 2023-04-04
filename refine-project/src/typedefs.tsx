// Custom types
export type QWEString = string;
export type QWENumber = number;
export type QWEBoolean = boolean;
export type QWEDate = Date;
export type QWEUUID = string;
export type QWEImage = string;
export type QWECode = string;
export type QWEURL = string;
export type QWERichText = string;
export type QWEStringList = QWEString[];
export type QWENumberList = QWENumber[];
export type QWEID = QWEUUID | QWENumber;

export enum QWEEnumsProcessStatus {
    "none" = "none",
    "done" = "done",
    "enqueued" = "enqueued",
    "in_progress" = "in_progress",
    "failed" = "failed",
}

interface gqlField<T> {
    edges?: {
        nodes?: Array<T>;
    };
    totalCount?: number;
}

export interface QWECompanyInterface {
    id: QWEUUID; // uuid.UUID
    name: QWEString; // string
    description: QWERichText; // string
    countries?: QWECountryInterface[];
    _countries?: gqlField<QWECountryInterface>;
    phones?: QWEPhoneInterface[];
    _phones?: gqlField<QWEPhoneInterface>;
    emails?: QWEEmailInterface[];
    _emails?: gqlField<QWEEmailInterface>;
    websites?: QWEWebsiteInterface[];
    _websites?: gqlField<QWEWebsiteInterface>;
    locations?: QWELocationInterface[];
    _locations?: gqlField<QWELocationInterface>;
    logoImage?: QWEImageInterface;
    coverImage?: QWEImageInterface;
    galleryImages?: QWEImageInterface[];
    _galleryImages?: gqlField<QWEImageInterface>;
}
export interface QWECountryInterface {
    id: QWEUUID; // uuid.UUID
    name: QWEString; // string
    code: QWEString; // string
    companies?: QWECompanyInterface[];
    _companies?: gqlField<QWECompanyInterface>;
    phones?: QWEPhoneInterface[];
    _phones?: gqlField<QWEPhoneInterface>;
    emails?: QWEEmailInterface[];
    _emails?: gqlField<QWEEmailInterface>;
    websites?: QWEWebsiteInterface[];
    _websites?: gqlField<QWEWebsiteInterface>;
    locations?: QWELocationInterface[];
    _locations?: gqlField<QWELocationInterface>;
}
export interface QWEEmailInterface {
    id: QWEUUID; // uuid.UUID
    title: QWEString; // string
    description: QWEString; // string
    address: QWEString; // string
    company?: QWECompanyInterface;
    country?: QWECountryInterface;
}
export interface QWEImageInterface {
    id: QWEUUID; // uuid.UUID
    title: QWEString; // string
    originalURL: QWEImage; // string
    galleryCompany?: QWECompanyInterface;
    logoCompany?: QWECompanyInterface;
    coverCompany?: QWECompanyInterface;
}
export interface QWELocationInterface {
    id: QWEUUID; // uuid.UUID
    title: QWEString; // string
    description: QWEString; // string
    latitude: QWEString; // float64
    longitude: QWEString; // float64
    address: QWEString; // string
    postcode: QWEString; // string
    type: QWEString; // string
    state: QWEString; // string
    suburb: QWEString; // string
    streetType: QWEString; // string
    streetName: QWEString; // string
    company?: QWECompanyInterface;
    country?: QWECountryInterface;
}
export interface QWEPhoneInterface {
    id: QWEUUID; // uuid.UUID
    title: QWEString; // string
    description: QWEString; // string
    number: QWEString; // string
    type: QWEString; // string
    company?: QWECompanyInterface;
    country?: QWECountryInterface;
}
export interface QWEProductInterface {
    id: QWEUUID; // uuid.UUID
    name: QWEString; // string
    description: QWERichText; // string
    image: QWEImage; // string
    url: QWEURL; // string
    lastSell: QWEDate; // time.Time
    createdAt: QWEDate; // time.Time
    status: QWEEnumsProcessStatus; // enums.ProcessStatus
    buildStatus: QWEEnumsProcessStatus; // enums.ProcessStatus
    warehouse?: QWEWarehouseInterface;
    vendor?: QWEVendorInterface;
}
export interface QWEVendorInterface {
    id: QWEUUID; // uuid.UUID
    name: QWEString; // string
    schema: QWECode; // string
    warehouses?: QWEWarehouseInterface[];
    _warehouses?: gqlField<QWEWarehouseInterface>;
    products?: QWEProductInterface[];
    _products?: gqlField<QWEProductInterface>;
}
export interface QWEWarehouseInterface {
    id: QWEUUID; // uuid.UUID
    name: QWEString; // string
    lastUpdate: QWEDate; // time.Time
    originalData: QWECode; // string
    enabled: QWEBoolean; // bool
    filters: QWEStringList; // []string
    products?: QWEProductInterface[];
    _products?: gqlField<QWEProductInterface>;
    vendor?: QWEVendorInterface;
}
export interface QWEWebsiteInterface {
    id: QWEUUID; // uuid.UUID
    title: QWEString; // string
    description: QWEString; // string
    url: QWEURL; // string
    company?: QWECompanyInterface;
    country?: QWECountryInterface;
}
