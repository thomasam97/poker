<script lang="ts">
    import type { Player } from "./player";

    export let player: Player;
    export let isRevealed: boolean = false;
    $: playerHasChoosen = player.chosenCard !== ""
</script>

<div class="player-card">
    <div class="admin">
        {#if player.isAdmin}
            👑
        {/if}
    </div>
    <div class="player">
        {player.name}
        
    </div>
    <div class="card">
        <div class="content">
            {#if playerHasChoosen && isRevealed}
                {player.chosenCard}
            {/if}
            {#if playerHasChoosen && !isRevealed}
                ☑️
            {/if}
        </div>
    </div>
</div>

<style>
    .player-card {
        font-family: "Helvetica Neue";
    }
     .card {
        border: gray thin solid;
        width:  12rem;
        height: 20rem;
        
        display:     grid;
        place-items: center;
        
        background:    black;
        color:         white;
        cursor:        pointer; 
        border-radius: 4px;
        transition:    all 0.3s ease;
    }

    .content {
        font-size:   5rem;
        font-weight: bold;
        position:    relative;
        z-index:     1;
    }

    .player{
        width:      12rem;
        text-align: center;
        font-size:  2rem;
    }
    .admin{
        width:      12rem;
        text-align: center;
        font-size:  3rem;
        height:     3rem;
    }
</style>