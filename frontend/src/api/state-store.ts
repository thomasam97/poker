import { writable } from "svelte/store";

export enum Status {
    Start =      "Start",
    InProgress = "InProgress",
    Revealed =   "Revealed",
}

export interface State {
    player?:  Player
    players?: Player[]
    status?:  Status
}

interface Player {
    id:   string
    name: string
}

const initalState: State = {
    players: [],
    status: Status.Start,
}

const state = writable(initalState)

export type UnsubscriberFn = () => void
export type HandlerFn = (state: State) => void 

export function subscribe(handler: HandlerFn): UnsubscriberFn {
    return state.subscribe(handler)
}

export const setState = state.set