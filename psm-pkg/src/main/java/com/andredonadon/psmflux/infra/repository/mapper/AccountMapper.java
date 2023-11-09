package com.andredonadon.psmflux.infra.repository.mapper;

import com.andredonadon.psmflux.domain.account.Account;
import com.andredonadon.psmflux.domain.account.AccountInput;
import com.andredonadon.psmflux.infra.repository.entity.AccountEntity;

public class AccountMapper {

    public static AccountEntity toEntity(AccountInput accountInput) {
        return AccountEntity.builder()
                .id(Long.parseLong(accountInput.getReferenceId()))
                .documentNumber(accountInput.getDocumentNumber())
                .build();
    }

    public static Account toDTO(AccountEntity accountEntity) {
        return Account.builder()
                .id(accountEntity.getId())
                .documentNumber(accountEntity.getDocumentNumber())
                .build();
    }

}
