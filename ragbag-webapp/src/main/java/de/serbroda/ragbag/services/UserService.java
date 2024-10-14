package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.User;
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

    public Optional<User> getUserByUsername(String username) {
        return accountRepository.findByUsernameIgnoreCase(username);
    }

    public Optional<User> getUserById(long userId) {
        return accountRepository.findById(userId);
    }

    @Transactional
    public User createAccount(User user) {
        if (getUserByUsername(user.getUsername()).isPresent()) {
            throw new EntityExistsException("User " + user.getUsername() + " already exists");
        }

        User entity = accountRepository.save(user);

        Space defaultSpace = new Space();
        defaultSpace.setName(
                StringUtils.capitalize(entity.getUsername().toLowerCase()) + "'s Space");
        spaceService.createSpace(defaultSpace, user);
        return entity;
    }

}
