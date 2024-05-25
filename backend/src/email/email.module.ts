import { Module } from '@nestjs/common';
import { EmailService } from './email.service';
import { PostmarkModule } from 'src/postmark/postmark.module';

// The Email module is used to import the Email service
// Use this module to send emails
// For an example, refer to the User module
@Module({
  providers: [EmailService],
  exports: [EmailService],
  imports: [PostmarkModule.forRootAsync()],
})
export class EmailModule {}
