<script lang="ts">
  import { twMerge } from 'tailwind-merge';

  let combinedClass: string = '';
  $: combinedClass = twMerge('hyperlink', $$restProps?.class);

  const { href, target, rel } = $$restProps;
</script>

<a {href} {target} {rel} class={combinedClass}><slot /></a>
