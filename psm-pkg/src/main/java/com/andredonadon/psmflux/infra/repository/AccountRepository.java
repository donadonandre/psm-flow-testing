package com.andredonadon.psmflux.infra.repository;

import com.andredonadon.psmflux.infra.repository.entity.AccountEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AccountRepository extends JpaRepository<AccountEntity, Long> {

    AccountEntity getById(Long id);

}
