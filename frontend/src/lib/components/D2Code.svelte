<script lang="ts">
  import { twMerge } from 'tailwind-merge';

  let combinedClass: string = '';
  $: combinedClass = twMerge('mockup-code my-8', $$restProps?.class);
</script>

<div class={combinedClass}>
  <pre data-prefix="$"><code><slot /></code></pre>
</div>
