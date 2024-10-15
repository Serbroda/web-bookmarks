package de.serbroda.ragbag.services;

import de.serbroda.ragbag.exceptions.UnauthorizedException;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.shared.SpaceRole;
import de.serbroda.ragbag.repositories.SpaceRepository;
import jakarta.persistence.EntityNotFoundException;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import java.util.Arrays;
import java.util.Optional;

@Service
public class AuthorizationService {

    private final SpaceRepository spaceRepository;

    public AuthorizationService(SpaceRepository spaceRepository) {
        this.spaceRepository = spaceRepository;
    }

    public static Optional<User> getAuthenticatedUser() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        return Optional.ofNullable((User) authentication.getPrincipal());
    }

    public static User getAuthenticatedUserRequired() {
        return getAuthenticatedUser().orElseThrow(() -> new UnauthorizedException("User is not authenticated"));
    }

    public boolean hasAccessToSpace(Authentication authentication, long spaceId, SpaceRole... roles) {
        Optional<Space> space = spaceRepository.findById(spaceId);
        if (!space.isPresent()) {
            throw new EntityNotFoundException("Space with id " + spaceId + " not found");
        }

        User user = (User) authentication.getPrincipal();
        return isAccessAllowed(user, space.get(), roles);
    }

    public static boolean isAccessAllowed(User user, Space space, SpaceRole... roles) {
        return space.getUsers().stream()
                .filter(s -> roles.length < 1 || Arrays.asList(roles).contains(s.getRole()))
                .map(s -> s.getUser())
                .anyMatch(a -> a.equals(user));
    }
}
