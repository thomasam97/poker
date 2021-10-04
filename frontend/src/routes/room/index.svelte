<script lang="ts">
    import { goto } from '$app/navigation';
    import { page } from "$app/stores";
    
    
    const { 
        roomid:     roomID,
    } = $page.params

    let playerName = "";
    let roomName = randomRoomName();
    if( roomID ){
        roomName = roomID
    }
    
    function routeToPage(route: string, replaceState: boolean = true) {
       goto(`/${route}`, { replaceState }) 
    }
    
    
    function onSubmit(event: Event){
        event.preventDefault();
    
        const path = ["room",roomName,playerName].join("/")
     
        routeToPage(path)
    }

    function randomRoomName(): string {
        const name = Math.random().toString(36).slice(2,7).toUpperCase()
        return name
    }
    
</script>


<svelte:head>
	<title>Poker: Room Creation</title>
</svelte:head>


    <main>
        <form on:submit={onSubmit}>
    
            <div>
                <label for="room-name">Room Name</label>
                <input id="room-name" bind:value={roomName} type="text" placeholder="random room name" />
            </div>

            <div>
                <label for="placyr-name" >Player Name</label>
                <input id="player-name" bind:value={playerName} type="text" placeholder="Jane" />
            </div>
    
    
            <div class="button-container">
                <button type="submit">Create</button>
            </div>
        </form>
    </main>
    
    <style>
        main {
            display: flex;
            justify-content: center;
        }
        form {
            display: grid;
            gap:     1rem;
            width:   auto;
        }
    
        label {
            display:        block;
            font-family:    "SoinSansProHeadline";
            text-transform: uppercase;
            font-weight:    400;
            color:          white;
        }
    
        .button-container {
            text-align: right;
        }
    
    </style>