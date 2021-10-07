<script lang=ts>
    import { fade, } from 'svelte/transition';
    import { flip } from 'svelte/animate';
    import { cubicInOut } from 'svelte/easing';
    import { Player, PlayerType } from "../../api";
    import PlayerCard from "./player-card.svelte";

    export let players: Player[] = []
    export let currentPlayer: Player;
    export let isRevealed: boolean = false

    $: playingPlayer = players.filter(p => p.type === PlayerType.Player)

</script>

<ol>
    {#each playingPlayer as player (player.id)}
        <li 
            transition:fade|local 
	        animate:flip="{{duration: 200, easing: cubicInOut}}"
        >
            <PlayerCard player={player} isCurrentPlayer={currentPlayer.id === player.id} isRevealed={isRevealed} />
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