<script lang="ts">
  import { onMount } from 'svelte';

  onMount(() => {
    window.location.href = '/';
  });
</script>

<svelte:head>
  <title>Logout</title>
</svelte:head>

<p>Logging out...</p>
