<script lang="ts">
    import type { Player } from "../../api";
    import CardBG from "./card-bg.png"

    export let player: Player;
    export let isCurrentPlayer: boolean;
    export let isRevealed: boolean = false;
    $: playerHasChoosen = player.chosenCard !== ""
</script>

<div 
    class="player-card" 
    class:playerHasChoosen
    on:click
>
    <div class="admin">
        {#if player.isAdmin}
            👑
        {/if}
    </div>
    <div class="player" test-id="player-card__player-name">
        {player.name}
    </div>
    <div 
        class="card flip-card" 
        class:revealed={isRevealed} 
    >
        <div class="flip-card-inner">
            <div 
                class="flip-card-front" 
                test-id="player-card__front"
                class:isCurrentPlayer 
            >
                <img src={CardBG} class="card-image" alt="SprintEins Logo small"/>
            </div>
            <div 
                class="flip-card-back"  
                test-id="player-card__back"
                class:isCurrentPlayer
            >
                <div class="content">
                    {player.chosenCard}
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .isCurrentPlayer{
        --shadow: inset 0px 0px 3px 1px var(--color-red);
        -webkit-box-shadow: var(--shadow);
	    box-shadow: 		var(--shadow);
    }
    .player-card {
        font-family: "Helvetica Neue";
        transition:   all 0.3s ease-in-out;
    }
     .card {
        width:  12rem;
        height: 20rem;

        color:         white;
        cursor:        pointer; 
        transition:    all 0.3s ease;
    }

    .player-card:not(.playerHasChoosen){
        opacity: 50%;
        transform: scale(0.9, 0.9);
    }

    .content {
        font-size:      5rem;
        font-weight:    bold;
        position:       relative;
        z-index:        1;
        width:          100%;
        height:         100%;
        display:        grid;
        place-items:    center;
        text-transform: uppercase;
        
    }

    .player{
        width:      12rem;
        text-align: center;
        font-size:  2rem;
        color:      white;
    }
    .admin{
        width:      12rem;
        text-align: center;
        font-size:  3rem;
        height:     3rem;
    }

/* ------- */


    .flip-card {
        border-radius:    4px;
        background-color: transparent;
        width:            12rem;
        height:           20rem;
        perspective:      1000px;
        display:          grid;
        place-items:      center;
    }

    .flip-card-inner {
        position:        relative;
        width:           100%;
        height:          100%;
        text-align:      center;
        transition:      transform 0.6s;
        transform-style: preserve-3d;
        /* box-shadow:      0 4px 8px 0 rgba(0,0,0,0.2); */
    }

    .flip-card.revealed .flip-card-inner{
    /* .flip-card:hover .flip-card-inner { */
        transform: rotateY(180deg);
    }

    .flip-card-front, .flip-card-back {
        position:      absolute;
        width:         100%;
        height:        100%;
        border-radius: 4px;
        display:       grid;
        place-items:   center;

        -webkit-backface-visibility: hidden;
        backface-visibility:         hidden;
    }

    .card-image {
        width: 11rem;
    }

    .flip-card-front {
        /* background-color: #bbb; */
        /* background-image: "url(./card-bg.png)"; */
        background-color: black;
        color: black;
    }

    .flip-card-back {
        /* background-color: #2980b9; */
        background-color: black;
        color: white;
        transform: rotateY(180deg);
    }

</style>