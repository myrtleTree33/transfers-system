import { PasswordResetService } from '$lib/services/password_reset_service';
import { fail, type Actions, redirect } from '@sveltejs/kit';
import { z } from 'zod';
import type { PageServerLoad } from './$types';
import { setError, superValidate } from 'sveltekit-superforms/server';
import { UserService } from '$lib/services/user_service';
import { JWTService } from '$lib/services/jwt_service';
import { zod } from 'sveltekit-superforms/adapters';

const startResetPasswordSchema = z.object({
  email: z.string().email('Email is not valid'),
});

export const load: PageServerLoad = async ({ url, params }) => {
  const startResetPasswordForm = await superValidate(zod(startResetPasswordSchema));
  return { startResetPasswordForm };
};

export const actions: Actions = {
  startResetPassword: async ({ request, locals, cookies }) => {
    const startResetPasswordForm = await superValidate(request, zod(startResetPasswordSchema));

    if (!startResetPasswordForm.valid) {
      return fail(400, { signupForm: startResetPasswordForm });
    }

    // Extract form data
    const { email } = startResetPasswordForm.data;

    // Reset password
    const passwordResetService = new PasswordResetService();
    const { error: startResetPasswordError } = await passwordResetService.create(email);
    if (startResetPasswordError) {
      return setError(
        startResetPasswordForm,
        '',
        startResetPasswordError.getFailureMessage() as string,
      );
    }

    // redirect to main page if successful
    throw redirect(301, `/`);
  },
};
