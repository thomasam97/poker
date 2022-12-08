<script lang="ts">
    import Card from "./card.svelte";
    import { createEventDispatcher } from 'svelte';
    import type{ Card as TypeCard } from "./card"

    const dispatch = createEventDispatcher();

    export let disableHoverAnimation = false
    export let cards: TypeCard[] = []
</script>


<ol>
    {#each cards as card}
        <li>
            <Card 
                label={card.label} 
                bgImagePath={card.bgImagePath} 
                disableHoverAnimation={disableHoverAnimation}
                on:click={ () => dispatch("choose", card)} />
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
        gap:             2rem;
        justify-content: center;
    }

   

</style>

