package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.repositories.UserRepository;
import jakarta.persistence.EntityExistsException;
import java.util.Optional;
import org.springframework.stereotype.Service;

@Service
public class UserService {

    private final UserRepository userRepository;
    private final LinkService linkService;

    public UserService(UserRepository userRepository, LinkService linkService) {
        this.userRepository = userRepository;
        this.linkService = linkService;
    }

    public Optional<User> getUserByUsername(String username) {
        return userRepository.findByUsernameIgnoreCase(username);
    }

    public User createUser(User user) {
        if (getUserByUsername(user.getUsername()).isPresent()) {
            throw new EntityExistsException("User " + user.getUsername() + " already exists");
        }

        User entity = userRepository.save(user);

        Space defaultSpace = new Space();
        defaultSpace.setName(entity.getUsername() + "'s Space");
        defaultSpace = linkService.createSpace(defaultSpace);

        return entity;
    }

}
