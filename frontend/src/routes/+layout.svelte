<script lang="ts">
  /**
   * This is the layout component for the entire application.
   * It is used to wrap all the pages in the application.
   * It contains the navigation bar, the main content, and the footer.
   */
  import { page } from '$app/stores';
  import { PUBLIC_PRODUCT_NAME } from '$env/static/public';
  import D2Toast from '$lib/components/D2Toast.svelte';
  import { setToast, toast_store } from '$lib/stores/stores';
  import '../app.css';

  export let data;

  // Check if the current page is a splash page
  let is_splash: boolean = false;
  $: is_splash =
    !$page.url.pathname.startsWith('/app') &&
    !$page.url.pathname.startsWith('/login') &&
    !$page.url.pathname.startsWith('/reset') &&
    !$page.url.pathname.startsWith('/reset-password') &&
    !$page.url.pathname.startsWith('/signup') &&
    !$page.url.pathname.startsWith('/logout');

  // This is the JWT access token exposed for use on the client-side, should API calls be necessary.
  const access_token: string = data?.access_token;

  // The toast message can be displayed on the client-side, by passing in the ?m= query parameter in the URL.
  setToast(data?.toast_message || '');
</script>

<svelte:head>
  <title>{PUBLIC_PRODUCT_NAME}: App</title>
</svelte:head>

<!-- Show toast message if available -->
<D2Toast />

<!-- Splash page navigation bar -->
{#if is_splash}
  <div class="fixed w-full z-[999]">
    <nav class="navbar bg-base-100 shadow py-2 lg:p-4">
      <div class="flex-1">
        <a class="btn btn-ghost text-xl" href="/">{PUBLIC_PRODUCT_NAME}</a>
        <ul class="flex-1 flex">
          <li>
            <a
              href="#pricing"
              class="btn btn-ghost"
              on:click={() => {
                document.getElementById('#pricing').scrollIntoView({ behavior: 'smooth' });
              }}>Pricing</a
            >
          </li>
        </ul>
      </div>
      <div class="flex-none">
        {#if access_token}
          <a href="/app" class="btn btn-primary text-md"> Dashboard</a>
        {:else}
          <div class="flex-1 flex gap-4">
            <a href="/login" class="btn btn-outline btn-sm lg:btn-md text-md">Login</a>
          </div>
        {/if}
      </div>
    </nav>
  </div>
{/if}

<!-- Main content -->
<slot />

<!-- Splash page footer -->
{#if is_splash}
  <footer
    class="py-24 bg-gray-900 text-primary-content w-full flex flex-col items-center text-center"
  >
    <div class="w-full flex flex-col items-center justify-center">
      <div class="w-full max-w-6xl px-8 text-sm">
        <div class="divider" />
        <div class="flex-1 grid lg:grid-cols-3 gap-8 text-center lg:text-left">
          <div class="flex flex-col gap-1 w-full">
            <div class="text-white text-lg mb-4">
              <a href="/">{PUBLIC_PRODUCT_NAME}</a>
            </div>
            <div class="">
              <a href="https://www.daytwo.io" class="btn btn-xs btn-outline btn-success mb-2"
                >Built with {PUBLIC_PRODUCT_NAME}</a
              >
            </div>
            <div class="text-xs">Copyright © 2024. All rights reserved.</div>
          </div>

          <div class="flex flex-col gap-1 w-full">
            <div class="text-white text-md mb-4">Product</div>
            <a href="#" class="hover:text-white">{PUBLIC_PRODUCT_NAME} for Indie developers</a>
            <a href="#" class="hover:text-white">{PUBLIC_PRODUCT_NAME} for side hustles</a>
            <a href="#" class="hover:text-white">{PUBLIC_PRODUCT_NAME} for startups</a>
          </div>

          <div class="flex flex-col gap-1 w-full">
            <div class="text-white text-md mb-4">Support</div>
            <a href="#" class="hover:text-white">Product</a>
            <a href="#" class="hover:text-white">Pricing</a>
            <a href="#" class="hover:text-white">Contact us</a>
          </div>
        </div>
      </div>
    </div>
  </footer>
{/if}
