<script lang="ts">
	import Header from '$lib/header/Header.svelte';
	import '../app.css';
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
		/* width: 100%; */
		/* max-width: 1024px; */
		/* margin: 0 auto; */
		box-sizing: border-box;
	}

	


	@media (min-width: 480px) {
		footer {
			padding: 40px 0;
		}
	}
</style>
