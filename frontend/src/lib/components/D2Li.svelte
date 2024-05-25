<script lang="ts">
  import { twMerge } from 'tailwind-merge';

  let combinedClass: string = '';
  $: combinedClass = twMerge('mb-2', $$restProps?.class);
</script>

<li class={combinedClass}><slot /></li>
