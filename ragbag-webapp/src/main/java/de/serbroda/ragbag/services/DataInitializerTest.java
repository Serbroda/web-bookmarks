package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.shared.UserRoles;
import de.serbroda.ragbag.repositories.UserRoleRepository;
import org.hibernate.Hibernate;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class DataInitializerTest extends DataInitializer {

    private final PageService pageService;

    public DataInitializerTest(
            UserRoleRepository userRoleRepository,
            UserService userService,
            PasswordEncoder passwordEncoder, PageService pageService) {
        super(userRoleRepository, userService, passwordEncoder);
        this.pageService = pageService;
    }

    @Transactional
    @Override
    public void initialize() {
        super.initialize();
        super.initializeUser("max", "s3cr3t", UserRoles.USER);
        User user = super.initializeUser("john", "s3cr3t", UserRoles.USER);

        Hibernate.initialize(user.getSpaces());

        user.getSpaces().stream().findFirst().ifPresent(space -> {
            Page page1 = new Page();
            page1.setName("Page 1");
            page1 = pageService.createPage(page1, user);

            Page page2 = new Page();
            page2.setName("Page 2");
            page2 = pageService.createPage(page2, user);

            Page page21 = new Page();
            page21.setName("Page 2.1");
            page21.setParent(page2);
            page21 = pageService.createPage(page21, user);

            Page page22 = new Page();
            page22.setName("Page 2.2");
            page22.setParent(page2);
            page22 = pageService.createPage(page22, user);

            Page page221 = new Page();
            page221.setName("Page 2.2.1");
            page221.setParent(page22);
            page221 = pageService.createPage(page221, user);
        });
    }
}
