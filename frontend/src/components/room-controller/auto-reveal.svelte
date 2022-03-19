<script lang="ts">

import { ButtonToggle } from "$lib/button-toggle"
import {createEventDispatcher} from "svelte"
    export let value = false

    const dispatch = createEventDispatcher()
    const keyActivate = "activate"
    
    let mainEl: HTMLElement

    function onChange(){
        disabled = true
    }

    function onActivate(event: CustomEvent){
        const selectedIndex = event.detail as number
        const value = options[selectedIndex].value
        dispatch(keyActivate, value)
    }

    $: disabled = false && value
    $: typeIndex = value ? 1 : 0;

    const options = [
        {
            label: "Off",
            value: false,
        },
        {
            label: "On",
            value: true,
        },
    ]
</script>

<main bind:this={mainEl}>
    <!-- <label> -->
        <span>Auto reveal</span>

        <ButtonToggle 
            labels={options.map(o=>o.label)} 
            on:activate={onActivate} 
            activeIndex={typeIndex}
        />
        <!-- <input 
            type="checkbox" 
            on:change 
            on:change={onChange} 
            checked={value} 
            disabled={disabled}
        /> -->
    <!-- </label> -->
</main>

<style>
    main{
        display:        flex;
        flex-direction: row;
        align-items:    center;
        /* border:         2px solid var(--color-red); */
        padding:        0;
        color:          white;
    }
    
    span {
        padding: 0 0.5rem;
    }
</style>