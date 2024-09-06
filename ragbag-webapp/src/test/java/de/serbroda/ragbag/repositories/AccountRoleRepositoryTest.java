package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.AccountRole;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class AccountRoleRepositoryTest extends AbstractRepositoryTest<AccountRole> {

    @Autowired
    private AccountRoleRepository repository;

    @Override
    protected JpaRepository<AccountRole, Long> getRepository() {
        return repository;
    }

    @Override
    protected AccountRole getCreateEntity() {
        AccountRole entity = new AccountRole();
        entity.setName("DUMMY");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(AccountRole entity) {
        entity.setName("EXAMPLE");
    }

    @Override
    protected Example<AccountRole> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        AccountRole entity = new AccountRole();
        entity.setName("DUMMY");
        return Example.of(entity, matcher);
    }
}