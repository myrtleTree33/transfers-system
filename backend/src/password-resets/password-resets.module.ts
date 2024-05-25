import { Module } from '@nestjs/common';
import { PasswordResetsService } from './password-resets.service';
import { PasswordResetsController } from './password-resets.controller';
import { PrismaModule } from 'src/prisma/prisma.module';
import { UsersModule } from 'src/users/users.module';
import { EmailModule } from 'src/email/email.module';

@Module({
  controllers: [PasswordResetsController],
  providers: [PasswordResetsService],
  imports: [PrismaModule, UsersModule, EmailModule],
})
export class PasswordResetsModule {}
