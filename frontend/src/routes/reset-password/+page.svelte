<script lang="ts">
  import D2Input from '$lib/components/D2Input.svelte';
  import { superForm } from 'sveltekit-superforms';
  import IcRoundWarningAmber from '~icons/ic/round-warning-amber';

  export let data;

  const { token } = data;

  // Superforms binding
  const {
    form: resetPasswordForm,
    submitting: resetPasswordFormSubmitting,
    errors: resetPasswordFormErrors,
    enhance: resetPasswordFormEnhance,
  } = superForm(data.resetPasswordForm, {
    onResult: ({ result, cancel }) => {},
    taintedMessage: null,
  });
</script>

{#if $resetPasswordFormErrors?._errors?.length > 0}
  <div role="alert" class="alert alert-error rounded-none">
    <IcRoundWarningAmber class="inline-block" />
    <span>
      {$resetPasswordFormErrors?._errors[0] || 'Please correct input errors and try again.'}
    </span>
  </div>
{/if}

<div class="min-h-screen bg-base-300 flex justify-center items-center">
  <form
    class="w-80 flex flex-col gap-6"
    method="POST"
    action="?/resetPassword"
    use:resetPasswordFormEnhance
  >
    <div>
      <h1 class="text-xl text-primary font-bold">Reset your password</h1>
      <p>Please enter your new password.</p>
    </div>

    <div class="grid gap-4">
      <D2Input
        label="password"
        placeholder="Password"
        name="password"
        type="password"
        bind:value={$resetPasswordForm.password}
        errorText={$resetPasswordFormErrors.password}
      />

      <D2Input
        label="confirm_password"
        placeholder="Confirm password"
        name="confirm_password"
        type="password"
        bind:value={$resetPasswordForm.confirm_password}
        errorText={$resetPasswordFormErrors.confirm_password}
      />

      <input type="hidden" name="token" value={token} />
    </div>

    <div class="grid gap-4">
      <button type="submit" class="btn btn-primary w-full">Reset my password</button>
    </div>
  </form>
</div>
