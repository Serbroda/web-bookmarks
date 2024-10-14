package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class UserRepositoryTest extends AbstractRepositoryTest<User> {

    @Autowired
    private AccountRepository repository;

    @Override
    protected JpaRepository<User, Long> getRepository() {
        return repository;
    }

    @Override
    protected User getCreateEntity() {
        User entity = new User();
        entity.setUsername("Fritz");
        entity.setPassword("s3cr3t");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(User entity) {
        entity.setUsername("Max");
    }

    @Override
    protected Example<User> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        User entity = new User();
        entity.setUsername("John");
        return Example.of(entity, matcher);
    }
}
