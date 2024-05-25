import {
  wrappedAuthFetch,
  wrappedFetch,
  type WrappedFetchRequest,
  type WrappedFetchResponse,
} from './fetch_helper';
import type { Nullable } from '../../types/Nullable';
import type { User } from '../../types/User';
import type { ServerErrorV2 } from '../../types/ServerErrorV2';

export class PasswordResetService {
  constructor() {}

  async create(email: string): Promise<{ error?: Nullable<ServerErrorV2> }> {
    const request: WrappedFetchRequest = {
      method: 'POST',
      endpoint: '/password-resets',
      body: {
        email,
      },
    };

    // Execute fetch
    const { error } = await wrappedFetch(request);

    // Return response
    return { error: error };
  }

  async confirm(
    token?: string,
    password?: string,
  ): Promise<{
    error: Nullable<ServerErrorV2>;
  }> {
    const request: WrappedFetchRequest = {
      method: 'POST',
      endpoint: '/password-resets/confirm',
      body: {
        token,
        password,
      },
    };

    // Execute fetch
    const { body, error } = await wrappedFetch(request);

    // Return response
    return { error: error };
  }
}
