<script lang="ts">
  import { twMerge } from 'tailwind-merge';
  import type { DictionaryType } from '../../types/DictionaryType';
  import type { KeyValue } from '../../types/KeyValue';
  import { createEventDispatcher, onMount } from 'svelte';

  const dispatch = createEventDispatcher();

  export let value: string = '';
  export let name: string = '';
  let selected_value: string = '';

  export let items: KeyValue<string, string>[] = [
    { k: '1:1', v: '1:1' },
    { k: '2:1', v: '2:1' },
    { k: '5:4', v: '5:4' },
    { k: '3:2', v: '3:2' },
    { k: '7:4', v: '7:4' },
    { k: '16:9', v: '16:9' },
    { k: '9:16', v: '9:16' },
    { k: '1:2', v: '1:2' },
  ];

  let selected: number = 0;

  $: value = selected_value;

  onMount(() => {
    selected_value = items[selected].v;
  });
</script>

<div class="flex gap-4">
  {#each items as item, i}
    <button
      class={twMerge('btn btn-primary', i !== selected ? 'btn-outline' : '')}
      on:click|preventDefault|stopPropagation={() => {
        selected = i;
        dispatch('click', item.v);
        selected_value = item.v;
      }}>{item.k}</button
    >
  {/each}
  <input {name} bind:value hidden />
</div>
