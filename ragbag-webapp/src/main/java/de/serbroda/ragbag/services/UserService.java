package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.repositories.AccountRepository;
import jakarta.persistence.EntityExistsException;
import java.util.Optional;
import org.apache.commons.lang3.StringUtils;
import org.springframework.stereotype.Service;

@Service
public class UserService {

    private final AccountRepository accountRepository;
    private final LinkService linkService;

    public UserService(AccountRepository accountRepository, LinkService linkService) {
        this.accountRepository = accountRepository;
        this.linkService = linkService;
    }

    public Optional<Account> getUserByUsername(String username) {
        return accountRepository.findByUsernameIgnoreCase(username);
    }

    public Account createAccount(Account account) {
        if (getUserByUsername(account.getUsername()).isPresent()) {
            throw new EntityExistsException("User " + account.getUsername() + " already exists");
        }

        Account entity = accountRepository.save(account);

        Space defaultSpace = new Space();
        defaultSpace.setName(
            StringUtils.capitalize(entity.getUsername().toLowerCase()) + "'s Space");
        linkService.createSpace(defaultSpace);

        return entity;
    }

}
