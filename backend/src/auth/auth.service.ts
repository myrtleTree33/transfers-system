import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
  ) {}

  async validateUser(email: string, password: string): Promise<any> {
    const user = await this.usersService.findOneByEmail(email);
    if (user && (await this.comparePassword(password, user.password))) {
      const { password, ...rest } = user;
      return rest;
    }
    return null;
  }

  async loginPasswordUser(user: any) {
    const payload = { email: user.email, sub: user.id };
    return {
      access_token: this.jwtService.sign(payload, {
        expiresIn: '60m',
        secret: process.env.JWT_SECRET,
      }),
    };
  }

  async signupPasswordUser(user: any) {
    const hashedPassword = await this.hashPassword(user.password);
    return this.usersService.create({ ...user, password: hashedPassword });
  }

  async updatePassword(id: string, password: string) {
    const hashedPassword = await this.hashPassword(password);
    return this.usersService.updatePassword(id, { password: hashedPassword });
  }

  async comparePassword(password, hash) {
    const isMatch = await bcrypt.compare(password, hash);
    return isMatch;
  }

  async hashPassword(password) {
    const hash = await bcrypt.hash(password, 10);
    return hash;
  }
}
