export * from "api"
import { API } from "api"
import { config } from "$lib/env"

export const api = new API(config.baseWS)