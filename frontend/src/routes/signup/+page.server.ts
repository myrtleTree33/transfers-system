import { fail, type Actions, redirect } from '@sveltejs/kit';
import { z } from 'zod';
import type { PageServerLoad } from './$types';
import { setError, superValidate } from 'sveltekit-superforms/server';
import { UserService } from '$lib/services/user_service';
import { JWTService } from '$lib/services/jwt_service';
import { zod } from 'sveltekit-superforms/adapters';

const signup_user_schema = z
  .object({
    email: z.string().email('Email is not valid'),
    name: z.string().min(1).max(170),
    password: z.string().min(8).max(100),
    confirm_password: z
      .string({ required_error: 'Confirm your password' })
      .min(1, 'Confirm your password'),
    a: z.number().min(1).max(100),
    b: z.number().min(1).max(100),
    security_question: z.number().min(1).max(100),
    redirect_url: z.string().optional(),
  })
  .refine((data) => data.a + data.b === data.security_question, {
    path: ['security_question'],
    message: 'Wrong answer',
  })
  .refine((data) => data.password === data.confirm_password, {
    path: ['confirm_password'],
    message: 'Passwords do not match',
  });

export const load: PageServerLoad = async ({ url, locals }) => {
  const a = Math.floor(Math.random() * 9) + 1;
  const b = Math.floor(Math.random() * 9) + 1;

  const signup_form = await superValidate(zod(signup_user_schema));
  const redirect_url = url.searchParams.get('redirect');
  return { signup_form, redirect_url, a, b };
};

export const actions: Actions = {
  signup: async ({ request, locals, cookies }) => {
    const signup_form = await superValidate(request, zod(signup_user_schema));

    if (!signup_form.valid) {
      return fail(400, { signupForm: signup_form });
    }

    // Extract form data
    const { email, name, password, confirm_password, redirect_url } = signup_form.data;

    // Signup user
    const userService = new UserService();
    const { error: createUserError } = await userService.signup(email, password, name);
    if (createUserError) {
      return setError(signup_form, '', createUserError.getFailureMessage() as string);
    }

    // Login user
    const { access_token, error: loginUserError } = await userService.login(email, password);
    if (loginUserError) {
      console.error(loginUserError);
      // Redirect to login page in event of error
      throw redirect(301, `/login`);
    }

    // Store token in cookie
    const jwtService = new JWTService(cookies);
    await jwtService.save(access_token as string);

    // redirect to dashboard if successful
    throw redirect(301, `/app?m=Welcome! Please verify your email address.`);
  },
};
