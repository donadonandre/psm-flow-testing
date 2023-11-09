package com.andredonadon.psmflux.service;

import com.andredonadon.psmflux.domain.operationtype.Operation;
import com.andredonadon.psmflux.domain.transaction.TransactionInput;
import com.andredonadon.psmflux.infra.repository.AccountRepository;
import com.andredonadon.psmflux.infra.repository.TransactionRepository;
import com.andredonadon.psmflux.infra.repository.entity.AccountEntity;
import com.andredonadon.psmflux.infra.repository.entity.TransactionEntity;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class TransactionService {

    private TransactionRepository transactionRepository;

    private AccountRepository accountRepository;

    public void save(TransactionInput transactionInput) {
        AccountEntity accountEntity = accountRepository.getById(transactionInput.getAccountId());

        TransactionEntity entity = TransactionEntity.builder()
                .id(UUID.randomUUID())
                .operationType(transactionInput.getOperationType())
                .amount(transactionInput.getAmount())
                .eventDate(LocalDateTime.now())
                .account(accountEntity)
                .build();

        transactionRepository.save(entity);
    }



}
