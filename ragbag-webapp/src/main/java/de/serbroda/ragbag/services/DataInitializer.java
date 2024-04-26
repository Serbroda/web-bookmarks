package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.AccountPermission;
import de.serbroda.ragbag.models.AccountRole;
import de.serbroda.ragbag.models.shared.AccountPermissions;
import de.serbroda.ragbag.models.shared.AccountRoles;
import de.serbroda.ragbag.repositories.AccountPermissionRepository;
import de.serbroda.ragbag.repositories.AccountRoleRepository;
import de.serbroda.ragbag.utils.RandomUtils;
import java.util.Collections;
import java.util.HashSet;
import java.util.Optional;
import java.util.Set;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class DataInitializer {

    private final AccountRoleRepository accountRoleRepository;
    private final AccountPermissionRepository accountPermissionRepository;
    private final UserService userService;
    private final PasswordEncoder passwordEncoder;

    public DataInitializer(AccountRoleRepository accountRoleRepository,
        AccountPermissionRepository accountPermissionRepository, UserService userService,
        PasswordEncoder passwordEncoder) {
        this.accountRoleRepository = accountRoleRepository;
        this.accountPermissionRepository = accountPermissionRepository;
        this.userService = userService;
        this.passwordEncoder = passwordEncoder;
    }

    public void initialize() {
        initializePermissions();
        initializeRoles();
        initializeAdmin();
    }

    private void initializeAdmin() {
        Optional<Account> adminOptional = userService.getUserByUsername("admin");
        if (adminOptional.isPresent()) {
            return;
        }

        final String password = RandomUtils.randomString(10);

        Account admin = new Account();
        admin.setUsername("admin");
        admin.setPassword(passwordEncoder.encode(password));
        admin.setRoles(new HashSet<>(
            Collections.singletonList(
                accountRoleRepository.findByNameIgnoreCase(AccountRoles.ADMIN.name()))));
        userService.createAccount(admin);

        System.out.println("Initialized admin user with password: " + password);
    }

    private void initializePermissions() {
        for (AccountPermissions permission : AccountPermissions.values()) {
            initializePermission(permission);
        }
    }

    private void initializeRoles() {
        for (AccountRoles role : AccountRoles.values()) {
            initializeRole(role);
        }
    }

    private void initializePermission(AccountPermissions permissionEnum) {
        final String name = permissionEnum.name();
        AccountPermission permission = accountPermissionRepository.findByNameIgnoreCase(name);
        if (permission == null) {
            permission = new AccountPermission();
            permission.setName(name);
            accountPermissionRepository.save(permission);
        }
    }

    private void initializeRole(AccountRoles roleEnum) {
        final String name = roleEnum.name();
        AccountRole role = accountRoleRepository.findByNameIgnoreCase(roleEnum.name());
        if (role == null) {
            Set<AccountPermission> permissions = accountPermissionRepository.findByNameIn(
                roleEnum.getPermissions().stream().map(AccountPermissions::name).toList());

            role = new AccountRole();
            role.setName(name);
            role.setPermissions(permissions);
            accountRoleRepository.save(role);
        }
    }
}
