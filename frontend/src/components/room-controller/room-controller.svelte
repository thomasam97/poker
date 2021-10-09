<script lang="ts">
    import { api } from "../../api"
    import { PlayerType, Status } from "../../api/state-store";
    import type { Player } from "../../api";
    import { ButtonToggle } from "$lib/button-toggle"
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


    function onActivate(event: CustomEvent<number>){
        const typeIndex = event.detail;
        const type = playerTypes[typeIndex]
        api.setPlayerType(type)
    }

    function isRevealDisabled(player: Player): boolean {
        return player.type !== PlayerType.Spectator && player.chosenCard === "" 
    }

    $: typeIndex = playerTypes.indexOf(player?.type)

    $: roomLink = `${window.location.origin}/room/${roomID}`
    function onRoomLinkClick(event: MouseEvent){
        event.preventDefault()
        navigator.clipboard.writeText(roomLink);
    }

</script>

<room-controller>
    <RoomLink link={roomLink} />

{#if player?.isAdmin }
    {#if status === Status.Start}
        <button on:click={onStartClick}> Start Voting</button>
    {/if}

    {#if status === Status.InProgress}
        <button on:click={onRevealClick} disabled={isRevealDisabled(player)}> Reveal Cards</button>
    {/if}
    
    {#if status === Status.Revealed}
        <button on:click={onResetClick}> Start Voting </button>
    {/if}
{/if}

    <ButtonToggle 
        labels={playerTypes} 
        on:activate={onActivate} 
        activeIndex={typeIndex}
    />
</room-controller>

<style>

    room-controller{
        display:         flex;
        justify-content: center;
        gap:             1rem;
    }

    button {
        width: 10rem;
    }
</style>