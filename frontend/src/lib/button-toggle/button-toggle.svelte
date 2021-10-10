<script lang="ts">
    import {createEventDispatcher} from "svelte"
    const dispatch = createEventDispatcher()
    const keyActivate = "activate"

    export let labels: string[] = []
    export let activeIndex = -1;

</script>


<button-toggle>
    {#each labels as label, li }
        <button 
            class:active={li===activeIndex} 
            on:click={ () => dispatch(keyActivate, li)}
        >
            {label}
        </button>
    {/each}

</button-toggle>


<style>
    button-toggle {
        margin:  0;
        padding: 0;
        height:  52px
    }

    button {
        margin: 0;
        height: 100%;
    }
            
    button + button {
        border-left: none;
    }

    button.active{
        background: var(--color-red);
	    color:    black;
    }
</style>