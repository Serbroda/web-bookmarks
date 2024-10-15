package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.repositories.UserRepository;
import jakarta.persistence.EntityExistsException;
import jakarta.transaction.Transactional;
import org.apache.commons.lang3.StringUtils;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
public class UserService {

    private final UserRepository userRepository;
    private final PageService pageService;
    private final SpaceService spaceService;

    public UserService(UserRepository userRepository, PageService pageService,
                       SpaceService spaceService) {
        this.userRepository = userRepository;
        this.pageService = pageService;
        this.spaceService = spaceService;
    }

    public Optional<User> getUserByUsername(String username) {
        return userRepository.findByUsernameIgnoreCase(username);
    }

    public Optional<User> getUserById(long userId) {
        return userRepository.findById(userId);
    }

    @Transactional
    public User createUser(User user) {
        if (getUserByUsername(user.getUsername()).isPresent()) {
            throw new EntityExistsException("User " + user.getUsername() + " already exists");
        }

        User entity = userRepository.save(user);

        Space defaultSpace = new Space();
        defaultSpace.setName(
                StringUtils.capitalize(entity.getUsername().toLowerCase()) + "'s Space");
        spaceService.createSpace(defaultSpace, user);
        return entity;
    }

}
