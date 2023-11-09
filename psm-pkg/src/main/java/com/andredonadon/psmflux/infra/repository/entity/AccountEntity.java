package com.andredonadon.psmflux.infra.repository.entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Table(name = "account")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class AccountEntity {
    @Id
    private Long id;

    @Column(name = "document_number")
    private String documentNumber;
}
