package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.AccountPermission;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class AccountPermissionRepositoryTest extends AbstractRepositoryTest<AccountPermission> {

    @Autowired
    private AccountPermissionRepository repository;

    @Override
    protected JpaRepository<AccountPermission, Long> getRepository() {
        return repository;
    }

    @Override
    protected AccountPermission getCreateEntity() {
        AccountPermission entity = new AccountPermission();
        entity.setName("DUMMY");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(AccountPermission entity) {
        entity.setName("EXAMPLE");
    }

    @Override
    protected Example<AccountPermission> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        AccountPermission entity = new AccountPermission();
        entity.setName("DUMMY");
        return Example.of(entity, matcher);
    }
}
