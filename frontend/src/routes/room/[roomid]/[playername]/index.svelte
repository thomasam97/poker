<script context="module">
	export const ssr = false;
</script>
<script>
import { page } from "$app/stores";
import { api } from "../../../../api"

const { 
    roomid: roomID,
    playername: playerName 
} = $page.params

console.debug('[DEBUG] ', {roomID, playerName} )
api.register(roomID, playerName)

</script>

<h1>Playername</h1>