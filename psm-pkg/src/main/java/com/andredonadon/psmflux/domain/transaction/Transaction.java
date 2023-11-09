package com.andredonadon.psmflux.domain.transaction;

import com.andredonadon.psmflux.domain.account.Account;
import com.andredonadon.psmflux.domain.operationtype.Operation;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Transaction {

    private UUID id;

    private Account account;

    private Operation operation;

    private BigDecimal amount;

    private LocalDateTime eventDate;

}
