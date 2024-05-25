import { Test, TestingModule } from '@nestjs/testing';
import { StripeCheckoutSessionsService } from './stripe-checkout-sessions.service';

describe('StripeCheckoutSessionsService', () => {
  let service: StripeCheckoutSessionsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [StripeCheckoutSessionsService],
    }).compile();

    service = module.get<StripeCheckoutSessionsService>(StripeCheckoutSessionsService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
