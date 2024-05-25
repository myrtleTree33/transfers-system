import { NODE_ENV } from '$env/static/private';
import type { Cookies } from '@sveltejs/kit';

export class JWTService {
  private cookies: Cookies;

  constructor(cookies: Cookies) {
    this.cookies = cookies;
  }

  async save(token: string) {
    this.cookies.set('token', token as string, {
      path: '/',
      httpOnly: true,
      sameSite: 'lax',
      secure: NODE_ENV !== 'development',
      maxAge: 60 * 60 * 24 * 30, // 30 days
    });
  }

  get() {
    return this.cookies.get('token');
  }

  async clear() {
    await this.cookies.delete('token', {
      path: '/',
      httpOnly: true,
      sameSite: 'lax',
      secure: NODE_ENV !== 'development',
      maxAge: 60 * 60 * 24 * 30,
    });
  }
}
