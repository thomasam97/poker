export * from "api"
import { API } from "api"
import { config } from "../env"

export const api = new API(config.baseWS)