<script>
   import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, Radio,
    TableHeadCell, Button, Toast, GradientButton,
    Drawer, CloseButton, Label, Input, Textarea, Dropdown } from 'flowbite-svelte';
   import { CloseCircleSolid, ChevronDownSolid, InfoCircleSolid } from 'flowbite-svelte-icons';
  import { sineIn } from 'svelte/easing';
   import { onMount } from "svelte";
   import axios from "axios";

   let displayToast = false;
   let ToastMessage = "";

  //  modal variables
  let newKey = ""
  let newValue = ""
  let transitionParams = {
    x: -320,
    duration: 200,
    easing: sineIn
  };
  let hideCreateSecretDrawer = true;

   let selectedRegion = "Select region";
   let selectedSecret = "Select secret";

   let secretNameList = [];
   let secretData = [];
   let regions = [];
  

  async function getRegions() {
    try {
      const response = await axios.get('http://localhost:9090/api/regions');
      regions = response.data.data;
    } catch (error) {
      console.error("error fetching secretes", error)
    }
  }

  async function getSecret(secretName) {
    try {
      const response = await axios.get(`http://localhost:9090/api/${selectedRegion}/${secretName}`);
      secretData = response.data.data;
    } catch (error) {
      console.error("error fetching secretes", error)
    }
  }

  async function listSecrets() {
    try {
      const response = await axios.get(`http://localhost:9090/api/${selectedRegion}/`);
      secretNameList = response.data.data;

      if (secretNameList.length > 0) {
        selectedSecret = "Select secret"
      } else {
        selectedSecret = "no secrets found"
      }

    } catch (error) {
      console.error("error fetching secretes", error)
      selectedSecret = "no secrets found"
    }
  }

  async function deleteSecretKey(keyName) {
    try {
      const response = await axios.delete(`http://localhost:9090/api/${selectedRegion}/${selectedSecret}?key=${keyName}`);
      if (response.status == 200) {
        displayToast = true;
        ToastMessage = `Item(`+{keyName}+` has been deleted.`
        setTimeout(() => {
          displayToast = false;
        }, 3000);
        getSecret(selectedSecret);
      }
    } catch (error) {
      console.error("error fetching secretes", error)
    }
  }

  async function adddSecretKey(){
    try {
      if (selectedRegion == "Select region" || selectedSecret == "Select secret") {
        ToastMessage = `Please select region and secret.`
        displayToast = true;
        return
      }

      const response = await axios.put(`http://localhost:9090/api/${selectedRegion}/${selectedSecret}`, {
        key: newKey,
        value: newValue
      });
      if (response.status == 200) {
        ToastMessage = `Item(`+ newKey +`) has been added to the secret(`+ selectedSecret +`).`
        console.log(ToastMessage)
        displayToast = true;
        hideCreateSecretDrawer = true;
        setTimeout(() => {
          displayToast = false;
        }, 3000);
        getSecret(selectedSecret);   
        return
      }
    } catch (error) {
      // TODO: Handle non 2xx in the try block itself
      console.error("error adding secret key", error)
      ToastMessage = `Item(`+ newKey +`) already exists in the secret(`+ selectedSecret +`).`
        displayToast = true;
        hideCreateSecretDrawer = true;
        setTimeout(() => {
          displayToast = false;
        }, 3000);
    }
  }

  function setRegion(regionName) {
    selectedRegion = regionName;
    secretData = [];
    listSecrets();
  }

  function setSecret(secretName) {
    selectedSecret = secretName;
    secretData = [];
    getSecret(secretName);
  }

  onMount(() => {
    getRegions()
    listSecrets();
  });

 </script>


<div>

  {#if displayToast}
    <Toast position="bottom-right" color="red">
      {ToastMessage}
    </Toast>
  {/if}

  <Drawer transitionType="fly" {transitionParams} bind:hidden={hideCreateSecretDrawer} id="sidebar3">
    <div class="flex items-center">
      <h5 id="drawer-label" class="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase dark:text-gray-400">
        <InfoCircleSolid class="w-4 h-4 mr-2.5" />Contact us
      </h5>
      <CloseButton on:click={() => (hideCreateSecretDrawer = true)} class="mb-4 dark:text-white" />
    </div>
    <!-- <form action="#" class="mb-6"> -->
    <form on:submit|preventDefault={() => adddSecretKey()} class="mb-6">
      <div class="mb-6">
        <Label for="Key" class="block mb-2">Key</Label>
        <Input id="secretKey" name="secretKey" bind:value={newKey} required placeholder="Enter secret key." />
      </div>
      <div class="mb-6">
        <Label for="Value" class="mb-2">Value</Label>
        <Textarea id="secretValue" bind:value={newValue} placeholder="Enter secret value." rows="4" name="secretValue" />
      </div>
      <Button on:click={() => adddSecretKey()} type="submit" class="w-full">Save</Button>  
    </form>
  </Drawer>

  <Button on:click={() => getRegions()}> {selectedRegion} <ChevronDownSolid class="w-3 h-3 ml-2 text-white dark:text-white" /></Button>
  <Dropdown class="w-44 p-3 space-y-3 text-sm">
    {#each regions as region}
      <li>
        <Radio name="regionName" on:change={() => setRegion(region)} on:click={() => setRegion(region)} value={region}>{region}</Radio>
      </li>
    {/each}
  </Dropdown>

  <Button>{selectedSecret}<ChevronDownSolid class="w-3 h-3 ml-2 text-white dark:text-white" /></Button>
  <Dropdown class="w-44 p-3 space-y-3 text-sm">
    {#each secretNameList as secretName}
      <li>
        <Radio name="secretName" bind:group={secretName} on:change={() => setSecret(secretName)} on:click={() => setSecret(secretName)} value={secretName}>{secretName}</Radio>
      </li>
    {/each}
  </Dropdown>

  <div class="h-4"></div>
  
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

  <GradientButton on:click={() => (hideCreateSecretDrawer = false)} outline pill color="purpleToBlue">Add secret</GradientButton>
</div>

<style>
</style>