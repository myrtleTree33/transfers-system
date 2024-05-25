<script lang="ts">
  import { twMerge } from 'tailwind-merge';

  let combinedClass: string = '';
  $: combinedClass = twMerge(
    'rounded-md border-2 border-primary bg-base-300 p-2 justify-center text-center text-xs lg:text-base',
    $$restProps?.class,
  );
</script>

<div class={combinedClass}><slot /></div>
