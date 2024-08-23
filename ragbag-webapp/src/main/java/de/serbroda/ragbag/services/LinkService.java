package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.repositories.LinkRepository;
import de.serbroda.ragbag.repositories.PageRepository;
import de.serbroda.ragbag.repositories.SpaceRepository;
import org.springframework.stereotype.Service;

@Service
public class LinkService {

    private final SpaceRepository spaceRepository;
    private final PageRepository pageRepository;
    private final LinkRepository linkRepository;

    public LinkService(SpaceRepository spaceRepository, PageRepository pageRepository,
                       LinkRepository linkRepository) {
        this.spaceRepository = spaceRepository;
        this.pageRepository = pageRepository;
        this.linkRepository = linkRepository;
    }

    public Space createSpace(Space space) {
        return spaceRepository.save(space);
    }
}
