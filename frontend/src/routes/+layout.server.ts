import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { JWTService } from '$lib/services/jwt_service';
import { UserService } from '$lib/services/user_service';
import type { User } from '../types/User';
import type { Nullable } from '../types/Nullable';

export const load: LayoutServerLoad = async ({ url, locals, cookies }) => {
  // Retrieve token
  const jwtService = new JWTService(cookies);
  const access_token = jwtService.get();
  let user: Nullable<User>;

  if (access_token) {
    const userService = new UserService(access_token);
    const { user: found_user, error } = await userService.getProfile();
    user = found_user;

    if (error) {
      console.error('Error getting user profile', error);
    }
  }

  if (!user) {
    jwtService.clear();
  }

  if (url.pathname.startsWith('/app') && !user) {
    throw redirect(301, '/login');
  }

  if ((url.pathname.startsWith('/login') || url.pathname.startsWith('/signup')) && user) {
    throw redirect(301, '/app');
  }

  let toast_message = url.searchParams.get('m');

  // Fully onboarded.  Proceed to dashboard.
  return { access_token, user, toast_message };
};
