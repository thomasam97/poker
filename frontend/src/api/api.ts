const BASE_URL = "ws://localhost:7788"

class API {
    constructor(
        private url: string,
    ){}

    private conn: WebSocket;

    public register(roomID: string, playerName: string){
        const url = this.urlRoomAndPlayer(roomID, playerName)
        this.conn = new WebSocket(url);

        this.conn.onopen = (e) => {
            console.debug("[open] Connection established");
            // alert("Sending to server");
            // this.conn.send("My name is John");
          };
          
          this.conn.onmessage = (event) => {
            console.debug(`[message] Data received from server: ${event.data}`);
          };
          
          this.conn.onclose = (event) => {
            if (event.wasClean) {
              console.debug(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
            } else {
              // e.g. server process killed or network down
              // event.code is usually 1006 in this case
              console.debug('[close] Connection died');
            }
          };
          
          this.conn.onerror = (error) =>  {
            console.debug(`[error] ${JSON.stringify(error)}`);
          };

    }

    public urlRoomAndPlayer(roomID: string, playerName: string): string {
        const path = [this.url,"poker","room",roomID,playerName].join("/")
        const url = new URL(path)
        return url.toString()
    }
}

export const api = new API(BASE_URL)