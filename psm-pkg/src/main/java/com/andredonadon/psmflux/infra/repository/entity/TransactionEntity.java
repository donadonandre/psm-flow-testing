package com.andredonadon.psmflux.infra.repository.entity;

import jakarta.persistence.*;
import lombok.*;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

@Entity
@Table(name = "transaction")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class TransactionEntity {

    @Id
    private UUID id;

    @ManyToOne
    private AccountEntity account;

    @Column(name = "operation_type")
    private Long operationType;

    private BigDecimal amount;

    private LocalDateTime eventDate;
}
