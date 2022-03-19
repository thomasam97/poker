<script lang="ts">
    export let value = false
    
    let mainEl: HTMLElement

    function onChange(){
        disabled = true
    }
    $: disabled = false && value
</script>

<main bind:this={mainEl}>
    <label>
        Auto reveal    
        <input 
            type="checkbox" 
            on:change 
            on:change={onChange} 
            checked={value} 
            disabled={disabled}
        />
    </label>
</main>

<style>
    main{
        display:        flex;
        flex-direction: row;
        align-items:    center;
        border:         2px solid var(--color-red);
        padding:        0.5rem;
        color:          white;
    }
</style>