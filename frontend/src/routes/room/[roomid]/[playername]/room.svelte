<script>
import { page } from "$app/stores";
import { api } from "../../../../api"
import { Cards } from "../../../../components/cards";
import { PlayerList } from "../../../../components/player-list"
import { RoomController } from "../../../../components/room-controller";
import { RoomStatus } from "../../../../components/room-status"


const { 
    roomid:     roomID,
    playername: playerName 
} = $page.params


let player;
let players = []
let roomStatus = ""

console.debug('[DEBUG] ', {roomID, playerName} )
api.register(roomID, playerName, (data) => {
    console.debug('[DEBUG] ', {data} )

    player = data.player
    players = data.players
    roomStatus = data.status

})

</script>


{#if player?.isAdmin}
<RoomController status={roomStatus}/>
{/if}

<RoomStatus status={roomStatus}/>
<PlayerList players={players} />
<Cards />