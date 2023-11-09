package com.andredonadon.psmflux.configuration;

import org.springframework.amqp.core.*;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RabbitMQConfig {

    @Bean
    public Exchange creationExchange() {
        return new TopicExchange("creation");
    }

    @Bean
    public Queue accountQueue() {
        return new Queue("account");
    }

    @Bean
    public Queue transactionQueue() {
        return new Queue("transaction");
    }

    @Bean
    public Binding accountBinding(Queue accountQueue, Exchange creationExchange) {
        return BindingBuilder.bind(accountQueue).to(creationExchange).with("account").noargs();
    }

    @Bean
    public Binding transactionBinding(Queue transactionQueue, Exchange creationExchange) {
        return BindingBuilder.bind(transactionQueue).to(creationExchange).with("transaction").noargs();
    }



}
