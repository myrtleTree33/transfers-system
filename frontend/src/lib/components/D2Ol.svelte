<script lang="ts">
  import { twMerge } from 'tailwind-merge';

  let combinedClass: string = '';
  $: combinedClass = twMerge('list-decimal list-inside', $$restProps?.class);
</script>

<ol class={combinedClass}><slot /></ol>
