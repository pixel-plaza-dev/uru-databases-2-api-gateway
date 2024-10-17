import {NestFactory} from '@nestjs/core';
import {AppModule} from './app.module';
import {API_GATEWAY} from "./config/api-gateway";

async function bootstrap() {
    const app = await NestFactory.create(AppModule);
    await app.listen(API_GATEWAY.PORT);
}

bootstrap();
