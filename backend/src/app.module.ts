import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { ConfigModule } from '@nestjs/config';
import { PrismaModule } from './prisma/prisma.module';
import { AuthModule } from './auth/auth.module';
import { PasswordResetsModule } from './password-resets/password-resets.module';
import { PostmarkService } from './postmark/postmark.service';
import { PostmarkModule } from './postmark/postmark.module';
import { EmailModule } from './email/email.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: '.env.development',
    }),
    PrismaModule,
    UsersModule,
    AuthModule,
    PasswordResetsModule,
    PostmarkModule,
    EmailModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
