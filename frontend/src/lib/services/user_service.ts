import {
  wrappedAuthFetch,
  wrappedFetch,
  type WrappedFetchRequest,
  type WrappedFetchResponse,
} from './fetch_helper';
import type { Nullable } from '../../types/Nullable';
import type { User } from '../../types/User';
import type { ServerErrorV2 } from '../../types/ServerErrorV2';

export class UserService {
  private auth_token?: string;

  constructor(auth_token?: string) {
    this.auth_token = auth_token;
  }

  async signup(
    email: string,
    password: string,
    name: string,
  ): Promise<{ user: Nullable<User>; error: Nullable<ServerErrorV2> }> {
    const request: WrappedFetchRequest = {
      method: 'POST',
      endpoint: '/users/signup/password',
      body: {
        email,
        password,
        name,
      },
    };

    // Execute fetch
    const { body, error } = await wrappedFetch(request);

    // Return response
    return { user: body, error: error };
  }

  async login(
    email?: string,
    password?: string,
  ): Promise<{
    access_token: Nullable<string>;
    error: Nullable<ServerErrorV2>;
  }> {
    const request: WrappedFetchRequest = {
      method: 'POST',
      endpoint: '/users/login/password',
      body: {
        email,
        password,
      },
    };

    // Execute fetch
    const { body, error } = await wrappedFetch(request);

    // Return response
    return { access_token: body?.access_token, error: error };
  }

  async getProfile(): Promise<{
    user: Nullable<User>;
    error: Nullable<ServerErrorV2>;
  }> {
    const request: WrappedFetchRequest = {
      method: 'GET',
      endpoint: '/users/profile',
    };

    // Execute fetch
    const { body, error } = await wrappedAuthFetch(request, this.auth_token);

    console.log(body);

    // Return response
    return { user: body, error };
  }

  async updateUser(
    id: string,
    name: string,
  ): Promise<{
    user: Nullable<User>;
    error: Nullable<ServerErrorV2>;
  }> {
    const request: WrappedFetchRequest = {
      method: 'PATCH',
      endpoint: `/users/${id}`,
      body: {
        name,
      },
    };

    // Execute fetch
    const { body, error } = await wrappedAuthFetch(request, this.auth_token);

    console.log(body);

    // Return response
    return { user: body, error };
  }
}
