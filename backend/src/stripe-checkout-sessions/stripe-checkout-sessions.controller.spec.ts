import { Test, TestingModule } from '@nestjs/testing';
import { StripeCheckoutSessionsController } from './stripe-checkout-sessions.controller';
import { StripeCheckoutSessionsService } from './stripe-checkout-sessions.service';

describe('StripeCheckoutSessionsController', () => {
  let controller: StripeCheckoutSessionsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [StripeCheckoutSessionsController],
      providers: [StripeCheckoutSessionsService],
    }).compile();

    controller = module.get<StripeCheckoutSessionsController>(StripeCheckoutSessionsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
