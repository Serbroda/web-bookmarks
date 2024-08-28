package de.serbroda.ragbag.utils;

import de.serbroda.ragbag.models.Account;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;

import java.util.Optional;

public class SessionUtils {

    public static Optional<Account> getAuthenticatedAccount() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        return Optional.ofNullable((Account) authentication.getPrincipal());
    }
}
