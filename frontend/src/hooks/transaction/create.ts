// This file is generated by greenstar scripts. DO NOT EDIT.
// noinspection DuplicatedCode

import { Transaction } from "../../models/transaction.ts"
import { DateProperties } from "../../models/transaction.ts"
import { BaseAPIURL } from "../../services/util.ts"
import { useOperation, Method, Hook } from "../../util/operation.ts"

export interface Request {
    amount: number
    currency: string
    date: Date
    description?: string
    referenceId: string
    sourceAccountId: string
    targetAccountId: string
}
export type Response = Transaction | undefined

export function useCreateTransaction(): Hook<Request, Response> {
    const opts = {
        initial: undefined,
        method: "POST" as Method,
        url: `${BaseAPIURL}/transactions`,
        dateProperties: DateProperties,
    }
    return useOperation<Request, Response>(opts)
}
