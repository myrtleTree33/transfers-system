import { Controller, Post, Body } from '@nestjs/common';
import { PasswordResetsService } from './password-resets.service';
import { CreatePasswordResetDto } from './dto/create-password-reset.dto';
import { ConfirmPasswordResetDto } from './dto/confirm-password-reset.dto';

@Controller('password-resets')
export class PasswordResetsController {
  constructor(private readonly passwordResetsService: PasswordResetsService) {}

  @Post()
  create(@Body() createPasswordResetDto: CreatePasswordResetDto) {
    return this.passwordResetsService.create(createPasswordResetDto);
  }

  @Post('/confirm')
  confirm(@Body() confirmPasswordResetDto: ConfirmPasswordResetDto) {
    return this.passwordResetsService.confirm(confirmPasswordResetDto);
  }
}
