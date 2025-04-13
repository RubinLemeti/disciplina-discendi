import { Response } from "express"
import { CollectionResponseModel, ErrorSubModel, FailureResponseModel, Meta, ResourceResponseModel, SuccessfulResponseModel } from "./genericResponseInterfaces"

export function SuccessJsonStandardResponse(r: Response, path: string, statusCode?: number) {
    let responseStatus = statusCode ? statusCode : 200
    let response: SuccessfulResponseModel = {
        path: path,
        success: true,
        statusCode: responseStatus,
        timestamp: new Date().toISOString().slice(0, 19)
    }
    r.status(responseStatus).send(response)
}

export function ErrorJsonStandardResponse(r: Response, path: string, statusCode?: number, title?: string, details?: string) {
    let error: ErrorSubModel = {
        title: title ? title : "Internal Server Error",
        details: details ? details : "An unexpected error occurred"
    }
    let responseStatus = statusCode ? statusCode : 500
    let response: FailureResponseModel = {
        error: error,
        path: path,
        success: false,
        statusCode: responseStatus,
        timestamp: new Date().toISOString().slice(0, 19)
    }
    r.status(responseStatus).send(response)
}

export function ResourceJsonStandardResponse(r: Response, path: string, data: any, statusCode?: number) {
    let responseStatus = statusCode ? statusCode : 200
    let response: ResourceResponseModel = {
        data: data,
        path: path,
        success: true,
        statusCode: responseStatus,
        timestamp: new Date().toISOString().slice(0, 19)
    }
    r.status(responseStatus).send(response)
}

export function CollectionJsonStandardResponse(r: Response, path: string, data: any, meta: Meta, statusCode?: number) {
    let responseStatus = statusCode ? statusCode : 200
    let response: CollectionResponseModel = {
        data: data,
        meta: meta,
        path: path,
        success: true,
        statusCode: responseStatus,
        timestamp: new Date().toISOString().slice(0, 19)
    }
    r.status(responseStatus).send(response)
}