package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.AccountRole;
import de.serbroda.ragbag.models.shared.AccountRoles;
import de.serbroda.ragbag.repositories.AccountRoleRepository;
import de.serbroda.ragbag.utils.RandomUtils;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Optional;
import java.util.stream.Collectors;
import java.util.stream.Stream;

@Service
public class DataInitializer {

    private static final String ADMIN_USERNAME = "admin";

    private final AccountRoleRepository accountRoleRepository;
    private final UserService userService;
    private final PasswordEncoder passwordEncoder;

    public DataInitializer(AccountRoleRepository accountRoleRepository, UserService userService,
                           PasswordEncoder passwordEncoder) {
        this.accountRoleRepository = accountRoleRepository;
        this.userService = userService;
        this.passwordEncoder = passwordEncoder;
    }

    public void initialize() {
        initializeRoles();
        initializeUser(ADMIN_USERNAME, RandomUtils.randomString(10), AccountRoles.ADMIN);
    }

    protected Account initializeUser(String username, String password, AccountRoles... roles) {
        Optional<Account> adminOptional = userService.getUserByUsername(username);
        if (adminOptional.isPresent()) {
            return null;
        }

        Account admin = new Account();
        admin.setUsername(username);
        admin.setPassword(passwordEncoder.encode(password));
        admin.setAccountRoles(Stream.of(roles)
                .map((r) -> accountRoleRepository.findByNameIgnoreCase(r.name()))
                .collect(Collectors.toSet()));
        Account account = userService.createAccount(admin);

        System.out.println("Initialized '" + ADMIN_USERNAME + "' user with password: " + password);
        return account;
    }

    private void initializeRoles() {
        for (AccountRoles role : AccountRoles.values()) {
            initializeRole(role);
        }
    }

    private void initializeRole(AccountRoles roleEnum) {
        final String name = roleEnum.name();
        AccountRole role = accountRoleRepository.findByNameIgnoreCase(roleEnum.name());
        if (role == null) {
            role = new AccountRole();
            role.setName(name);
            accountRoleRepository.save(role);
        }
    }
}
