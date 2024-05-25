import { Inject, Injectable, Logger } from '@nestjs/common';
import { ServerClient } from 'postmark';
import { POSTMARK_SERVICE } from 'src/postmark/postmark.constans';

@Injectable()
export class EmailService {
  private readonly logger = new Logger(EmailService.name);

  constructor(@Inject(POSTMARK_SERVICE) private readonly postmarkProvider: ServerClient) {}

  sendEmail(from: string, to: string, templateAlias: string, model: object) {
    this.postmarkProvider.sendEmailWithTemplate({
      From: from,
      To: to,
      TemplateAlias: templateAlias,
      // TemplateId: 35624032,
      TemplateModel: model,
    });
  }

  sendBotEmail(to: string, templateAlias: string, model: object) {
    const { POSTMARK_FROM_EMAIL } = process.env;
    return this.sendEmail(POSTMARK_FROM_EMAIL, to, templateAlias, model);
  }
}
