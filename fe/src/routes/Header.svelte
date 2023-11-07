<script>
	import { page } from '$app/stores';
	import { MoonSolid, SunSolid } from 'flowbite-svelte-icons';

	let theme = 'light';
	function toggle() {
		window.document.getElementById('app')?.classList.toggle('dark-mode');
		if (theme === 'light') {
			theme = 'dark';
		} else {
			theme = 'light';
		}
	}
</script>

<header class="relative w-full p-0">
	<nav class="flex justify-center">
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L1,2 C1.5,3 1.5,3 2,3 L2,0 Z" class="fill-[var(--background)]" />
		</svg>
		<ul
			class="relative p-0 m-0 h-[3em] flex justify-center items-center list-none bg-[var(--background)] bg-contain"
		>
			<li aria-current={$page.url.pathname === '/' ? 'page' : undefined} class="relative h-full">
				<a
					href="/"
					class="flex h-full items-center text-[color:var(--color-text)] font-bold text-[0.8rem] uppercase tracking-widest no-underline transition-[color] duration-[0.2s] ease-linear px-2 py-0"
					>Home</a
				>
			</li>
			<!-- TODO: Add this later -->
			<!-- <li
				aria-current={$page.url.pathname === '/about' ? 'page' : undefined}
				class="relative h-full"
			>
				<a
					href="/about"
					class="flex h-full items-center text-[color:var(--color-text)] font-bold text-[0.8rem] uppercase tracking-widest no-underline transition-[color] duration-[0.2s] ease-linear px-2 py-0"
					>About</a
				>
			</li> -->
		</ul>
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L0,3 C0.5,3 0.5,3 1,2 L2,0 Z" />
		</svg>
	</nav>
	<button on:click={toggle} class="absolute right-10 top-3">
		{#if theme === 'dark'}
			<SunSolid class="w-6 h-6 text-[#ff3e00]" />
		{:else}
			<MoonSolid class="w-6 h-6 text-[#ff3e00]" />
		{/if}
	</button>

	<div class="corner" />
</header>

<style>
	.corner {
		width: 3em;
		height: 3em;
	}

	nav {
		--background: rgba(255, 255, 255, 0.7);
	}

	svg {
		width: 2em;
		height: 3em;
		display: block;
	}

	path {
		fill: var(--background);
	}

	li[aria-current='page']::before {
		--size: 6px;
		content: '';
		width: 0;
		height: 0;
		position: absolute;
		top: 0;
		left: calc(50% - var(--size));
		border: var(--size) solid transparent;
		border-top: var(--size) solid var(--color-theme-1);
	}

	nav a:hover {
		color: var(--color-theme-1);
	}
</style>
