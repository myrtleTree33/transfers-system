<script>
  import { page } from '$app/stores';
</script>

<div class="w-full min-h-screen flex items-center justify-center p-4">
  {#if $page?.status === 404}
    <div>
      <h2>Whoops! This is not the page you are looking for.</h2>
      <p class="text-center">
        Head back to the <a href="/" class="hyperlink">home page</a> or
        <a href="/app" class="hyperlink">dashboard</a>.
      </p>
    </div>
  {:else}
    <div>
      <h1>{$page.status}: {$page.error.message}</h1>
    </div>
  {/if}
</div>
