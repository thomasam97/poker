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

<div class="room-controller-root">
{#if status === Status.Start}
    <button on:click={onStartClick}> Start Voting</button>
{/if}

{#if status === Status.InProgress}
    <button on:click={onRevealClick} disabled={player.chosenCard === ""}> Reveal Cards</button>
{/if}
 
{#if status === Status.Revealed}
    <button on:click={onResetClick}> Reset </button>
{/if}
</div>

<style>

    .room-controller-root{
        display:         flex;
        justify-content: center;
    }

    button {
        width: 10rem;
    }
</style>