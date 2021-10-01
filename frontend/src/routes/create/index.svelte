<script lang="ts">
import { goto } from '$app/navigation';

let roomName = Math.random().toString(36).slice(2,7).toUpperCase()
let playerName =""



function routeToPage(route: string, replaceState: boolean = true) {
   goto(`/${route}`, { replaceState }) 
}


function onSubmit(event: Event){
    event.preventDefault();

    const path = ["room",roomName,playerName].join("/")
 
    routeToPage(path)
}

</script>

<form on:submit={onSubmit}>

    <div>
        <label for="placyr-name" >Player Name</label>
        <input id="player-name" bind:value={playerName} type="text" placeholder="Jane" />
    </div>

    <div>
        <label for="room-name">Room Name</label>
        <input id="room-name" bind:value={roomName} type="text" placeholder="random room name" />
    </div>

    <div>
        <label for="select-deck">Select Deck</label>
        <select id="select-deck">
            <option value="" selected></option>
            <option value="simple">Simple</option>
            <option value="fibonacchi">Fibonacchi</option>
        </select>
    </div>

<div>
    <button type="submit">Create</button>
</div>
</form>

<style>
    form {
        display: grid;
        gap: 1rem;
    }
</style>