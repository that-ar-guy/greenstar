// This file is generated by greenstar scripts. DO NOT EDIT.
// noinspection DuplicatedCode

import { Tenant } from "../../models/tenant.ts"
import { DateProperties } from "../../models/tenant.ts"
import { BaseAPIURL, QueryNilValue } from "../../services/util.ts"
import { useOperation, Method, Hook } from "../../util/operation.ts"

export type SortableProperty = "displayName"

export interface Request {
    offset?: number
    count?: number
    sort?: SortableProperty[]
    displayName?: string
    id?: string
}
export interface Response {
    offset: number
    totalCount: number
    items: Tenant[]
}

function buildURL(req: Request): string {
    const urlParams = new URLSearchParams()

    if (req.offset !== undefined && req.offset >= 0) {
        urlParams.set("_offset", req.offset + "")
    }
    if (req.count !== undefined && req.count >= 0) {
        urlParams.set("_count", req.count + "")
    }
    if (req.sort !== undefined && req?.sort.length) {
        req.sort.forEach((s) => urlParams.append("_sort", s))
    }
    if (req.displayName !== undefined) {
        if (req.displayName === null) {
            urlParams.set("displayName", QueryNilValue)
        } else {
            urlParams.set("displayName", req.displayName)
        }
    }
    if (req.id !== undefined) {
        if (req.id === null) {
            urlParams.set("id", QueryNilValue)
        } else {
            urlParams.set("id", req.id)
        }
    }

    return `${BaseAPIURL}/tenants${urlParams.size ? "?" + urlParams.toString() : ""}`
}

export function useListTenants(): Hook<Request, Response> {
    const opts = {
        initial: { offset: 0, totalCount: 0, items: [] },
        method: "GET" as Method,
        url: buildURL,
        dateProperties: DateProperties,
    }
    return useOperation<Request, Response>(opts)
}
