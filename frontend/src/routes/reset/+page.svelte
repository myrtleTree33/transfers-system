<script lang="ts">
  import D2Input from '$lib/components/D2Input.svelte';
  import { superForm } from 'sveltekit-superforms';
  import IcRoundWarningAmber from '~icons/ic/round-warning-amber';

  export let data;

  // Superforms binding
  const {
    form: startResetPasswordForm,
    submitting: startResetPasswordFormSubmitting,
    errors: startResetPasswordFormErrors,
    enhance: startResetPasswordFormEnhance,
  } = superForm(data.startResetPasswordForm, {
    onResult: ({ result, cancel }) => {},
    taintedMessage: null,
  });
</script>

{#if $startResetPasswordFormErrors?._errors?.length > 0}
  <div role="alert" class="alert alert-error rounded-none">
    <IcRoundWarningAmber class="inline-block" />
    <span>
      {$startResetPasswordFormErrors?._errors[0] || 'Please correct input errors and try again.'}
    </span>
  </div>
{/if}

<div class="min-h-screen bg-base-300 flex justify-center items-center">
  <form
    class="w-80 flex flex-col gap-6"
    method="POST"
    action="?/startResetPassword"
    use:startResetPasswordFormEnhance
  >
    <div>
      <h1 class="text-xl text-primary font-bold">Forgot your password?</h1>
      <p class="text-sm">
        Remembered your password? <a href="/login" class="hyperlink">Sign in</a>
      </p>
    </div>

    <div class="grid gap-4">
      <D2Input
        label="email"
        placeholder="Email"
        name="email"
        type="email"
        bind:value={$startResetPasswordForm.email}
        errorText={$startResetPasswordFormErrors.email}
      />
    </div>

    <button class="btn btn-primary w-full" type="submit">Send me password reset instructions</button
    >
  </form>
</div>
