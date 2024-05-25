<script lang="ts">
  import { toast_store } from '$lib/stores/stores';
  import { createEventDispatcher } from 'svelte';
  import { twMerge } from 'tailwind-merge';
  import IcRoundLock from '~icons/ic/round-lock';
  import IcRoundWarningAmber from '~icons/ic/round-warning-amber';
  import IcOutlineInfo from '~icons/ic/outline-info';

  const dispatch = createEventDispatcher();

  let combinedClass: string = '';
  $: combinedClass = twMerge(
    'alert alert-info rounded-md text-sm text-left flex',
    is_toast_visible ? '' : 'animate-fade animate-duration-300 animate-reverse',
  );

  let is_toast_visible: boolean = false;

  // flash toast if message changed
  const unsubscribe = toast_store.subscribe((value) => {
    is_toast_visible = true;

    setTimeout(() => {
      is_toast_visible = false;
    }, 1000);
  });
</script>

{#if $toast_store}
  <div class="toast toast-top z-[99999] w-full">
    <div class={combinedClass}>
      <div>
        <IcOutlineInfo class="inline" />
      </div>
      <div class="text-wrap">{$toast_store}</div>
    </div>
  </div>
{/if}
