package com.andredonadon.psmflux.domain.transaction;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.math.BigDecimal;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class TransactionInput {

    @JsonProperty("account_id")
    private Long accountId;

    @JsonProperty("operation_type")
    private Long operationType;

    private BigDecimal amount;

}
