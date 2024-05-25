import { Controller, Get, Post, Body, Patch, Param, Delete } from '@nestjs/common';
import { StripeCheckoutSessionsService } from './stripe-checkout-sessions.service';
import { CreateStripeCheckoutSessionDto } from './dto/create-stripe-checkout-session.dto';
import { UpdateStripeCheckoutSessionDto } from './dto/update-stripe-checkout-session.dto';

@Controller('stripe-checkout-sessions')
export class StripeCheckoutSessionsController {
  constructor(private readonly stripeCheckoutSessionsService: StripeCheckoutSessionsService) {}

  @Get(':id')
  findOneByStripeCheckoutSessionId(@Param('id') id: string) {
    return this.stripeCheckoutSessionsService.findOne(id);
  }
}
