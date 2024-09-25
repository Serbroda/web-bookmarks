package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Bookmark;
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
public class BookmarkRepositoryTest extends AbstractRepositoryTest<Bookmark> {

    @Autowired
    private SpaceRepository spaceRepository;

    @Autowired
    private PageRepository pageRepository;

    @Autowired
    private BookmarkRepository repository;

    @Override
    protected JpaRepository<Bookmark, Long> getRepository() {
        return repository;
    }

    @Override
    protected Bookmark getCreateEntity() {
        Space space = new Space();
        space.setName("space");
        space = spaceRepository.save(space);

        Page page = new Page();
        page.setSpace(space);
        page.setName("page");
        page.setVisibility(PageVisibility.PUBLIC);
        page = pageRepository.save(page);

        Bookmark entity = new Bookmark();
        entity.setName("bookmark1");
        entity.setUrl("http://example.com");
        entity.setPage(page);
        return entity;
    }

    @Override
    protected void modifyUpdateEntity(Bookmark entity) {
        entity.setName("bookmark2");
    }

    @Override
    protected Example<Bookmark> getExample() {
        ExampleMatcher matcher = ExampleMatcher.matchingAny()
                .withMatcher("name", ExampleMatcher.GenericPropertyMatchers.contains().ignoreCase());
        Bookmark entity = new Bookmark();
        entity.setName("bookmark1");
        return Example.of(entity, matcher);
    }
}
