package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Role;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class RoleRepositoryTest extends AbstractRepositoryTest<Role> {

    @Autowired
    private RoleRepository repository;

    @Override
    protected JpaRepository<Role, Long> getRepository() {
        return repository;
    }

    @Override
    protected Role getCreateEntity() {
        Role entity = new Role();
        entity.setName("ADMIN");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(Role entity) {
        entity.setName("USER");
    }

    @Override
    protected Example<Role> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        Role entity = new Role();
        entity.setName("ADMIN");
        return Example.of(entity, matcher);
    }
}
