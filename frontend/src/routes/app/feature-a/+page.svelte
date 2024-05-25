<script lang="ts">
</script>

<svelte:head>
  <title>Feature A</title>
</svelte:head>

<div>
  <h1>Feature A</h1>
  <p>This is an example of a page. You can being coding here.</p>

  <p>There are 4 different types of pages in SvelteKit. They are:</p>
  <ul class="list-disc list-inside">
    <li>{'+page.svelte'}: Run client-side. A page that is part of a route.</li>
    <li>{'+layout.svelte'}: Run client-side. A layout that is part of a route.</li>
    <li>
      {'+page.server.ts'}: Run server-side. A page that is part of a route and is used to fetch data
      on the server.
    </li>
    <li>
      {'+layout.server.ts'}: Run server-side. A layout that is part of a route and is used to fetch
      data on the server.
    </li>
    <li>
      {'+error.svelte'}: Run in the event of an error. A page that is part of a route and is used to
      display an error.
    </li>
  </ul>
</div>
