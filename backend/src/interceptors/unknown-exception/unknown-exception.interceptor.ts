import {
  CallHandler,
  ExecutionContext,
  HttpException,
  Injectable,
  NestInterceptor,
} from '@nestjs/common';
import { Observable, catchError, throwError } from 'rxjs';

@Injectable()
export class UnknownExceptionsInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    return next.handle().pipe(
      catchError((error) => {
        console.error(error);
        let status = 500;
        const response: any = {
          code: 'unknown_error',
          description: error?.message ?? 'Internal server error',
        };

        if (error.exposeCustom_) {
          status = error.status || 500;
          response.code = error.message || 'unknown_error';

          if (error.description) {
            response.description = error.description;
          }
          if (error.exposeMeta) {
            response.meta = error.exposeMeta;
          }
        }

        return throwError(() => new HttpException(response, status, { cause: error }));
      }),
    );
  }
}
