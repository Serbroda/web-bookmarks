package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Group;
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
public class GroupRepositoryTest extends AbstractRepositoryTest<Group> {

    @Autowired
    private SpaceRepository spaceRepository;

    @Autowired
    private GroupRepository repository;

    @Override
    protected JpaRepository<Group, Long> getRepository() {
        return repository;
    }

    @Override
    protected Group getCreateEntity() {
        Space space = new Space();
        space.setName("space");
        space = spaceRepository.save(space);

        Group entity = new Group();
        entity.setSpace(space);
        entity.setName("group1");
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(Group entity) {
        entity.setName("group2");
    }

    @Override
    protected Example<Group> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        Group entity = new Group();
        entity.setName("group1");
        return Example.of(entity, matcher);
    }
}
