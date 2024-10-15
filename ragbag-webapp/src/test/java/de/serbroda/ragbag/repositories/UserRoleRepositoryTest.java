package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.UserRole;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class UserRoleRepositoryTest extends AbstractRepositoryTest<UserRole> {

    @Autowired
    private UserRoleRepository repository;

    @Override
    protected JpaRepository<UserRole, Long> getRepository() {
        return repository;
    }

    @Override
    protected UserRole getCreateEntity() {
        UserRole entity = new UserRole();
        entity.setName("DUMMY");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(UserRole entity) {
        entity.setName("EXAMPLE");
    }

    @Override
    protected Example<UserRole> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        UserRole entity = new UserRole();
        entity.setName("DUMMY");
        return Example.of(entity, matcher);
    }
}
