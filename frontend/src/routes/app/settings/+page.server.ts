import { fail, type Actions, redirect } from '@sveltejs/kit';
import { z } from 'zod';
import type { PageServerLoad } from './$types';
import { setError, superValidate } from 'sveltekit-superforms/server';
import { UserService } from '$lib/services/user_service';
import { JWTService } from '$lib/services/jwt_service';
import { zod } from 'sveltekit-superforms/adapters';
import { PasswordResetService } from '$lib/services/password_reset_service';

const changeNameSchema = z.object({
  id: z.string().min(1).max(100),
  name: z.string().min(1).max(100),
});

const changePasswordSchema = z.object({
  email: z.string().min(1).max(100),
  password: z.string().min(8).max(100).optional(),
});

export const load: PageServerLoad = async ({ request, url, cookies }) => {
  const jwtService = new JWTService(cookies);
  const access_token = jwtService.get();

  let userService = new UserService(access_token);
  let { user, error } = await userService.getProfile();
  if (error) {
    console.error('Error getting user profile', error);
  }

  const changeNameForm = await superValidate({ name: user?.name }, zod(changeNameSchema));
  const changePasswordForm = await superValidate(zod(changePasswordSchema));

  return { changeNameForm, changePasswordForm };
};

export const actions: Actions = {
  changeName: async ({ request, url, cookies, locals }) => {
    const changeNameForm = await superValidate(request, zod(changeNameSchema));

    if (!changeNameForm.valid) {
      return fail(400, { changeNameForm });
    }

    const jwtService = new JWTService(cookies);
    const access_token = jwtService.get();

    // Extract form data
    const { id, name } = changeNameForm.data;

    // // Login user
    const userService = new UserService(access_token);
    const { user, error } = await userService.updateUser(id, name);
    if (error) {
      return setError(changeNameForm, '', error.getFailureMessage() as string);
    }
  },

  changePassword: async ({ request, url, cookies, locals }) => {
    const changePasswordForm = await superValidate(request, zod(changePasswordSchema));

    if (!changePasswordForm.valid) {
      return fail(400, { changePasswordForm });
    }

    // Extract form data
    const { email } = changePasswordForm.data;

    const passwordResetService = new PasswordResetService();
    const { error } = await passwordResetService.create(email);
    if (error) {
      return setError(changePasswordForm, '', error.getFailureMessage() as string);
    }
  },
};
