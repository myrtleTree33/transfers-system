import { Module } from '@nestjs/common';
import { PrismaService } from './prisma.service';

// PrismaService is a provider that can be injected into other services or controllers.
// Use this provider to access your database.
@Module({
  providers: [PrismaService],
  exports: [PrismaService],
})
export class PrismaModule {}
