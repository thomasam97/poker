export * from "api"
import { Stash } from "$lib/api/stash"
import { API } from "api"
import { config } from "../env"

export const api = new API(config.baseWS)
export const stash = new Stash()