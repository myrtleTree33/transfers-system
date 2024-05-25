import { DynamicModule, Module, Provider } from '@nestjs/common';
import { POSTMARK_SERVICE } from './postmark.constans';
import * as postmark from 'postmark';

// This module is used to import the Postmark service
// It is a direct dependency of the Email module.
// Do not use this module directly in other modules.
// Instead, use the Email module.
// Refer to the User module for an example of how to use the Email module.
@Module({})
export class PostmarkModule {
  public static forRootAsync(): DynamicModule {
    const postmarkProvider: Provider = {
      provide: POSTMARK_SERVICE,
      useFactory: async () => {
        const postmarkClient = new postmark.ServerClient(process.env.POSTMARK_API_KEY);

        return postmarkClient;
      },
    };

    return {
      exports: [postmarkProvider],
      module: PostmarkModule,
      providers: [postmarkProvider],
    };
  }
}
