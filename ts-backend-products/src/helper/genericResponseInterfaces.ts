interface Href {
    href: string
}

interface Links {
    first: Href
    previous: Href
    current: Href
    next: Href
    last: Href
}

interface Records {
    first: number
    last: number
    total: number
    limit: number
    offset: number
}

export interface Pagination {
    first: number
    previous: number
    current: number
    next: number
    last: number
    links: Links
    records: Records
}

export interface Meta {
    pagination: Pagination
}

export interface ResourceResponseModel {
    data: any;
    path: string;
    success: boolean;
    statusCode: number;
    timestamp: string;
}

export interface CollectionResponseModel {
    data: any;
    meta: Meta;
    path: string;
    success: boolean;
    statusCode: number;
    timestamp: string;
}

export interface SuccessfulResponseModel {
    path: string;
    success: boolean;
    statusCode: number;
    timestamp: string;
}

export interface ErrorSubModel {
    title: string
    details: string
}

export interface FailureResponseModel {
    error: ErrorSubModel
    path: string;
    success: boolean;
    statusCode: number;
    timestamp: string;
}