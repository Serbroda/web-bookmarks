package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Link;
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
public class LinkRepositoryTest extends AbstractRepositoryTest<Link> {

    @Autowired
    private SpaceRepository spaceRepository;

    @Autowired
    private PageRepository pageRepository;

    @Autowired
    private LinkRepository repository;

    @Override
    protected JpaRepository<Link, Long> getRepository() {
        return repository;
    }

    @Override
    protected Link getCreateEntity() {
        Space space = new Space();
        space.setName("space");
        space = spaceRepository.save(space);

        Page page = new Page();
        page.setSpace(space);
        page.setName("page");
        page = pageRepository.save(page);

        Link entity = new Link();
        entity.setName("link1");
        entity.setUrl("http://example.com");
        entity.setPage(page);
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(Link entity) {
        entity.setName("link2");
    }

    @Override
    protected Example<Link> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        Link entity = new Link();
        entity.setName("link1");
        return Example.of(entity, matcher);
    }
}