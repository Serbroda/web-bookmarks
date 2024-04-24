package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Role;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.shared.Roles;
import de.serbroda.ragbag.repositories.RoleRepository;
import de.serbroda.ragbag.repositories.UserRepository;
import de.serbroda.ragbag.utils.RandomUtils;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Arrays;
import java.util.Collections;
import java.util.HashSet;
import java.util.Optional;

@Service
public class DataInitializer {

    private final RoleRepository roleRepository;
    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;

    public DataInitializer(RoleRepository roleRepository, UserRepository userRepository, PasswordEncoder passwordEncoder) {
        this.roleRepository = roleRepository;
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
    }

    public void initialize() {
        initializeRoles();
        initializeAdmin();
    }

    private void initializeAdmin() {
        Optional<User> adminOptional = userRepository.findByUsernameIgnoreCase("admin");
        if (adminOptional.isPresent()) {
            return;
        }

        final String password = RandomUtils.randomString(10);

        User admin = new User();
        admin.setUsername("admin");
        admin.setPassword(passwordEncoder.encode(password));
        admin.setRoles(new HashSet<>(Collections.singletonList(roleRepository.findByNameIgnoreCase(Roles.ADMIN.name()))));
        userRepository.save(admin);

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
