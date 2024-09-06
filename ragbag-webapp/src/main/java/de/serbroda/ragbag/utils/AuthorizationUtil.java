package de.serbroda.ragbag.utils;

import de.serbroda.ragbag.exceptions.UnauthorizedException;
import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.shared.PageRole;
import de.serbroda.ragbag.models.shared.SpaceRole;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;

import java.nio.file.AccessDeniedException;
import java.util.Arrays;
import java.util.Optional;

public class AuthorizationUtil {

    public static Optional<Account> getAuthenticatedAccount() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        return Optional.ofNullable((Account) authentication.getPrincipal());
    }

    public static Account getAuthenticatedAccountRequired() {
        return getAuthenticatedAccount().orElseThrow(() -> new UnauthorizedException("User is not authenticated"));
    }

    public static void checkAccessAllowed(Account account, Space space, SpaceRole... roles) throws AccessDeniedException {
        if (!isAccessAllowed(account, space, roles)) {
            throw new AccessDeniedException("Not allowed to access space " + space.getId());
        }
    }

    public static void checkAccessAllowed(Account account, Page page, PageRole... roles) throws AccessDeniedException {
        if (!isAccessAllowed(account, page, roles)) {
            throw new AccessDeniedException("Not allowed to access page " + page.getId());
        }
    }

    public static boolean isAccessAllowed(Account account, Space space, SpaceRole... roles) {
        return account.getSpaces().stream()
                .filter(sa -> roles.length < 1 || Arrays.asList(roles).contains(sa.getRole()))
                .map(sa -> sa.getSpace())
                .anyMatch(a -> a.equals(space));
    }

    public static boolean isAccessAllowed(Account account, Page page, PageRole... roles) {
        return account.getPages().stream()
                .filter(pa -> roles.length < 1 || Arrays.asList(roles).contains(pa.getRole()))
                .map(sa -> sa.getPage())
                .anyMatch(a -> a.equals(page));
    }
}
