import { config } from "$lib/env"
import type { HandlerFn, PlayerType } from "./state-store";
const BASE_URL = config.baseWS

enum ActionTypes {
    StartVoting   = "StartVoting",
    Reveal        = "Reveal",
    Reset         = "Reset",
    Choose        = "Choose",
    SetPlayerType = "SetPlayerType",
    SetCards      = "SetCards",
}

export interface Message<T> {
    type:     ActionTypes,
    payload?: T, 
}

class API {
    constructor(
        private url: string,
    ){}
        
    private conn: WebSocket;

    public setPlayerType(type: PlayerType){
        const msg: Message<PlayerType> = {
            type:     ActionTypes.SetPlayerType,
            payload: type,
        }
        this.send(msg);
    }

    public startRoom() {
        const msg = {
            type: ActionTypes.StartVoting
        }
        this.send(msg)
    }

    public revealRoom() {
        const msg = {
            type: ActionTypes.Reveal
        }
        this.send(msg)
    }
  
    public resetRoom() {
        const msg = {
            type: ActionTypes.Reset
        }
        this.send(msg)
    }

    public choose(card: string){
        const msg = {
            type: ActionTypes.Choose,
            payload: card,
        }
        this.send(msg)
    }

    public setCards(cards: string[]) {
        const msg = {
            type: ActionTypes.SetCards,
            payload: cards,
        }
        this.send(msg)
    }

    private send<T>(msg: Message<T>){
        if(!this.conn){ return }

        const payload = JSON.stringify(msg)
        this.conn.send(payload);
    }
    
    public register(roomID: string, playerName: string, handlerFn: HandlerFn){
        const url = this.urlRoomAndPlayer(roomID, playerName)
        let conn = this.conn
        if( !this.conn ){
            this.conn = new WebSocket(url);
            conn = this.conn
        }
        
        conn.onopen = (e) => {
            console.debug("[open] Connection established");
        };
        
        conn.onmessage = (event) => {
            console.debug(`[message] Data received from server: ${event.data}`);
            const parsedData = JSON.parse(event.data)
            handlerFn(parsedData)
        };
        
        conn.onclose = (event) => {
            if (event.wasClean) {
                console.debug(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
            } else {
                // e.g. server process killed or network down
                // event.code is usually 1006 in this case
                console.debug('[close] Connection died');
            }
        };
        
        conn.onerror = (error) =>  {
            console.debug(`[error] ${JSON.stringify(error)}`);
        };
        
    }
    
        public urlRoomAndPlayer(roomID: string, playerName: string): string {
            const path = [this.url,"room",roomID,playerName].join("/")
            const url = new URL(path)
            return url.toString()
        }
    }
    
    export const api = new API(BASE_URL)