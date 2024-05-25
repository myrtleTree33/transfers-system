import { Test, TestingModule } from '@nestjs/testing';
import { PasswordResetsController } from './password-resets.controller';
import { PasswordResetsService } from './password-resets.service';

describe('PasswordResetsController', () => {
  let controller: PasswordResetsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [PasswordResetsController],
      providers: [PasswordResetsService],
    }).compile();

    controller = module.get<PasswordResetsController>(PasswordResetsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
