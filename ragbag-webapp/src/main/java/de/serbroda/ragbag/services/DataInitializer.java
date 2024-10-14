package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.UserRole;
import de.serbroda.ragbag.models.shared.UserRoles;
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
        initializeUser(ADMIN_USERNAME, RandomUtils.randomString(10), UserRoles.ADMIN);
    }

    protected User initializeUser(String username, String password, UserRoles... roles) {
        Optional<User> adminOptional = userService.getUserByUsername(username);
        if (adminOptional.isPresent()) {
            return null;
        }

        User admin = new User();
        admin.setUsername(username);
        admin.setPassword(passwordEncoder.encode(password));
        admin.setAccountRoles(Stream.of(roles)
                .map((r) -> accountRoleRepository.findByNameIgnoreCase(r.name()))
                .collect(Collectors.toSet()));
        User user = userService.createAccount(admin);

        System.out.println("Initialized '" + ADMIN_USERNAME + "' user with password: " + password);
        return user;
    }

    private void initializeRoles() {
        for (UserRoles role : UserRoles.values()) {
            initializeRole(role);
        }
    }

    private void initializeRole(UserRoles roleEnum) {
        final String name = roleEnum.name();
        UserRole role = accountRoleRepository.findByNameIgnoreCase(roleEnum.name());
        if (role == null) {
            role = new UserRole();
            role.setName(name);
            accountRoleRepository.save(role);
        }
    }
}
