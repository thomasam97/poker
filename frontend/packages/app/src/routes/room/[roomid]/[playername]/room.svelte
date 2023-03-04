<script lang="ts">
	import { page } from "$app/stores"; 
	import { api, stash, type Set, Status, type Player, PlayerType, type State } from "$lib/api"
	import { Cards } from "$lib/components/cards"
  	import type { Card as TypeCard, } from "$lib/components/cards";
	import { PlayerList } from "$lib/components/player-list"
	import { RoomController } from "$lib/components/room-controller";
	import { RoomPanels } from "$lib/components/room-panels";
	import { RoomSettings } from "$lib/components/room-settings";
	import { CardBacks } from "$lib/card-backs";
	
	const { 
		roomid:     roomID,
		playername: playerName,
	} = $page.params

	
	let player: Player | undefined;
	let players: Player[] | undefined = []
	let roomStatus: Status | undefined = Status.Init
	let cards: TypeCard[] = []
	let sets: Set[] = []
	let autoReveal: boolean = false

	// 
	// API Connection
	// 
	api.register(roomID, playerName, onMessage, onOpen)

	function onMessage(data: State){
		player = data.player
		players = data.players
		roomStatus = data.status
		cards = data.cards.map( card => ({label: card, value:""}))
		sets = data.sets
		autoReveal = data.autoReveal
	}
	function onOpen(){
		let cardBack = stash.CardBack()

		const isSupportedCardBack = Object.values(CardBacks).indexOf(cardBack) >= 0
		if(!isSupportedCardBack){
			cardBack = CardBacks.logo
			stash.storeCardBack(cardBack)
		}

		api.switchCardBack(cardBack)
	}
	
	

	// ---
	
	function onChoose(card: TypeCard){
		api.choose(card.label)
	}
	
	function onRevote(){
		if(!isVotingInProgress(roomStatus)){
			return
		}
		api.reVote()
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

	// 
	// RoomPanels
	// 
	let isSettingsOpen = false;
	function handleDone(){
		isSettingsOpen = false
	}
	function handleOpenSettings(){
		isSettingsOpen = true
	}
	
</script>

<svelte:head>
	<title>Poker: Room {roomID}</title>
</svelte:head>

<room>
	
	
	
	<RoomPanels open={isSettingsOpen} on:opensettings={handleOpenSettings}>
		<RoomSettings 
			slot="settings"
			player={player}
			players={players}
			status={roomStatus}
			autoReveal={autoReveal}
			sets={sets}

			on:done={handleDone}
		/>

		<RoomController
			slot="controls"
			status={roomStatus} 
			player={player} 
			roomID={roomID} 
			
		/>
		
	</RoomPanels>
	
	
	{#if isPlayerSpectator(player) || hasPlayerChosen(player) || !isVotingInProgress(roomStatus)}
		<PlayerList 
			players={players} 
			currentPlayer={player} 
			isRevealed={isRevealed(roomStatus)}
			on:revote={onRevote}
		/>
	{/if}
	
	{#if !isPlayerSpectator(player) && ( !hasPlayerChosen(player) && isGameRunnin(roomStatus) )}
		<Cards 
			cards={cards} 
			on:choose={(event) => onChoose(event.detail)} 
		/>
	{/if}
	
	<footer>
		<span> &nbsp; </span>
		<a href="https://www.sprinteins.com" class="logo" target="_blank">
			<img src="/img/logo_white.svg" class="logo" alt="Scrum-Poker Logo"/>
		</a> 
		
		<a href="https://www.sprinteins.com" target="_blank" class="se-logo">
			<img src="/img/se_logo_white.png" alt="SprintEins Logo"/>
		</a> 
	</footer>
	
	
</room>

<style>
	room {
		height:             100%;
		display:            grid;
		gap:                2rem;
		grid-template-rows: auto 1fr auto;
	}
	
	footer {
		display:     grid;
		align-items: baseline;
		height:      190px;
		padding:     0;
		grid-template-columns: 1fr 1fr 200px;
	}
	
	
	a.logo{
		width:      200px;
		height:     inherit;
		display:    block;
		text-align: center;
		padding:    0;
	}
	
	a.logo img {
		width:           200px;
		height:          inherit;
		object-fit:      none;
		object-position: 50% 71%;
	}
	
	.se-logo{
		width: 200px;
	}
	
	.se-logo img{
		width: inherit;
	}
	
	
</style>