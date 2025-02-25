// This file is generated by greenstar scripts. DO NOT EDIT.
// noinspection DuplicatedCode

import { Tenant } from "../../models/tenant.ts"
import { DateProperties } from "../../models/tenant.ts"
import { BaseAPIURL } from "../../services/util.ts"
import { useOperation, Method, Hook } from "../../util/operation.ts"

export interface Request {
    id?: string
    displayName: string
}
export type Response = Tenant | undefined

export function useCreateTenant(): Hook<Request, Response> {
    const opts = {
        initial: undefined,
        method: "POST" as Method,
        url: `${BaseAPIURL}/tenants`,
        dateProperties: DateProperties,
    }
    return useOperation<Request, Response>(opts)
}
