import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import helmet from 'helmet';
import { UnknownExceptionsInterceptor } from './interceptors/unknown-exception/unknown-exception.interceptor';
import { NotFoundInterceptor } from './interceptors/not-found/not-found.interceptor';
import { LoggingInterceptor } from './logging/logging.interceptor';

const { PROJECT_NAME, CORS_URLS } = process.env;

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.setGlobalPrefix('api');

  const config = new DocumentBuilder()
    .setTitle(`${PROJECT_NAME}`)
    .setDescription(`${PROJECT_NAME} API documentation`)
    .setVersion('0.1')
    .build();

  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('docs', app, document);

  // Apply global middleware
  app.use(helmet());

  // Set global interceptors
  app.useGlobalInterceptors(
    new UnknownExceptionsInterceptor(),
    new NotFoundInterceptor(),
    new LoggingInterceptor(),
  );

  // Enable CORS
  const corsUrls = CORS_URLS?.split(',')?.map((url) => url?.trim()) ?? [];
  app.enableCors({
    origin: corsUrls,
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    preflightContinue: false,
    optionsSuccessStatus: 204,
  });

  await app.listen(3000);
}

bootstrap();
