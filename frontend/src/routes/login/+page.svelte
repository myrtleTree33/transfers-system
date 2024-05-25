<script lang="ts">
  import D2Input from '$lib/components/D2Input.svelte';
  import { superForm } from 'sveltekit-superforms';
  import IcRoundWarningAmber from '~icons/ic/round-warning-amber';

  export let data;

  // Superforms binding
  const {
    form: loginForm,
    submitting: loginFormSubmitting,
    errors: loginFormErrors,
    enhance: loginFormEnhance,
  } = superForm(data.loginForm, {
    onResult: ({ result, cancel }) => {},
    taintedMessage: null,
  });
</script>

{#if $loginFormErrors?._errors?.length > 0}
  <div role="alert" class="alert alert-error rounded-none">
    <IcRoundWarningAmber class="inline-block" />
    <span>
      {$loginFormErrors?._errors[0] || 'Please correct input errors and try again.'}
    </span>
  </div>
{/if}

<div class="min-h-screen bg-base-300 flex justify-center items-center">
  <form class="w-80 flex flex-col gap-6" method="POST" action="?/login" use:loginFormEnhance>
    <div>
      <h1 class="text-xl text-primary font-bold">Login to your account</h1>
      <p class="text-sm">Not a member? <a href="/signup" class="hyperlink">Sign up</a></p>
    </div>

    <div class="grid gap-4">
      <D2Input
        label="email"
        placeholder="Email"
        name="email"
        bind:value={$loginForm.email}
        errorText={$loginFormErrors.email}
      />

      <D2Input
        label="password"
        placeholder="Password"
        name="password"
        type="password"
        bind:value={$loginForm.password}
        errorText={$loginFormErrors.password}
      />
    </div>

    <div class="grid gap-4">
      <div class="form-control">
        <label class="label cursor-pointer">
          <span class="label-text">Remember me</span>
          <input type="checkbox" class="toggle" checked />
        </label>
      </div>

      <button type="submit" class="btn btn-primary w-full">Log in</button>

      <div class="flex justify-between text-sm">
        <a href="/reset" class="hyperlink">Forgot password?</a>
        <a href="/signup" class="hyperlink">Sign up</a>
      </div>
    </div>
  </form>
</div>
