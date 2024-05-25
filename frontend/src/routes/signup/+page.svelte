<script lang="ts">
  /**
   * This is the signup page.
   */

  import D2Input from '$lib/components/D2Input.svelte';
  import { superForm } from 'sveltekit-superforms';

  export let data;

  const {
    form: signup_form,
    submitting: signup_form_submitting,
    errors: signup_form_errors,
    enhance: signup_form_enhance,
  } = superForm(data.signup_form, {
    onResult: ({ result, cancel }) => {},
    taintedMessage: null,
  });
</script>

<div class="min-h-screen bg-base-300 flex justify-center items-center">
  <form class="w-80 flex flex-col gap-6" method="POST" action="?/signup">
    <div>
      <h1 class="text-xl text-primary font-bold">Sign up new account</h1>
      <p class="text-sm">Already a member? <a href="/login" class="hyperlink">Sign in</a></p>
    </div>

    <div class="grid gap-4">
      <D2Input
        label="email"
        placeholder="Email"
        name="email"
        bind:value={$signup_form.email}
        errorText={$signup_form_errors.email}
      />

      <D2Input
        label="name"
        placeholder="Name"
        name="name"
        bind:value={$signup_form.name}
        errorText={$signup_form_errors.name}
      />

      <D2Input
        label="password"
        placeholder="Password"
        name="password"
        type="password"
        bind:value={$signup_form.password}
        errorText={$signup_form_errors.password}
      />

      <D2Input
        label="confirm_password"
        placeholder="Confirm password"
        name="confirm_password"
        type="password"
        bind:value={$signup_form.password}
        errorText={$signup_form_errors.password}
      />

      <D2Input
        is_display_label={true}
        label={`What is ${data.a} + ${data.b}?`}
        placeholder="Security question"
        name="security_question"
        type="security_question"
        bind:value={$signup_form.security_question}
        errorText={$signup_form_errors.security_question}
      />

      <input name="redirect_url" hidden value={data.redirect_url} />
      <input name="a" hidden value={data.a} />
      <input name="b" hidden value={data.b} />
    </div>

    <div class="grid gap-4">
      <button class="btn btn-primary w-full">Sign up</button>

      <div class="flex justify-between text-sm">
        <a href="/reset" class="hyperlink">Forgot password?</a>
        <a href="/login" class="hyperlink">Sign in</a>
      </div>
    </div>
  </form>
</div>
