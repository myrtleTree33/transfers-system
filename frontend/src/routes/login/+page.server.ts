import { fail, type Actions, redirect } from '@sveltejs/kit';
import { z } from 'zod';
import type { PageServerLoad } from './$types';
import { setError, superValidate } from 'sveltekit-superforms/server';
import { UserService } from '$lib/services/user_service';
import { JWTService } from '$lib/services/jwt_service';
import { zod } from 'sveltekit-superforms/adapters';

const login_user_schema = z.object({
  email: z.string().email('Email is not valid'),
  password: z.string().min(1).max(100),
  remember_me: z.boolean().optional(),
  redirect_url: z.string().optional(),
});

export const load: PageServerLoad = async ({ url, locals, cookies }) => {
  const loginForm = await superValidate(zod(login_user_schema));

  return { loginForm };
};

export const actions: Actions = {
  login: async ({ request, url, cookies, locals }) => {
    const loginForm = await superValidate(request, zod(login_user_schema));

    if (!loginForm.valid) {
      return fail(400, { loginForm });
    }

    // Extract form data
    const { email, password, redirect_url } = loginForm.data;

    // Login user
    const userService = new UserService();
    const { access_token, error } = await userService.login(email, password);
    if (error) {
      return setError(loginForm, '', error.getFailureMessage() as string);
    }

    // Store token in cookie
    const jwtService = new JWTService(cookies);
    jwtService.save(access_token as string);

    // redirect to app if successful
    if (redirect_url?.startsWith('/') && redirect_url?.startsWith('/login')) {
      throw redirect(301, redirect_url);
    }
    // Default to app
    console.log('-------------------');
    throw redirect(301, `/app`);
  },
};
