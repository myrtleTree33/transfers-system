import { Module } from '@nestjs/common';
import { StripeWebhooksService } from './stripe-webhooks.service';
import { StripeWebhooksController } from './stripe-webhooks.controller';
import { PrismaModule } from 'src/prisma/prisma.module';
import { StripeCheckoutSessionsModule } from 'src/stripe-checkout-sessions/stripe-checkout-sessions.module';

@Module({
  controllers: [StripeWebhooksController],
  providers: [StripeWebhooksService],
  imports: [PrismaModule, StripeCheckoutSessionsModule],
  exports: [StripeWebhooksService],
})
export class StripeWebhooksModule {}
