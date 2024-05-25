import {
  Headers,
  Body,
  Controller,
  Post,
  HttpException,
  RawBodyRequest,
  Req,
} from '@nestjs/common';
import stripe from 'stripe';
import { StripeWebhooksService } from './stripe-webhooks.service';

// This is the controller for the Stripe Webhooks
// It listens for events from Stripe and handles them
// The events are verified using the Stripe secret
// The events are then handled based on their type
// Use the StripeWebhooksService to handle Stripe events
@Controller('webhooks/stripe')
export class StripeWebhooksController {
  constructor(private readonly stripeWebhooksService: StripeWebhooksService) {}

  @Post()
  async create(@Headers() headers, @Req() req: RawBodyRequest<Request>) {
    const sig = headers['stripe-signature'];
    let event;

    try {
      event = stripe.webhooks.constructEvent(req.rawBody, sig, process.env.STRIPE_WEBHOOK_SECRET);
    } catch (err) {
      throw new HttpException(`Webhook Error: ${err.message}`, 400);
    }

    // Handle the event
    switch (event.type) {
      case 'checkout.session.completed':
        const checkoutSessionCompleted = event.data.object;
        await this.stripeWebhooksService.handleCheckoutSessionCompleted(checkoutSessionCompleted);
        break;

      default:
        console.log(`Unhandled event type ${event.type}`);
    }

    return {};
  }
}
