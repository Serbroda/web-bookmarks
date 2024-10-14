package de.serbroda.ragbag.security;

import de.serbroda.ragbag.exceptions.UnauthorizedException;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.shared.PageRole;
import de.serbroda.ragbag.models.shared.SpaceRole;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;

import java.nio.file.AccessDeniedException;
import java.util.Arrays;
import java.util.Optional;

public class AuthorizationService {

    private AuthorizationService() {

    }

    public static Optional<User> getAuthenticatedAccount() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        return Optional.ofNullable((User) authentication.getPrincipal());
    }

    public static User getAuthenticatedAccountRequired() {
        return getAuthenticatedAccount().orElseThrow(() -> new UnauthorizedException("User is not authenticated"));
    }

    public static void checkAccessAllowed(Space space, SpaceRole... roles) throws AccessDeniedException {
        checkAccessAllowed(getAuthenticatedAccountRequired(), space, roles);
    }

    public static void checkAccessAllowed(Page page, PageRole... roles) throws AccessDeniedException {
        checkAccessAllowed(getAuthenticatedAccountRequired(), page, roles);
    }

    public static void checkAccessAllowed(User user, Space space, SpaceRole... roles) throws AccessDeniedException {
        if (!isAccessAllowed(user, space, roles)) {
            throw new AccessDeniedException("Not allowed to access space " + space.getId());
        }
    }

    public static void checkAccessAllowed(User user, Page page, PageRole... roles) throws AccessDeniedException {
        if (!isAccessAllowed(user, page, roles)) {
            throw new AccessDeniedException("Not allowed to access page " + page.getId());
        }
    }

    public static boolean isAccessAllowed(User user, Space space, SpaceRole... roles) {
        return user.getSpaces().stream()
                .filter(sa -> roles.length < 1 || Arrays.asList(roles).contains(sa.getRole()))
                .map(sa -> sa.getSpace())
                .anyMatch(a -> a.equals(space));
    }

    public static boolean isAccessAllowed(User user, Page page, PageRole... roles) {
        return user.getPages().stream()
                .filter(pa -> roles.length < 1 || Arrays.asList(roles).contains(pa.getRole()))
                .map(sa -> sa.getPage())
                .anyMatch(a -> a.equals(page));
    }
}
