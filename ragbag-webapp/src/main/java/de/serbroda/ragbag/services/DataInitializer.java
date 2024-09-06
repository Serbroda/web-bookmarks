package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.AccountRole;
import de.serbroda.ragbag.models.shared.AccountRoles;
import de.serbroda.ragbag.repositories.AccountRoleRepository;
import de.serbroda.ragbag.utils.RandomUtils;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.HashSet;
import java.util.Optional;

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
        initializeAdmin();
    }

    private void initializeAdmin() {
        Optional<Account> adminOptional = userService.getUserByUsername(ADMIN_USERNAME);
        if (adminOptional.isPresent()) {
            return;
        }

        final String password = RandomUtils.randomString(10);

        Account admin = new Account();
        admin.setUsername(ADMIN_USERNAME);
        admin.setPassword(passwordEncoder.encode(password));
        admin.setAccountRoles(new HashSet<>(
                Collections.singletonList(
                        accountRoleRepository.findByNameIgnoreCase(AccountRoles.ADMIN.name()))));
        userService.createAccount(admin);

        System.out.println("Initialized '" + ADMIN_USERNAME + "' user with password: " + password);
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