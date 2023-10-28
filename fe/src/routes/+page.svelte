<script>
   import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, Radio,
    TableHeadCell, Checkbox, TableSearch, Button, Toast, GradientButton,
    Dropdown, DropdownItem, DropdownDivider, DropdownHeader, Search } from 'flowbite-svelte';
   import { CheckCircleSolid, ExclamationCircleSolid, FireOutline, CloseCircleSolid, ChevronDownSolid, UserRemoveSolid } from 'flowbite-svelte-icons';
   import { onMount } from "svelte";
   import axios from "axios";

   let displayToast = false;
   let regionName = "";
   let secretName = "";
   let deletedItem = "";
   let tableData = [];

  //  export let tableData = [
  //   { id: 1, name: "Apple MacBook Pro 17", description: "Laptop" },
  //   { id: 2, name: "Microsoft Surface Pro", description: "Tab" },
  //   { id: 3, name: "Magic Mouse 2", description: "Accessary" }
  // ];

  async function getSecrets() {
    try {
      const response = await axios.get('http://localhost:8080/api/secrets');
      tableData = response.data.data;
    } catch (error) {
      console.error("error fetching secretes", error)
    }
  }

  function deleteSecretItem(item) {
      displayToast = true;
      deletedItem = item;
      setTimeout(() => {
        displayToast = false;
      }, 3000);
  }

  onMount(() => {
    getSecrets();
  });

 </script>


<div>

  {#if displayToast}
    <!-- <Toast position="bottom-right" color="green">
      <svelte:fragment slot="icon">
        <CheckCircleSolid class="w-5 h-5" />
        <span class="sr-only">Check icon</span>
      </svelte:fragment>
      Item moved successfully.
    </Toast> -->
  
    <Toast position="bottom-right" color="red">
      <svelte:fragment slot="icon">
        <CloseCircleSolid class="w-5 h-5" />
        <span class="sr-only">Error icon</span>
      </svelte:fragment>
      Item({deletedItem}) has been deleted.
    </Toast>
  {/if}

  <Button>Select region<ChevronDownSolid class="w-3 h-3 ml-2 text-white dark:text-white" /></Button>
  <Dropdown class="w-44 p-3 space-y-3 text-sm">
    <li>
      <Radio name="regionName" bind:group={regionName} value={"us-east-1"}>us-east-1</Radio>
    </li>
    <li>
      <Radio name="regionName" bind:group={regionName} value={"us-east-2"}>us-east-2</Radio>
    </li>
    <li>
      <Radio name="regionName" bind:group={regionName} value={"us-east-3"}>us-east-3</Radio>
    </li>
  </Dropdown>

  <Button>Select secret<ChevronDownSolid class="w-3 h-3 ml-2 text-white dark:text-white" /></Button>
  <Dropdown class="w-44 p-3 space-y-3 text-sm">
    <li>
      <Radio name="secretName" bind:group={secretName} value={"ingester"}>ingester</Radio>
    </li>
    <li>
      <Radio name="secretName" bind:group={secretName} value={"postal"}>postal</Radio>
    </li>
    <li>
      <Radio name="secretName" bind:group={secretName} value={"nile"}>nile</Radio>
    </li>
  </Dropdown>

<!-- add few line spaces -->
  <div class="h-4"></div>
  
  <!-- TODO: Make it searchable -->
  <Table hoverable={true} shadow>
    <TableHead>
      <TableHeadCell>ID</TableHeadCell>
      <TableHeadCell>Product name</TableHeadCell>
      <TableHeadCell>description</TableHeadCell>
      <TableHeadCell>Action</TableHeadCell>
    </TableHead>
    <TableBody>
      {#each tableData as item}
      <TableBodyRow>
        <TableBodyCell>{item.id}</TableBodyCell>
        <TableBodyCell>{item.name}</TableBodyCell>
        <TableBodyCell>{item.description}</TableBodyCell>
        <TableBodyCell>
          <Button on:click={() => deleteSecretItem(item.id)} color="red" pill>Delete</Button>
        </TableBodyCell>
      </TableBodyRow>
      {/each}
    </TableBody>
  </Table>

  <GradientButton outline pill color="purpleToBlue">Add secret</GradientButton>
</div>

<style>
</style>