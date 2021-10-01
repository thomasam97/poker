<script lang=ts>
    import type { Player } from "./player";
    import PlayerCard from "./player-card.svelte";

    export let players: Player[] = []
    export let isRevealed: boolean = false


</script>

<ol>
    {#each players as player}
        <li>
            <PlayerCard player={player} isRevealed={isRevealed} />
        </li>
    {/each}
</ol>


<style>
    ol{
        list-style-type: none;
        margin:          0;
        padding:         0;

        display:        flex;
        flex-direction: row;
        flex-wrap:      wrap;
        gap:            2rem;
    }
</style>