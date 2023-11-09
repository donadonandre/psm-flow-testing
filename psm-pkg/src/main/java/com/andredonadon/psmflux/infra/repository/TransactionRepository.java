package com.andredonadon.psmflux.infra.repository;

import com.andredonadon.psmflux.infra.repository.entity.TransactionEntity;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface TransactionRepository extends JpaRepository<TransactionEntity, UUID> {



}
