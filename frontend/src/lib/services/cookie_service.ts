import { NODE_ENV } from '$env/static/private';
import type { Cookies } from '@sveltejs/kit';

export class CookieService {
  private cookies: Cookies;

  constructor(cookies: Cookies) {
    this.cookies = cookies;
  }

  async save(key: string, token: string) {
    this.cookies.set(key, token as string, {
      path: '/',
      httpOnly: true,
      sameSite: 'lax',
      secure: NODE_ENV !== 'development',
      maxAge: 60 * 60 * 24 * 30, // 30 days
    });
  }

  get(key: string) {
    return this.cookies.get(key);
  }

  async clear(key: string) {
    await this.cookies.delete(key, {
      path: '/',
      httpOnly: true,
      sameSite: 'lax',
      secure: NODE_ENV !== 'development',
      maxAge: 60 * 60 * 24 * 30,
    });
  }
}
