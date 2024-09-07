package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.repositories.BookmarkRepository;
import de.serbroda.ragbag.repositories.PageRepository;
import de.serbroda.ragbag.repositories.SpaceRepository;
import org.springframework.stereotype.Service;

@Service
public class BookmarkService {

    private final SpaceRepository spaceRepository;
    private final PageRepository pageRepository;
    private final BookmarkRepository bookmarkRepository;

    public BookmarkService(SpaceRepository spaceRepository, PageRepository pageRepository,
                           BookmarkRepository bookmarkRepository) {
        this.spaceRepository = spaceRepository;
        this.pageRepository = pageRepository;
        this.bookmarkRepository = bookmarkRepository;
    }

    public Space createSpace(Space space) {
        return spaceRepository.save(space);
    }
}
