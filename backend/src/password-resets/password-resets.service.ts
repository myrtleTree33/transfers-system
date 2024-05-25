import { Injectable } from '@nestjs/common';
import { CreatePasswordResetDto } from './dto/create-password-reset.dto';
import { ConfirmPasswordResetDto } from './dto/confirm-password-reset.dto';
import { PrismaService } from 'src/prisma/prisma.service';
import { AuthService } from 'src/auth/auth.service';
import { UsersService } from 'src/users/users.service';
import { EmailService } from 'src/email/email.service';

@Injectable()
export class PasswordResetsService {
  constructor(
    private readonly prismaService: PrismaService,
    private readonly usersService: UsersService,
    private readonly authService: AuthService,
    private readonly emailService: EmailService,
  ) {}

  async create(createPasswordResetDto: CreatePasswordResetDto) {
    const user = await this.usersService.findOneByEmail(createPasswordResetDto?.email);

    if (!user) {
      return;
    }

    // Invalidate all old password resets
    await this.prismaService.passwordReset.updateMany({
      where: { user_id: user.id },
      data: {
        status: 'completed',
      },
    });

    // Create new password reset
    const token = Math.random().toString(36).substring(2, 15);
    await this.prismaService.passwordReset.create({
      data: {
        user_id: user.id,
        token,
        created_at: new Date(),
        updated_at: new Date(),
        status: 'pending',
      },
    });

    // Send email
    await this.emailService.sendBotEmail(user.email, 'reset_password_template', {
      product_name: process?.env?.PRODUCT_NAME,
      name: user.name,
      action_url: `${process.env.FRONTEND_URL}/reset-password?token=${token}`,
    });

    return {};
  }

  async confirm(confirmPasswordResetDto: ConfirmPasswordResetDto) {
    const passwordReset = await this.findOneByToken(confirmPasswordResetDto.token);

    if (!passwordReset) {
      throw new Error('Invalid token');
    }

    const user = await this.usersService.findOne(passwordReset.user_id);

    if (!user) {
      return;
    }

    await this.authService.updatePassword(user.id, confirmPasswordResetDto.password);

    // Update all other requests with user ID to make them stale
    return this.prismaService.passwordReset.updateMany({
      where: { user_id: user.id },
      data: {
        status: 'completed',
      },
    });
  }

  findOneByToken(token: string) {
    return this.prismaService.passwordReset.findFirst({
      where: {
        token: token,
        status: 'pending',
        // created_at: {
        //   gte: new Date(new Date().getTime() - 1000 * 60 * 60), // 1 hour
        // },
      },
      orderBy: {
        created_at: 'desc',
      },
    });
  }
}
