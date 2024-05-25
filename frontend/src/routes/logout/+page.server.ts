import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from '../login/$types';
import { JWTService } from '$lib/services/jwt_service';

export const load: PageServerLoad = async ({ locals, cookies }) => {
  const jwtService = new JWTService(cookies);
  await jwtService.clear();
  throw redirect(301, '/');
};
