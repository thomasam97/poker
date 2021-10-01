<script>
    import Card from "./card.svelte";
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();


    let cards = [
        "1",
        "2",
        "3",
        "5",
        "8",
        "13",
        "20",
        "40",
        "100",
        "∞",
        "?"
    ]
</script>


<ol>
    {#each cards as card}
        <li>
            <Card label={card} on:click={ () => dispatch("choose", card)} />
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

