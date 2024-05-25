<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { twMerge } from 'tailwind-merge';
  import IcRoundLock from '~icons/ic/round-lock';
  import IcRoundWarningAmber from '~icons/ic/round-warning-amber';

  const dispatch = createEventDispatcher();

  export let errorText: string = '';
  export let is_display_label: boolean = false;

  let combinedClass: string = '';
  $: combinedClass = twMerge(
    'input input-bordered w-full',
    $$restProps?.disabled ? 'text-gray-400' : '',
    !!errorText ? 'input-error' : '',
    $$restProps?.class,
  );
</script>

<label class="form-control w-full" for={$$restProps?.name}>
  {#if is_display_label}
    <div class="label">
      <span class="label-text">{$$restProps?.label}</span>
    </div>
  {/if}

  <div class="relative">
    <input
      {...$$restProps}
      class={combinedClass}
      aria-invalid={!!errorText}
      placeholder={$$restProps?.placeholder
        ? $$restProps?.placeholder
        : `Enter your ${$$restProps?.label?.toLowerCase()}...`}
    />

    {#if $$restProps?.disabled}
      <div class="text-md pointer-events-none absolute right-2 top-1/2 -translate-y-1/2 transform">
        <IcRoundLock className="inline" />
      </div>
    {/if}
  </div>

  {#if errorText}
    <div class="label">
      <span
        class="label-text-alt text-red-500 animate animate-fade-down animate-duration-100 animate-ease-out"
      >
        <IcRoundWarningAmber class="mr-1 inline-block" />
        {errorText}
      </span>
      <span class="label-text-alt"></span>
    </div>
  {/if}
</label>
