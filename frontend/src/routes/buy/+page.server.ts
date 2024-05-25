import { redirect } from '@sveltejs/kit';

import { PUBLIC_STRIPE_CHECKOUT_LINK } from '$env/static/public';

export function load() {
  throw redirect(302, PUBLIC_STRIPE_CHECKOUT_LINK);
}
