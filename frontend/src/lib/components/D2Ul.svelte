<script lang="ts">
  import { twMerge } from 'tailwind-merge';

  let combinedClass: string = '';
  $: combinedClass = twMerge('list-disc list-inside', $$restProps?.class);
</script>

<ul class={combinedClass}><slot /></ul>
