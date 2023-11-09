package com.andredonadon.psmflux.service;

import com.andredonadon.psmflux.domain.account.AccountInput;
import com.andredonadon.psmflux.infra.repository.AccountRepository;
import com.andredonadon.psmflux.infra.repository.mapper.AccountMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;


@Service
@RequiredArgsConstructor
public class AccountService {

    private final AccountRepository accountRepository;

    public void save(AccountInput accountInput) {
        accountRepository.save(AccountMapper.toEntity(accountInput));
    }

    public String getById(Long id) {
        return "";
    }



}
