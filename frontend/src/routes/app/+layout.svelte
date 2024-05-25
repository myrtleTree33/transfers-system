<script type="ts">
  import IcRoundHome from '~icons/ic/round-home';
  import IcRoundSupervisorAccount from '~icons/ic/round-supervisor-account';
  import IcRoundPerson from '~icons/ic/round-person';
  import { PUBLIC_PRODUCT_NAME } from '$env/static/public';

  import IcRoundMenu from '~icons/ic/round-menu';
  import GrommetIconsGithub from '~icons/grommet-icons/github';
  import IcOutlineInsertDriveFile from '~icons/ic/outline-insert-drive-file';

  export let data;

  const docs = [
    {
      title: 'General',
      link: '',
      type: 'header',
    },
    {
      title: 'Github Repository',
      link: 'https://github.com/elishah-consulting/daytwo-js',
      type: 'github',
    },
    {
      title: 'General',
      link: '',
      type: 'header',
    },
    {
      title: 'Feature A',
      link: '/app/feature-a',
    },
    {
      title: 'Sample Markdown',
      link: '/app/sample-markdown',
    },
    {
      title: 'Settings',
      link: '',
      type: 'header',
    },
    {
      title: 'Settings',
      link: '/app/settings',
    },
    {
      title: 'Logout',
      link: '/logout',
    },
  ];
</script>

<svelte:head>
  <title>{PUBLIC_PRODUCT_NAME} App</title>
</svelte:head>

<div class="w-full min-h-screen flex">
  <!-- This is a side drawer navigation.  Feel free to use this widget should your app have more than 4 main pages. -->
  <nav class="p-4 hidden lg:inline">
    <ul class="menu bg-base-200 w-72 max-w-full rounded-box">
      <li class="menu-title text-primary font-bold mb-1">
        <div class="flex-1 flex flex-col gap-2">
          {data?.user?.email}
          <div class="badge badge-outline badge-sm font-normal">Member</div>
        </div>
      </li>

      {#each docs as doc}
        {#if doc?.type === 'header'}
          <li class="menu-title">
            {doc.title}
          </li>
        {:else if doc?.type === 'github'}
          <li>
            <a href={doc.link} class="text-accent font-bold">
              <GrommetIconsGithub class="inline text-lg" />
              {doc.title}
            </a>
          </li>
        {:else}
          <li>
            <a href={doc.link}>
              <IcOutlineInsertDriveFile class="inline text-lg" />
              {doc.title}
            </a>
          </li>
        {/if}
      {/each}
    </ul>
  </nav>

  <div class="drawer lg:hidden fixed z-[9999]">
    <input id="my-drawer" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content">
      <label
        for="my-drawer"
        class="btn btn-primary w-12 h-12 rounded-full m-4 fixed bottom-0 right-0"
        ><IcRoundMenu class="inline" /></label
      >
    </div>
    <div class="drawer-side">
      <label for="my-drawer" aria-label="close sidebar" class="drawer-everlay"></label>
      <ul class="menu p-4 w-80 max-w-full min-h-full bg-base-300">
        <!-- Sidebar content here -->
        {#each docs as doc}
          {#if doc?.type === 'header'}
            <li class="menu-title">
              {doc.title}
            </li>
          {:else if doc?.type === 'github'}
            <li class="text-accent">
              <a href={doc.link}>
                <GrommetIconsGithub class="inline text-lg" />
                {doc.title}
              </a>
            </li>
          {:else}
            <li>
              <a href={doc.link}>
                <IcOutlineInsertDriveFile class="inline text-lg" />
                {doc.title}
              </a>
            </li>
          {/if}
        {/each}
      </ul>
    </div>
  </div>

  <div class="flex-1">
    <slot />
    <div class="mb-20" />
  </div>

  <!-- This is a mobile bottom navigation.  Feel free to use this widget should your app have less than 4 main pages. -->

  <!-- <div class="btm-nav btm-nav-sm lg:hidden bg-base-200">
		<a href="/app">
			<IcRoundHome class="inline" />
			<span class="btm-nav-label">Home</span>
		</a>
		<a href="/app/feature-a">
			<IcRoundSupervisorAccount class="inline" />
			<span class="btm-nav-label">Feature A</span>
		</a>
		<a href="/app/settings">
			<IcRoundPerson class="inline" />
			<span class="btm-nav-label">Settings</span>
		</a>
	</div> -->
</div>
