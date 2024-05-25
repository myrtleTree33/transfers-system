<script lang="ts">
  import { twMerge } from 'tailwind-merge';

  let combinedClass: string = '';
  $: combinedClass = twMerge('py-8 grid grid-cols-4 gap-4', $$restProps?.class);
</script>

<div class={combinedClass}>
  <slot />
</div>
