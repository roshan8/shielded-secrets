<script>
   import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, 
    TableHeadCell, Checkbox, TableSearch, Button, Toast } from 'flowbite-svelte';
   import { CheckCircleSolid, ExclamationCircleSolid, FireOutline, CloseCircleSolid } from 'flowbite-svelte-icons';
   import { onMount } from "svelte";
   import axios from "axios";

   let displayToast = false;
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
      tableData = response.data;
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
</div>

<style>
</style>