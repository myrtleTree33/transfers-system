import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ url, locals, cookies }) => {
  // Fully onboarded.  Proceed to dashboard.
  return {};
};
