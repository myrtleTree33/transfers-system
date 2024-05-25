import { UnknownExceptionsInterceptor } from './unknown-exception.interceptor';

describe('UnknownExceptionInterceptor', () => {
  it('should be defined', () => {
    expect(new UnknownExceptionsInterceptor()).toBeDefined();
  });
});
