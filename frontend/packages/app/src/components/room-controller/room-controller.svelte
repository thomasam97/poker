<script lang="ts">

    import { PlayerType, Set, Status, api } from "$lib/api";
    import type { Player } from "api";
    import { ButtonToggle } from "$lib/button-toggle"
    import RoomLink from "./room-link.svelte";
    import AutoReveal from "./auto-reveal.svelte";


    export let status = "";
    export let player: Player
    export let players: Player[] = []
    export let roomID = "";
    export let sets: Set[] = [];
    export let autoReveal: boolean = false;

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
    
    function onAdminChange(event: Event){
        const selectEl = event.target as HTMLSelectElement
        const index = selectEl.value
        const player = players[index]
        api.setAdmin(player)
    }

    function onAutoRevealChange(event: CustomEvent){
        const autoReveal = event.detail as boolean
        api.setAutoReveal(autoReveal)
    }

    function isRevealDisabled(player: Player): boolean {
        return player.type !== PlayerType.Spectator && player.chosenCard === "" 
    }

    $: typeIndex = playerTypes.indexOf(player?.type)
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
    
    <select disabled={status === Status.InProgress} on:change={onCardChange}>
        {#each sets as set, si}
            <option value={si}>{set.label}: {set.cards}</option>
        {/each}
    </select>

    <AutoReveal
        value={autoReveal}
        on:activate={onAutoRevealChange}
    />

    <label>
        <span class="give-admin-label">Give admin to</span>
        <select disabled={status === Status.InProgress} on:change={onAdminChange}>
        {#each players as player, pi}
            <option value={pi}>{player.name}</option>
        {/each}
    </select>
</label>
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
        gap:             3rem;
    }

    button {
        min-width: 10rem;
        height:    4rem;
        padding:   0 0.5rem;

    }

    select{
        width:         200px;
        text-overflow: ellipsis;
        font-size:     1rem;
        border-width:  0.25rem;
        height:        4rem;
    }

    .give-admin-label{
        letter-spacing: 0.1em;
        font-family:    var(--font-all-capital);
        text-transform: uppercase;
    }

</style>