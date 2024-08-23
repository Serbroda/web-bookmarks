package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.controller.v1.api.UserApi;
import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.mappers.UserMapper;
import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.repositories.AccountRepository;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1")
public class UserController implements UserApi {

    private final AccountRepository accountRepository;

    public UserController(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    @Override
    public ResponseEntity<List<UserDto>> getUsers() {
        List<Account> accounts = accountRepository.findAll();
        return ResponseEntity.ok(UserMapper.INSTANCE.mapAll(accounts));
    }
}
