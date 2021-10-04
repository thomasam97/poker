<script lang=ts>
    import { fade, } from 'svelte/transition';
    import { flip } from 'svelte/animate';
    import { cubicInOut } from 'svelte/easing';
    import type { Player } from "./player";
    import PlayerCard from "./player-card.svelte";

    export let players: Player[] = []
    export let isRevealed: boolean = false


</script>

<ol>
    {#each players as player (player.id)}
        <li 
            transition:fade|local 
	        animate:flip="{{duration: 200, easing: cubicInOut}}"
        >
            <PlayerCard player={player} isRevealed={isRevealed} />
        </li>
    {/each}
</ol>


<style>
    ol{
        list-style-type: none;
        margin:          0;
        padding:         0;

        display:         flex;
        flex-direction:  row;
        flex-wrap:       wrap;
        justify-content: center;
        gap:             2rem;
        transition:      all 0.3s ease;
    }
</style>