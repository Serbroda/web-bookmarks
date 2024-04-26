package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class AccountRepositoryTest extends AbstractRepositoryTest<Account> {

    @Autowired
    private AccountRepository repository;

    @Override
    protected JpaRepository<Account, Long> getRepository() {
        return repository;
    }

    @Override
    protected Account getCreateEntity() {
        Account entity = new Account();
        entity.setUsername("Fritz");
        entity.setPassword("s3cr3t");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(Account entity) {
        entity.setUsername("Max");
    }

    @Override
    protected Example<Account> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        Account entity = new Account();
        entity.setUsername("John");
        return Example.of(entity, matcher);
    }
}
