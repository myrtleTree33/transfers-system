import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma.service';
import { StripeCheckoutSessionsService } from 'src/stripe-checkout-sessions/stripe-checkout-sessions.service';

// This service handles the Stripe Webhooks
// It listens for events from Stripe and handles them
// The events are verified using the Stripe secret
// The events are then handled based on their type
@Injectable()
export class StripeWebhooksService {
  constructor(
    private readonly prismaService: PrismaService,
    private readonly stripeCheckoutSessionsService: StripeCheckoutSessionsService,
  ) {}

  async handleCheckoutSessionCompleted(checkoutSessionCompleted: any) {
    const email: string = checkoutSessionCompleted?.customer_details?.email;
    const checkoutSessionId: string = checkoutSessionCompleted?.id;

    await this.stripeCheckoutSessionsService.create({
      email,
      stripe_checkout_session_id: checkoutSessionId,
    });
  }
}
