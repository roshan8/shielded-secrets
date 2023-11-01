<script>
	import '../app.postcss';
	import Header from './Header.svelte';
	import './styles.css';
</script>

<div class="app" id="app">
	<Header />

	<main>
		<slot />
	</main>
</div>

<style>
	.app {
		min-height: 100vh;
		height: auto;
	}
	:global(.app) {
		background-color: #f5eeee;
		transition: background-color 0.3s;
	}
	:global(.app.dark-mode) {
		background-color: #1d3040;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 1rem;
		width: 100%;
		max-width: 64rem;
		margin: 0 auto;
		box-sizing: border-box;
	}
</style>
