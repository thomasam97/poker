<script lang="ts">
	import Header from '$lib/header/Header.svelte';
	import '../app.css';
	import {initConsoleDetection} from  "$lib/dev-tools/console-detector"
	import {initConsoleMessage} from "./console-message"
	

	initConsoleDetection()
	initConsoleMessage()


</script>

<!-- <Header /> -->

<main>
	<slot />
</main>



<style>
	main {
		flex: 1;
		display: flex;
		flex-direction:  column;
		justify-content: flex-start;
		padding: 1rem;
		box-sizing: border-box;
		height: 100vh;
	}

	
</style>
