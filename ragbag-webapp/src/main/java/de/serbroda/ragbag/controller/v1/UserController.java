package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.mappers.UserMapper;
import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.repositories.AccountRepository;
import de.serbroda.ragbag.security.AuthorizationService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1/users")
public class UserController {

    private final AccountRepository accountRepository;

    public UserController(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    @PreAuthorize("hasAnyRole('ADMIN')")
    @GetMapping
    public ResponseEntity<List<UserDto>> getUsers() {
        List<Account> accounts = accountRepository.findAll();
        return ResponseEntity.ok(UserMapper.INSTANCE.mapAll(accounts));
    }

    @GetMapping("/me")
    public ResponseEntity<UserDto> authenticatedUser() {
        return AuthorizationService.getAuthenticatedAccount()
                .map(UserMapper.INSTANCE::map)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.status(HttpStatus.UNAUTHORIZED).build());
    }

}
