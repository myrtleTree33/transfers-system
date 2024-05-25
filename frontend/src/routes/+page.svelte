<script lang="ts">
  /**
   * This is the splash page for the application.
   * It is the first page that users see when they visit the site.
   */

  export let data;
  const { is_splash } = data;
  import Typewriter from 'typewriter-effect/dist/core';
  import MarqueeTextWidget from 'svelte-marquee-text-widget';
  import D2Pricing from '$lib/components/D2Pricing.svelte';
  import { PUBLIC_PRODUCT_NAME } from '$env/static/public';
  import { onMount } from 'svelte';

  onMount(() => {
    const typewriter = new Typewriter('#typewriter', {
      strings: [`From Idea to Startup in 2 Hours.`],
      autoStart: true,
      delay: 5,
      cursor: '|',
      deleteSpeed: 10000000,
      loop: false,
    });

    typewriter.start();
  });
</script>

<svelte:head>
  <title>{PUBLIC_PRODUCT_NAME} - Get your app up on the second day.</title>
</svelte:head>

<!-- Splash section -->
<section class="bg-gradient-to-b from-[#e53935] to-[#e35d5b] min-h-screen w-full pt-7">
  <div class="p-4 flex-1 flex flex-col justify-center items-center min-h-screen gap-8">
    <a href="/signup" class="text-5xl font-bold text-gray-400 pb-4">
      <h1
        id="typewriter"
        class="text-primary-content text-center uppercase bg-primary p-4 rounded-md"
      >
        From Idea to Startup in 2 Hours.
      </h1>
    </a>
    <div class="w-full flex justify-center items-center flex-col gap-4">
      <div class="text-lg text-center text-primary-content font-fjalla-one">
        The only NestJS + SvelteKit boilerplate you'll ever need to launch your idea in a single
        day.
      </div>

      <div class="text-white w-full text-sm max-w-prose font-mono">
        <MarqueeTextWidget duration={30}>
          <ul class="flex w-full gap-4 list-none mx-4">
            <li>NestJS</li>
            <li>SvelteKit</li>
            <li>Prisma</li>
            <li>Postgres</li>
            <li>Docker</li>
            <li>Devbox</li>
            <li>Sentry</li>
            <li>Stripe</li>
            <li>Postmark</li>
            <li>Fly.io</li>
          </ul>
        </MarqueeTextWidget>
      </div>

      <div class="mt-8 grid gap-4 w-64">
        <a class="text-white btn btn-lg btn-primary uppercase font-fjalla-one" href="/signup"
          >Get {PUBLIC_PRODUCT_NAME}</a
        >
        <a class="text-white btn btn-lg btn-outline" href="/login">Login</a>
      </div>

      <div>
        <ul class="grid gap-x-8 max-w-4xl font-fjalla-one text-primary-content text-lg">
          <li class="li-checked">Separate frontend and backend</li>
          <li class="li-checked">Ship on day 2</li>
          <li class="li-checked">User authentication</li>
          <li class="li-checked">App monitoring + CI / CD</li>
        </ul>
      </div>
    </div>
  </div>
</section>

<!-- FAQ section -->
<section class="py-20 bg-base-100 w-full flex flex-col items-center justify-center p-4">
  <div class="w-full flex flex-col gap-8 items-center justify-center">
    <h2 class="text-3xl font-bold text-primary text-center">FAQ</h2>
    <div class="grid lg:grid-cols-2 gap-4">
      <div class="flex flex-col gap-4">
        <h4 class="text-primary">Learn why {PUBLIC_PRODUCT_NAME} is better than others.</h4>
        <p>
          Have a question? <a href="mailto:joel.tong@gmail.com" class="hyperlink text-sm"
            >Send a query</a
          >
          {' '}and we'll get in touch.
        </p>
      </div>
      <div class="join join-vertical w-full bg-primary-content max-w-prose">
        <div class="collapse collapse-arrow join-item border border-base-300">
          <input type="radio" name="my-accordion-3" />
          <div class="collapse-title text-sm font-bold">What is packaged?</div>
          <div class="collapse-content text-sm">
            <p>
              1. The NestJS + SvelteKit starter with all you need to run password-based user logins,
              components, subscriptions and payments integration.
            </p>
            <p>2. The comprehensive documentation to get you up and shipping in 3 hours.</p>
            <p>
              3. A comprehensive guide by our seasoned sales professional co-founder, to write copy
              that sells.
            </p>
            <p>
              4. One-on-one session with the creators to help you get started and answer any
              questions you may have.
            </p>
            <p>
              5. Complete boilerplate and setup guide to launching monitoring + CI / CD deployments
              with Fly.io
            </p>
          </div>
        </div>
        <div class="collapse collapse-arrow join-item border border-base-300">
          <input type="radio" name="my-accordion-3" />
          <div class="collapse-title text-sm font-bold">
            Can I deploy to Vercel / Netlify / Heroku / AWS instead of Fly.io?
          </div>
          <div class="collapse-content text-sm">
            <p>
              Yes absolutely! You could use Mailgun instead of Postmark, MongoDB instead of
              Postgres, for example. Our founders have personally found our chosen stack to work and
              scale fastest for early startups and indie projects however.
            </p>
          </div>
        </div>
        <div class="collapse collapse-arrow join-item border border-base-300">
          <input type="radio" name="my-accordion-3" />
          <div class="collapse-title text-sm font-bold">Are we just a website template?</div>
          <div class="collapse-content text-sm">
            <p>
              Absolutely not! When you purchase {PUBLIC_PRODUCT_NAME}, you purchase a suite of tools
              to help your startup succeed.
            </p>
          </div>
        </div>

        <div class="collapse collapse-arrow join-item border border-base-300">
          <input type="radio" name="my-accordion-3" />
          <div class="collapse-title text-sm font-bold">Are there any costs associated?</div>
          <div class="collapse-content text-sm">
            <p>
              Our chosen choice of provider, Fly.io, lets you deploy your site at minimum cost
              (startup from $5 / mo.), giving you the flexibiliy to prototype your app and keep
              costs low.
            </p>
          </div>
        </div>

        <div class="collapse collapse-arrow join-item border border-base-300">
          <input type="radio" name="my-accordion-3" />
          <div class="collapse-title text-sm font-bold">
            How often is {PUBLIC_PRODUCT_NAME} updated?
          </div>
          <div class="collapse-content text-sm">
            <p>
              {PUBLIC_PRODUCT_NAME} was built with {PUBLIC_PRODUCT_NAME}. We ship actively, and get
              new updates in regularly.
            </p>
          </div>
        </div>

        <div class="collapse collapse-arrow join-item border border-base-300">
          <input type="radio" name="my-accordion-3" />
          <div class="collapse-title text-sm font-bold">
            Is authentication dependent on any providers?
          </div>
          <div class="collapse-content text-sm">
            <p>
              No. We built our own authentication system. However, {PUBLIC_PRODUCT_NAME} can be easily
              integrated with any auth provider such as Auth0, Firebase, or Supabase.
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- Pricing section -->
<section class="min-h-screen bg-base-100 w-full flex flex-col items-center justify-center h-full">
  <div class="w-full flex flex-col gap-8 items-center justify-center px-4 py-12">
    <a id="#pricing" class=""></a>
    <h2 class="text-3xl font-bold text-white text-center">Pricing</h2>
    <div class="grid lg:grid-cols-3 gap-8">
      <D2Pricing href="/">
        <svelte:fragment slot="header">Free</svelte:fragment>
        <svelte:fragment slot="usualPrice">$269</svelte:fragment>
        <svelte:fragment slot="currentPrice">$100</svelte:fragment>
        <svelte:fragment slot="currency">USD</svelte:fragment>
        <svelte:fragment>
          <div class="">Forever</div>
          <div class="divider"></div>
          <div class="grid gap-4">
            <div class="">For hobbyists and small projects</div>
            <ul class="grid gap-1">
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 project</li>
              <li class="li-checked">1 environment</li>
              <li class="li-checked">1 database</li>
            </ul>
            <a href="/signup" class="btn btn-accent w-full">Get {PUBLIC_PRODUCT_NAME}</a>
          </div>
        </svelte:fragment>
      </D2Pricing>

      <D2Pricing selectedHeader="Popular" isSelected={true} class="order-first lg:order-none">
        <svelte:fragment slot="header"><span class="text-success">Personal</span></svelte:fragment>
        <svelte:fragment slot="usualPrice">$269</svelte:fragment>
        <svelte:fragment slot="currentPrice">$150</svelte:fragment>
        <svelte:fragment slot="currency">USD</svelte:fragment>
        <svelte:fragment>
          <div class="">Monthly</div>
          <div class="divider"></div>
          <div class="grid gap-4 w-full">
            <div class="">For indie hackers</div>
            <ul class="grid gap-1">
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 project</li>
              <li class="li-checked">1 environment</li>
              <li class="li-checked">1 database</li>
            </ul>
            <a href="/signup" class="btn btn-accent w-full">Get {PUBLIC_PRODUCT_NAME}</a>
          </div>
        </svelte:fragment>
      </D2Pricing>

      <D2Pricing href="/">
        <svelte:fragment slot="header">Professional</svelte:fragment>
        <svelte:fragment slot="usualPrice">$269</svelte:fragment>
        <svelte:fragment slot="currentPrice">$250</svelte:fragment>
        <svelte:fragment slot="currency">USD</svelte:fragment>
        <svelte:fragment>
          <div class="">Monthly</div>
          <div class="divider"></div>
          <div class="grid gap-4">
            <div class="">For professionals</div>
            <ul class="grid gap-1">
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 user</li>
              <li class="li-checked">1 project</li>
              <li class="li-checked">1 environment</li>
              <li class="li-checked">1 database</li>
            </ul>
            <a href="/signup" class="btn btn-accent w-full">Get {PUBLIC_PRODUCT_NAME}</a>
          </div>
        </svelte:fragment>
      </D2Pricing>
    </div>
  </div>
</section>

<!-- Second Call to Action section -->
<section class="min-h-screen bg-accent w-full flex flex-col items-center justify-center">
  <div class="w-full flex flex-col gap-8 items-center justify-center">
    <div class="flex flex-col flex-1 gap-4 items-center text-primary-content text-center mb-20">
      <h2 class="text-3xl font-bold pb-0 text-primary-content">Ship to production with</h2>
      <div class="p-4 rounded-md bg-primary w-fit mt-2 font-fjalla-one uppercase text-2xl">
        {PUBLIC_PRODUCT_NAME}
      </div>
    </div>
    <a class="text-white btn btn-lg btn-primary uppercase font-fjalla-one text-2xl" href="/signup"
      >Get {PUBLIC_PRODUCT_NAME}</a
    >
  </div>
</section>
