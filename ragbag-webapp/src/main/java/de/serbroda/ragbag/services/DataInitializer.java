package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Role;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.shared.Roles;
import de.serbroda.ragbag.repositories.RoleRepository;
import de.serbroda.ragbag.utils.RandomUtils;
import java.util.Collections;
import java.util.HashSet;
import java.util.Optional;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class DataInitializer {

    private final RoleRepository roleRepository;
    private final UserService userService;
    private final PasswordEncoder passwordEncoder;

    public DataInitializer(RoleRepository roleRepository, UserService userService,
        PasswordEncoder passwordEncoder) {
        this.roleRepository = roleRepository;
        this.userService = userService;
        this.passwordEncoder = passwordEncoder;
    }

    public void initialize() {
        initializeRoles();
        initializeAdmin();
    }

    private void initializeAdmin() {
        Optional<User> adminOptional = userService.getUserByUsername("admin");
        if (adminOptional.isPresent()) {
            return;
        }

        final String password = RandomUtils.randomString(10);

        User admin = new User();
        admin.setUsername("admin");
        admin.setPassword(passwordEncoder.encode(password));
        admin.setRoles(new HashSet<>(
            Collections.singletonList(roleRepository.findByNameIgnoreCase(Roles.ADMIN.name()))));
        userService.createUser(admin);

        System.out.println("Initialized admin user with password: " + password);
    }

    private void initializeRoles() {
        for (Roles role : Roles.values()) {
            initializeRole(role.name());
        }
    }

    private void initializeRole(String name) {
        Role role = roleRepository.findByNameIgnoreCase(name);
        if (role == null) {
            role = new Role();
            role.setName(name);
            roleRepository.save(role);
        }
    }
}
