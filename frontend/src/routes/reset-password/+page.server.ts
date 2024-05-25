import { PasswordResetService } from '$lib/services/password_reset_service';
import { fail, type Actions, redirect } from '@sveltejs/kit';
import { z } from 'zod';
import type { PageServerLoad } from './$types';
import { setError, superValidate } from 'sveltekit-superforms/server';
import { UserService } from '$lib/services/user_service';
import { JWTService } from '$lib/services/jwt_service';
import { zod } from 'sveltekit-superforms/adapters';

const resetPasswordSchema = z
  .object({
    token: z.string(),
    password: z.string().min(8).max(100),
    confirm_password: z
      .string({ required_error: 'Confirm your password' })
      .min(1, 'Confirm your password'),
  })
  .refine((data) => data.password === data.confirm_password, {
    path: ['confirm_password'],
    message: 'Passwords do not match',
  });

export const load: PageServerLoad = async ({ url, params }) => {
  const token = url.searchParams.get('token');

  const resetPasswordForm = await superValidate(zod(resetPasswordSchema));
  const redirect_url = url.searchParams.get('redirect');
  return { token, resetPasswordForm, redirect_url };
};

export const actions: Actions = {
  resetPassword: async ({ request, locals, cookies }) => {
    const resetPasswordForm = await superValidate(request, zod(resetPasswordSchema));

    if (!resetPasswordForm.valid) {
      return fail(400, { signupForm: resetPasswordForm });
    }

    // Extract form data
    const { password, confirm_password, token } = resetPasswordForm.data;

    // Reset password
    const passwordResetService = new PasswordResetService();
    const { error: resetPasswordError } = await passwordResetService.confirm(token, password);
    if (resetPasswordError) {
      return setError(resetPasswordForm, '', resetPasswordError.getFailureMessage() as string);
    }

    // redirect to dashboard if successful
    throw redirect(301, `/app`);
  },
};
