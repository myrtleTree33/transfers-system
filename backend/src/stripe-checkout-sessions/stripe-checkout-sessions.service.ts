import { Injectable } from '@nestjs/common';
import { CreateStripeCheckoutSessionDto } from './dto/create-stripe-checkout-session.dto';
import { UpdateStripeCheckoutSessionDto } from './dto/update-stripe-checkout-session.dto';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class StripeCheckoutSessionsService {
  constructor(private readonly prismaService: PrismaService) {}

  async create(createStripeCheckoutSessionDto: CreateStripeCheckoutSessionDto) {
    const { email, stripe_checkout_session_id } = createStripeCheckoutSessionDto;
    const stripeCheckoutSession = await this.prismaService.stripeCheckoutSession.create({
      data: {
        email,
        stripe_checkout_session_id,
        status: 'pending',
      },
    });

    return stripeCheckoutSession;
  }

  async findOne(id: string) {
    return this.prismaService.stripeCheckoutSession.findUnique({
      where: { stripe_checkout_session_id: id },
    });
  }

  update(
    stripe_checkout_session_id: string,
    updateStripeCheckoutSessionDto: UpdateStripeCheckoutSessionDto,
  ) {
    return this.prismaService.stripeCheckoutSession.update({
      where: { stripe_checkout_session_id },
      data: {
        ...updateStripeCheckoutSessionDto,
        updated_at: new Date(),
      },
    });
  }

  remove(id: number) {
    return `This action removes a #${id} stripeCheckoutSession`;
  }
}
