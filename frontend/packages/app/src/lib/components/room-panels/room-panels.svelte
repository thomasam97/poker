<script>
    import { createEventDispatcher } from 'svelte';
    
    // 
    // Open Settings
    // 
    export let open = true;
    const dispatch = createEventDispatcher();
    function dispatchOpenSettings(){
        dispatch('opensettings')
    }


</script>

<room-panels>

    <div class="controls">
        <slot name="controls"></slot>
    </div>

    <div class="settings" class:open>
        <slot name="settings"></slot>
    </div>

    <div class="panel-control">
        <button class="text" on:click={dispatchOpenSettings}>
            Settings ⬇
        </button>
    </div>

</room-panels>

<style>
    room-panels {
        --height: 4rem;


        display:  flex;
        justify-content: center;
        position: relative;
        width:    auto;
        height:   var(--height);
        z-index:  200;
        
    }

    .settings {
        position:       absolute;
        width:          100%;
        transition:     all ease-in-out 0.5s;
        height:         auto;
        transform:      translateY(-200%);
        z-index:        100;
    }

    .settings.open {
        transform: translateY(0)
    }

    .controls {
        position: relative;
    }

    .panel-control{
        position:        relative;
        flex-grow:       1;
        justify-content: end;
        justify-items:   end;
        display:         flex;
    }

</style>