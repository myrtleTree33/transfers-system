import { Module } from '@nestjs/common';
import { StripeCheckoutSessionsService } from './stripe-checkout-sessions.service';
import { StripeCheckoutSessionsController } from './stripe-checkout-sessions.controller';
import { PrismaModule } from 'src/prisma/prisma.module';

@Module({
  controllers: [StripeCheckoutSessionsController],
  providers: [StripeCheckoutSessionsService],
  imports: [PrismaModule],
  exports: [StripeCheckoutSessionsService],
})
export class StripeCheckoutSessionsModule {}
