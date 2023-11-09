package com.andredonadon.psmflux.infra.messagebroker;

import com.andredonadon.psmflux.domain.account.AccountInput;
import com.andredonadon.psmflux.domain.transaction.TransactionInput;
import com.andredonadon.psmflux.service.AccountService;
import com.andredonadon.psmflux.service.TransactionService;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class Consumer {
    private final ObjectMapper objectMapper;
    private final AccountService accountService;
    private final TransactionService transactionService;

    @RabbitListener(queues = "account")
    public void processAccountMessage(String message) {
        try {
            AccountInput accountInput = objectMapper.readValue(message, AccountInput.class);

            accountService.save(accountInput);
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
    }

    @RabbitListener(queues = "transaction")
    public void processTransactionMessage(String message) {
        try {
            TransactionInput transactionInput = objectMapper.readValue(message, TransactionInput.class);

            transactionService.save(transactionInput);
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }


    }

}
