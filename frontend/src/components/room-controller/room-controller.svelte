<script lang="ts">
    import { api } from "../../api"
    import { Status } from "../../api/state-store";

    export let status = "";

    function onStartClick(){
        api.startRoom()
    }

    function onRevealClick(){
        api.revealRoom()
    }

    function onResetClick(){
        api.resetRoom()
    }

</script>

<div class="room-controller-root">
{#if status === Status.Start}
    <button on:click={onStartClick}> Start </button>
{/if}

{#if status === Status.InProgress}
    <button on:click={onRevealClick}> Reveal </button>
{/if}
 
{#if status === Status.Revealed}
    <button on:click={onResetClick}> Reset </button>
{/if}
</div>

<style>
    button {
        width: 5rem;
    }
</style>