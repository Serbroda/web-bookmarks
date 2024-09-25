package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.repositories.AccountRepository;
import jakarta.persistence.EntityExistsException;
import jakarta.transaction.Transactional;
import org.apache.commons.lang3.StringUtils;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
public class UserService {

    private final AccountRepository accountRepository;
    private final PageService pageService;
    private final SpaceService spaceService;

    public UserService(AccountRepository accountRepository, PageService pageService,
                       SpaceService spaceService) {
        this.accountRepository = accountRepository;
        this.pageService = pageService;
        this.spaceService = spaceService;
    }

    public Optional<Account> getUserByUsername(String username) {
        return accountRepository.findByUsernameIgnoreCase(username);
    }

    public Optional<Account> getUserById(long userId) {
        return accountRepository.findById(userId);
    }

    @Transactional
    public Account createAccount(Account account) {
        if (getUserByUsername(account.getUsername()).isPresent()) {
            throw new EntityExistsException("User " + account.getUsername() + " already exists");
        }

        Account entity = accountRepository.save(account);

        Space defaultSpace = new Space();
        defaultSpace.setName(
                StringUtils.capitalize(entity.getUsername().toLowerCase()) + "'s Space");
        spaceService.createSpace(defaultSpace, account);
        return entity;
    }

}
