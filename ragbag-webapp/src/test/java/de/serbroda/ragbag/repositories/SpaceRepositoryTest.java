package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Role;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class SpaceRepositoryTest extends AbstractRepositoryTest<Space> {

    @Autowired
    private SpaceRepository repository;

    @Override
    protected JpaRepository<Space, Long> getRepository() {
        return repository;
    }

    @Override
    protected Space getCreateEntity() {
        Space entity = new Space();
        entity.setName("space1");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(Space entity) {
        entity.setName("space2");
    }

    @Override
    protected Example<Space> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        Space entity = new Space();
        entity.setName("space1");
        return Example.of(entity, matcher);
    }
}
