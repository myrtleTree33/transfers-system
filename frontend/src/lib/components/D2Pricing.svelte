<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { twMerge } from 'tailwind-merge';

  const dispatch = createEventDispatcher();

  export let isSelected: boolean = false;
  export let selectedHeader: string = '';

  $: cardContainerClass = twMerge(
    'text-primary-content bg-primary p-8 pt-0 rounded-lg w-full lg:max-w-96',
    selectedHeader ? 'indicator border-2 border-neutral' : '',
    isSelected ? 'indicator border-2 border-success' : '',
    $$restProps?.class,
  );

  $: indicatorClass = twMerge(
    'indicator-item indicator-top indicator-center badge uppercase text-xs font-bold',
    selectedHeader ? 'badge-neutral' : '',
    isSelected ? 'badge badge-success' : '',
  );
</script>

<div class={cardContainerClass}>
  <section class="w-full">
    {#if selectedHeader}
      <span class={indicatorClass}>{selectedHeader}</span>
    {/if}
    <div class="text-2xl font-bold">
      <h3 class="text-primary-content"><slot name="header">Free</slot></h3>
    </div>
    <div class="flex gap-4 items-end">
      {#if $$slots?.usualPrice}
        <span class="text-sm line-through">
          <slot name="usualPrice" />
        </span>
      {/if}

      {#if $$slots?.currentPrice}
        <span class="text-4xl font-bold text-white">
          <slot name="currentPrice" />
        </span>
      {/if}

      {#if $$slots?.currency}
        <span class="text-sm">
          <slot name="currency" />
        </span>
      {/if}
    </div>
    <slot />
  </section>
</div>
