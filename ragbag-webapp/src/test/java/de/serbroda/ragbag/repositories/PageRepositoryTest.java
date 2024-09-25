package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.shared.PageVisibility;
import de.serbroda.ragbag.repositories.base.AbstractRepositoryTest;
import de.serbroda.ragbag.repositories.base.TransactionalProfileSpringBootTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Example;
import org.springframework.data.domain.ExampleMatcher;
import org.springframework.data.jpa.repository.JpaRepository;

@SpringBootTest
@TransactionalProfileSpringBootTest
public class PageRepositoryTest extends AbstractRepositoryTest<Page> {

    @Autowired
    private SpaceRepository spaceRepository;

    @Autowired
    private PageRepository repository;

    @Override
    protected JpaRepository<Page, Long> getRepository() {
        return repository;
    }

    @Override
    protected Page getCreateEntity() {
        Space space = new Space();
        space.setName("space");
        space = spaceRepository.save(space);

        Page entity = new Page();
        entity.setSpace(space);
        entity.setName("page1");
        entity.setVisibility(PageVisibility.PUBLIC);
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(Page entity) {
        entity.setName("page2");
    }

    @Override
    protected Example<Page> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        Page entity = new Page();
        entity.setName("page1");
        return Example.of(entity, matcher);
    }
}
