<script lang="ts">
import { page } from "$app/stores";
import { api } from "../../../../api"
import { Status } from "../../../../api/state-store";
import { Cards } from "../../../../components/cards";
import { PlayerList } from "../../../../components/player-list"
import type { Player } from "../../../../components/player-list/player";
import { RoomController } from "../../../../components/room-controller";
import { RoomStatus } from "../../../../components/room-status";
import Header from "../../../../lib/header/Header.svelte"


const { 
    roomid:     roomID,
    playername: playerName,
} = $page.params


let player;
let players = []
let roomStatus: Status = Status.Init

api.register(roomID, playerName, (data) => {
    console.debug('[DEBUG] ', {data} )

    player = data.player
    players = data.players
    roomStatus = data.status

})

function onChoose(card: string){
    api.choose(card)
}

function hasPlayerChosen(player: Player): boolean {
    if(!player){
        return false
    }

    return player.chosenCard !== ""
}

function isGameRunnin(roomStatus: Status): boolean {
    return roomStatus === Status.InProgress
}

function isVotingInProgress(roomStatus: Status): boolean {
    return roomStatus === Status.InProgress
}

function isRevealed(roomStatus: Status): boolean {
    return roomStatus === Status.Revealed
}

</script>

<main>

{#if player?.isAdmin}
    <RoomController status={roomStatus} player={player} />
{:else}
    <div class="ghost" />
{/if}


<!-- <RoomStatus status={roomStatus}/> -->

{#if hasPlayerChosen(player) || !isVotingInProgress(roomStatus)}
    <PlayerList players={players} isRevealed={isRevealed(roomStatus)} />
{/if}

{#if !hasPlayerChosen(player) && isGameRunnin(roomStatus)}
    <Cards on:choose={(event) => onChoose(event.detail) }/>
{/if}

</main>

<style>
    main {
        display: grid;
        gap: 2rem;
    }

    .ghost {
        height: 40px;
    }
</style>