import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { PrismaService } from 'src/prisma/prisma.service';
import { UpdateUserPasswordDto } from './dto/update-user-password.dto';

@Injectable()
export class UsersService {
  constructor(private prismaService: PrismaService) {}

  async create(createUserDto: CreateUserDto) {
    return this.prismaService.user.create({
      data: {
        ...createUserDto,
      },
    });
  }

  findAll() {
    return `This action returns all users`;
  }

  findOne(id: string) {
    return this.prismaService.user.findUnique({
      where: { id },
    });
  }

  findOneByEmail(email: string) {
    return this.prismaService.user.findUnique({
      where: { email },
    });
  }

  update(id: string, updateUserDto: UpdateUserDto) {
    return this.prismaService.user.update({
      where: { id },
      data: {
        ...updateUserDto,
      },
    });
  }

  updatePassword(id: string, updateUserPasswordDto: UpdateUserPasswordDto) {
    return this.prismaService.user.update({
      where: { id },
      data: {
        ...updateUserPasswordDto,
      },
    });
  }

  remove(id: number) {
    return `This action removes a #${id} user`;
  }
}
