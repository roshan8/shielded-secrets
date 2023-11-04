<script lang="ts">
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		Radio,
		TableHeadCell,
		Button,
		Toast,
		GradientButton,
		Drawer,
		CloseButton,
		Label,
		Input,
		Textarea,
		Dropdown,
		Spinner
	} from 'flowbite-svelte';
	import { ChevronDownSolid, PlusOutline } from 'flowbite-svelte-icons';
	import { sineIn } from 'svelte/easing';
	import { onMount } from 'svelte';
	import axios from 'axios';

	const API_URL = 'http://localhost:9090';

	let displayToast = false;
	let ToastMessage = '';
	let loading = false;

	//  modal variables
	let newKey = '';
	let newValue = '';
	let transitionParams = {
		x: -320,
		duration: 200,
		easing: sineIn
	};
	let hideCreateSecretDrawer = true;

	let selectedRegion = '';
	let selectedSecret = '';

	let secretNameList: string[] = [];
	let secretData: string[] = [];
	let regions: string[] = [];

	async function getRegions() {
		try {
			loading = true;
			const response = await axios.get(`${API_URL}/api/regions`);
			regions = response.data.data;

			// choose the first region by default
			if (regions.length > 0) {
				selectedRegion = regions[0];
				listSecrets();
			}
		} catch (error) {
			console.error('error fetching secretes', error);
		} finally {
			loading = false;
		}
	}

	async function getSecret(secretName: string) {
		try {
			loading = true;
			const response = await axios.get(`${API_URL}/api/regions/${selectedRegion}/${secretName}`);
			secretData = response.data.data;
		} catch (error) {
			console.error('error fetching secretes', error);
			// TODO: Handle 403 for requests failing due to auth/allowed ip failures
		} finally {
			loading = false;
		}
	}

	async function listSecrets() {
		try {
			loading = true;
			const response = await axios.get(`${API_URL}/api/regions/${selectedRegion}/`);
			secretNameList = response.data.data;

			if (secretNameList.length > 0) {
				selectedSecret = '';
			} else {
				selectedSecret = '';
			}
		} catch (error) {
			console.error('error fetching secretes', error);
			selectedSecret = '';
			// TODO: Handle 403 for requests failing due to auth/allowed ip failures
		} finally {
			loading = false;
		}
	}

	async function deleteSecretKey(keyName: string) {
		try {
			loading = true;
			const response = await axios.delete(
				`${API_URL}/api/regions/${selectedRegion}/${selectedSecret}?key=${keyName}`
			);
			if (response.status == 200) {
				displayToast = true;
				ToastMessage = `Item(` + { keyName } + ` has been deleted.`;
				setTimeout(() => {
					displayToast = false;
				}, 3000);
				getSecret(selectedSecret);
			}
		} catch (error) {
			console.error('error fetching secretes', error);
		} finally {
			loading = false;
		}
	}

	async function addSecretKey() {
		try {
			if (selectedRegion == '' || selectedSecret == '') {
				ToastMessage = `Please select region and secret.`;
				displayToast = true;
				return;
			}

			const response = await axios.put(
				`${API_URL}/api/regions/${selectedRegion}/${selectedSecret}`,
				{
					key: newKey,
					value: newValue
				}
			);
			if (response.status == 200) {
				ToastMessage = `Item(` + newKey + `) has been added to the secret(` + selectedSecret + `).`;
				console.log(ToastMessage);
				displayToast = true;
				hideCreateSecretDrawer = true;
				setTimeout(() => {
					displayToast = false;
				}, 3000);
				getSecret(selectedSecret);
				return;
			}
		} catch (error) {
			// TODO: Handle non 2xx in the try block itself
			console.error('error adding secret key', error);
			ToastMessage = `Item(` + newKey + `) already exists in the secret(` + selectedSecret + `).`;
			displayToast = true;
			hideCreateSecretDrawer = true;
			setTimeout(() => {
				displayToast = false;
			}, 3000);
		}
	}

	function setRegion(regionName: string) {
		selectedRegion = regionName;
		secretData = [];
		listSecrets();
	}

	function setSecret(secretName: string) {
		selectedSecret = secretName;
		secretData = [];
		getSecret(secretName);
	}

	onMount(() => {
		getRegions();
		if (selectedRegion.length) {
			listSecrets();
		}
	});
</script>

<div>
	{#if displayToast}
		<Toast position="bottom-right" color="red">
			{ToastMessage}
		</Toast>
	{/if}

	{#if loading}
		<div
			class="fixed bg-black bg-opacity-50 h-full w-full inset-0 z-[999] flex justify-center items-center"
		>
			<Spinner class="w-10 h-10 text-white" />
		</div>
	{/if}

	<Drawer transitionType="fly" {transitionParams} bind:hidden={hideCreateSecretDrawer} id="sidebar">
		<div class="flex items-center">
			<h5
				id="drawer-label"
				class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400"
			>
				<PlusOutline class="w-4 h-4 mr-2.5 outline-none focus:outline-none" />Add new secret
			</h5>
			<CloseButton on:click={() => (hideCreateSecretDrawer = true)} class="mb-4 dark:text-white" />
		</div>
		<form on:submit|preventDefault={() => addSecretKey()} class="mb-6">
			<div class="mb-6">
				<Label for="Key" class="block mb-2">Key</Label>
				<Input
					id="secretKey"
					name="secretKey"
					bind:value={newKey}
					required
					placeholder="Enter secret key."
				/>
			</div>
			<div class="mb-6">
				<Label for="Value" class="mb-2">Value</Label>
				<Textarea
					id="secretValue"
					bind:value={newValue}
					placeholder="Enter secret value."
					rows="4"
					name="secretValue"
				/>
			</div>
			<Button on:click={() => addSecretKey()} type="submit" class="w-full">Save</Button>
		</form>
	</Drawer>

	<div class="flex gap-2 justify-between">
		<div class="flex gap-2">
			<Button on:click={() => getRegions()}>
				{selectedRegion.length ? selectedRegion : 'Select region'}
				<ChevronDownSolid
					class="w-3 h-3 ml-2 text-white dark:text-white outline-none focus:outline-none"
				/></Button
			>
			<Dropdown class="w-full p-3 space-y-3 text-sm">
				<div slot="header" class="px-4 py-2">
					<span class="block text-sm text-gray-900 dark:text-white"> Available regions </span>
					<span class="block truncate text-xs font-medium"> Select a region to view secrets </span>
				</div>
				<div class="space-y-2.5 max-h-[400px] overflow-y-scroll w-full">
					{#each regions as region}
						<li>
							<Radio
								name="regionName"
								on:change={() => setRegion(region)}
								on:click={() => setRegion(region)}
								checked={region == selectedRegion}
								value={region}>{region}</Radio
							>
						</li>
					{/each}
				</div>
			</Dropdown>

			<Button
				>{selectedSecret.length ? selectedSecret : 'Select secret'}<ChevronDownSolid
					class="w-3 h-3 ml-2 text-white dark:text-white outline-none focus:outline-none"
				/></Button
			>
			<Dropdown class="w-44 p-3 space-y-3 text-sm">
				<div slot="header" class="px-4 py-2">
					{#if secretNameList.length == 0}
						<span class="block text-sm text-gray-900 dark:text-white"> No secrets found </span>
						<span class="block truncate text-xs font-medium"> Add a secret to view </span>
					{/if}
					{#if secretNameList.length > 0}
						<span class="block text-sm text-gray-900 dark:text-white"> Available secrets </span>
						<span class="block truncate text-xs font-medium"> Select a secret to view keys </span>
					{/if}
				</div>
				<div class="space-y-2.5 max-h-[400px] overflow-y-scroll w-full">
					{#each secretNameList as secretName}
						<li>
							<Radio
								name="secretName"
								bind:group={secretName}
								on:change={() => setSecret(secretName)}
								on:click={() => setSecret(secretName)}
								checked={secretName == selectedSecret}
								value={secretName}>{secretName}</Radio
							>
						</li>
					{/each}
				</div>
			</Dropdown>
		</div>
		<div class="flex gap-2">
			<Button on:click={() => (hideCreateSecretDrawer = false)} outline pill>
				<PlusOutline class="w-3 h-3 mr-2.5 outline-none focus:outline-none" />
				Add new secret</Button
			>
		</div>
	</div>

	<div class="h-4" />

	<Table hoverable={true} shadow>
		<TableHead>
			<TableHeadCell>Key</TableHeadCell>
			<TableHeadCell>Value</TableHeadCell>
			<TableHeadCell>Action</TableHeadCell>
		</TableHead>
		<TableBody>
			{#each secretData as item}
				<TableBodyRow>
					<TableBodyCell>{item}</TableBodyCell>
					<TableBodyCell>**********</TableBodyCell>
					<TableBodyCell>
						<Button on:click={() => deleteSecretKey(item)} color="red" pill>Delete</Button>
					</TableBodyCell>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
</div>
