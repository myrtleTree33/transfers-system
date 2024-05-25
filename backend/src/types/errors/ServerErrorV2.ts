import { Nullable } from '../Nullable';

// This class is used to represent a server error.
export class ServerErrorV2 extends Error {
  constructor(
    private failure_code?: string,
    private failure_message?: string,
  ) {
    super(`[Server error] failure_code=${failure_code} failure_message=${failure_message}`);

    Object.setPrototypeOf(this, ServerErrorV2.prototype);
  }

  getFailureCode(): Nullable<string> {
    return this.failure_code;
  }

  getFailureMessage(): Nullable<string> {
    return this.failure_message;
  }
}
