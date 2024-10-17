// Microservices
import * as process from "node:process";

export interface RabbitMQConfig {
    readonly NAME: string;
    readonly URL: string;
    readonly QUEUE: string;
    readonly QUEUE_OPTIONS: {
        readonly DURABLE: boolean;
    };
}

export const RABBITMQ_AUTH_SERVICE: RabbitMQConfig = {
    NAME: 'auth-service',
    URL: process.env.RABBITMQ_AUTH_URL,
    QUEUE: process.env.RABBITMQ_AUTH_QUEUE,
    QUEUE_OPTIONS: {DURABLE: true},
} as const;

export const RABBITMQ_USERS_SERVICE: RabbitMQConfig = {
    NAME: 'users-service',
    URL: process.env.RABBITMQ_USERS_URL,
    QUEUE: process.env.RABBITMQ_USERS_QUEUE,
    QUEUE_OPTIONS: {DURABLE: true},
} as const;