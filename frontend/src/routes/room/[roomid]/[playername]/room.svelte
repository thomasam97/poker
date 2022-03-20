<script lang="ts">
import { page } from "$app/stores";
import { api, Cards as TypeCards, Set } from "../../../../api"
import { Player, PlayerType, Status } from "../../../../api/state-store";
import { Cards } from "../../../../components/cards";
import { PlayerList } from "../../../../components/player-list"
import { RoomController } from "../../../../components/room-controller";


const { 
    roomid:     roomID,
    playername: playerName,
} = $page.params


let player;
let players = []
let roomStatus: Status = Status.Init
let cards: TypeCards = []
let sets: Set[] = []
let autoReveal: boolean = false

api.register(roomID, playerName, (data) => {

    player = data.player
    players = data.players
    roomStatus = data.status
    cards = data.cards
    sets = data.sets
    autoReveal = data.autoReveal

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

function isPlayerSpectator( player: Player): boolean{
    return player && player.type === PlayerType.Spectator
}

</script>

<main>

    <RoomController
        status={roomStatus} 
        player={player} 
        players={players}
        roomID={roomID} 
        sets={sets} 
        autoReveal={autoReveal}    
    />



{#if isPlayerSpectator(player) || hasPlayerChosen(player) || !isVotingInProgress(roomStatus)}
    <PlayerList players={players} currentPlayer={player} isRevealed={isRevealed(roomStatus)} />
{/if}

{#if !isPlayerSpectator(player) && ( !hasPlayerChosen(player) && isGameRunnin(roomStatus) )}
    <Cards on:choose={(event) => onChoose(event.detail)} cards={cards} />
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