<script lang="ts">
    import { PlayerType, Status, api } from "$lib/api";
    import type { Player } from "api";
    import RoomLink from "./room-link.svelte";


    export let status = "";
    export let player: Player
    export let roomID = "";

    function onStartClick(){
        api.startRoom()
    }

    function onRevealClick(){
        api.revealRoom()
    }

    function onResetClick(){
        api.resetRoom()
        api.startRoom()
    }

    const playerTypes = [
        PlayerType.Player,
        PlayerType.Spectator,
    ]

    function isRevealDisabled(player: Player): boolean {
        return player.type !== PlayerType.Spectator && player.chosenCard === "" 
    }

    $: roomLink = `${window.location.origin}/room/${roomID}`
    
    const key="button"

   
</script>

<room-controller>
    <RoomLink link={roomLink} />

{#if player?.isAdmin }
    {#if status === Status.Start}
        <button 
            class="primary"
            on:click={onStartClick}
            test-id="button-start-voting"
        > 
            Start Voting
        </button>
    {/if}

    {#if status === Status.InProgress}
        <button 
            class="primary"
            on:click={onRevealClick} 
            test-id="button-reveal-cards"
            disabled={isRevealDisabled(player)}
        > 
            Reveal Cards
        </button>
    {/if}
    
    {#if status === Status.Revealed}
        <button 
            class="primary"
            on:click={onResetClick}
        > 
            Start Voting 
        </button>
    {/if}    

{/if}


</room-controller>

<style>

    room-controller{
        display:         flex;
        justify-content: center;
        gap:             3rem;
    }

    button {
        min-width: 10rem;
        height:    4rem;
        padding:   0 0.5rem;

    }


</style>