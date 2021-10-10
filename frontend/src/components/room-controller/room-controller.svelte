<script lang="ts">
    import { api } from "../../api"
    import { PlayerType, Set, Status } from "../../api/state-store";
    import type { Player } from "../../api";
    import { ButtonToggle } from "$lib/button-toggle"
    import RoomLink from "./room-link.svelte";

    export let status = "";
    export let player: Player
    export let roomID = "";
    export let sets: Set[] = [];

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

    function onCardChange(event: Event){
        const selectEl = event.target as HTMLSelectElement
        const index = selectEl.value
        const cards = sets[index].cards
        api.setCards(cards)
    }

    function isRevealDisabled(player: Player): boolean {
        return player.type !== PlayerType.Spectator && player.chosenCard === "" 
    }

    $: typeIndex = playerTypes.indexOf(player?.type)
    $: roomLink = `${window.location.origin}/room/${roomID}`
    

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
    
    <select disabled={status === Status.InProgress} on:change={onCardChange}>
        {#each sets as set, si}
            <option value={si}>{set.label}: {set.cards}</option>
        {/each}
    </select>

    
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

    select{
        width:         200px;
        text-overflow: ellipsis;
    }
</style>