<script lang="ts">
    import { api } from "../../api"
    import { Status } from "../../api/state-store";
import type { Player } from "../player-list/player";

    export let status = "";
    export let player: Player

    function onStartClick(){
        api.startRoom()
    }

    function onRevealClick(){
        api.revealRoom()
    }

    function onResetClick(){
        api.resetRoom()
    }

</script>
<room-controller>
    {#if status === Status.Start}
        <button on:click={onStartClick}> Start Voting</button>
    {/if}

    {#if status === Status.InProgress}
        <button on:click={onRevealClick} disabled={player.chosenCard === ""}> Reveal Cards</button>
    {/if}
    
    {#if status === Status.Revealed}
        <button on:click={onResetClick}> Reset </button>
    {/if}


    <div>
        <label for="topic">Topic</label>
        <input id="topic" type="text" />
    </div>

</room-controller>

<style>

    room-controller {
        background:      transparent;
        /* border:          1px white solid; */
        display:         flex;
        /* padding:         1rem; */
        width:           auto;
        gap:             2rem;
        justify-content: center;
        
        --shadow-color: rgb(209, 209, 209);
        --shadow: 1px 0px 3px 1px var(--shadow-color), inset 0px 0px 3px 1px var(--shadow-color); ; 

        outline:            none;
        /* -webkit-box-shadow: var(--shadow); */
        /* box-shadow: 		var(--shadow); */
    }
/* 
    room-controller:focus,
    room-controller:focus-within,
    room-controller:active {
        
    } */

    button {
        width: 10rem;
    }
</style>