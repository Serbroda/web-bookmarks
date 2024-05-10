package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceAccount;
import de.serbroda.ragbag.models.keys.SpaceAccountKey;
import de.serbroda.ragbag.models.shared.SpaceRole;
import de.serbroda.ragbag.repositories.AccountRepository;
import de.serbroda.ragbag.repositories.SpaceAccountRepository;
import jakarta.persistence.EntityExistsException;
import jakarta.transaction.Transactional;
import java.util.Optional;
import org.apache.commons.lang3.StringUtils;
import org.springframework.stereotype.Service;

@Service
public class UserService {

    private final AccountRepository accountRepository;
    private final LinkService linkService;
    private final SpaceService spaceService;

    public UserService(AccountRepository accountRepository, LinkService linkService,
        SpaceService spaceService) {
        this.accountRepository = accountRepository;
        this.linkService = linkService;
        this.spaceService = spaceService;
    }

    public Optional<Account> getUserByUsername(String username) {
        return accountRepository.findByUsernameIgnoreCase(username);
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
