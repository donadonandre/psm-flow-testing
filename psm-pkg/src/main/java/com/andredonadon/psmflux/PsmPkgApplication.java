package com.andredonadon.psmflux;

import org.springframework.amqp.rabbit.annotation.EnableRabbit;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableRabbit
public class PsmPkgApplication {

    public static void main(String[] args) {
        SpringApplication.run(PsmPkgApplication.class, args);
    }

}
