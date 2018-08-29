package org.rls.rest;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableAutoConfiguration
public class RLSRestApplication {
    public static void main(String[] args) {
        SpringApplication.run(RLSRestApplication.class, args);
    }
}
