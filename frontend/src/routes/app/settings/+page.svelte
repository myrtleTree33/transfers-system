<script lang="ts">
  import D2Input from '$lib/components/D2Input.svelte';
  import { setToast } from '$lib/stores/stores.js';
  import { superForm } from 'sveltekit-superforms';
  import IcRoundWarningAmber from '~icons/ic/round-warning-amber';

  export let data;

  const { id } = data?.user || {};

  // Superforms binding
  const {
    form: changeNameForm,
    submitting: changeNameFormSubmitting,
    errors: changeNameFormErrors,
    enhance: changeNameFormEnhance,
  } = superForm(data.changeNameForm, {
    onResult: async ({ result, cancel }) => {
      if (result.type !== 'success') {
        return;
      }

      await setToast('Name successfully changed.');
    },
    taintedMessage: null,
  });

  const {
    form: changePasswordForm,
    submitting: changePasswordFormSubmitting,
    errors: changePasswordFormErrors,
    enhance: changePasswordFormEnhance,
  } = superForm(data.changePasswordForm, {
    onResult: async ({ result, cancel }) => {
      if (result.type !== 'success') {
        return;
      }

      await setToast(
        'Password reset successful. Please check your email for further instructions.',
      );
    },
    taintedMessage: null,
  });
</script>

<div class="flex-1 p-4 flex min-h-screen">
  <div class="w-full flex flex-col gap-4">
    <div>
      <div>
        <h1 class="text-xl text-white font-bold mb-8">Account settings</h1>
      </div>
      <form
        class="flex-1 w-full flex flex-col gap-4 lg:max-w-96"
        method="POST"
        action="?/changeName"
        use:changeNameFormEnhance
      >
        {#if $changeNameFormErrors?._errors?.length > 0}
          <div role="alert" class="alert alert-error rounded-none">
            <IcRoundWarningAmber class="inline-block" />
            <span>
              {$changeNameFormErrors?._errors[0] || 'Please correct input errors and try again.'}
            </span>
          </div>
        {/if}

        <div>
          <h5>General settings</h5>
        </div>
        <div class="grid gap-4">
          <D2Input
            label="name"
            placeholder="Name"
            name="name"
            bind:value={$changeNameForm.name}
            errorText={$changeNameFormErrors.name}
          />

          <input type="hidden" name="id" value={data?.user?.id} />
        </div>

        <button type="submit" class="btn btn-primary w-full">Change name</button>
      </form>
    </div>

    <div>
      <form
        class="flex-1 w-full flex flex-col gap-4 lg:max-w-96"
        method="POST"
        action="?/changePassword"
        use:changePasswordFormEnhance
      >
        <div>
          <h5>Password</h5>
        </div>
        <div class="grid gap-4">
          <D2Input
            label="password"
            placeholder="Password"
            name="password"
            type="password"
            value={'123456789'}
            disabled
            errorText={$changePasswordFormErrors.password}
          />
          <input type="hidden" name="email" value={data?.user?.email} />
        </div>

        <button type="submit" class="btn btn-primary w-full">Reset password</button>
      </form>
    </div>
  </div>
</div>
