import {Module} from '@nestjs/common';
import {AppController} from './app.controller';
import {AppService} from './app.service';
import {ConfigModule} from "@nestjs/config";
import {ClientsModule, Transport} from "@nestjs/microservices";
import {RABBITMQ_AUTH_SERVICE, RABBITMQ_USERS_SERVICE} from "./config/microservices";
import { UsersModule } from './users/users.module';
import { AuthModule } from './auth/auth.module';
import { UsersModule } from './users/users.module';

@Module({
    imports: [
        ConfigModule.forRoot({
            isGlobal: true,
            envFilePath: '.env',
        }),
        ClientsModule.register([
            {
                name: RABBITMQ_USERS_SERVICE.NAME,
                transport: Transport.RMQ,
                options: {
                    urls: [RABBITMQ_USERS_SERVICE.URL],
                    queue: RABBITMQ_USERS_SERVICE.QUEUE,
                    queueOptions: {
                        durable: RABBITMQ_USERS_SERVICE.QUEUE_OPTIONS.DURABLE
                    }
                },
            },
            {
                name: RABBITMQ_AUTH_SERVICE.NAME,
                transport: Transport.RMQ,
                options: {
                    urls: [RABBITMQ_AUTH_SERVICE.URL],
                    queue: RABBITMQ_AUTH_SERVICE.QUEUE,
                    queueOptions: {
                        durable: RABBITMQ_AUTH_SERVICE.QUEUE_OPTIONS.DURABLE
                    }
                },
            },
        ]),
        UsersModule,
        AuthModule
    ],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {
}
