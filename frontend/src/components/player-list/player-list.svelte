<script lang=ts>
    import type { Player } from "./player";

    export let players: Player[] = []

</script>

<ul>
    {#each players as player}
        <li>{player.name} isAdmin: {player.isAdmin}</li>
    {/each}
</ul>