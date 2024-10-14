package de.serbroda.ragbag.services;

import de.serbroda.ragbag.dtos.page.CreatePageDto;
import de.serbroda.ragbag.mappers.PageMapper;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.PageUser;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.shared.PageRole;
import de.serbroda.ragbag.models.shared.PageVisibility;
import de.serbroda.ragbag.models.shared.SpaceRole;
import de.serbroda.ragbag.repositories.PageAccountRepository;
import de.serbroda.ragbag.repositories.PageRepository;
import de.serbroda.ragbag.repositories.SpaceRepository;
import jakarta.persistence.EntityNotFoundException;
import org.springframework.stereotype.Service;

import java.nio.file.AccessDeniedException;
import java.util.List;
import java.util.Optional;
import java.util.Set;

import static de.serbroda.ragbag.security.AuthorizationService.checkAccessAllowed;

@Service
public class PageService {

    private final SpaceRepository spaceRepository;
    private final PageRepository pageRepository;
    private final PageAccountRepository pageAccountRepository;

    public PageService(
            SpaceRepository spaceRepository,
            PageRepository pageRepository,
            PageAccountRepository pageAccountRepository) {
        this.spaceRepository = spaceRepository;
        this.pageRepository = pageRepository;
        this.pageAccountRepository = pageAccountRepository;
    }

    public Page createPage(CreatePageDto createPageDto, User user) throws AccessDeniedException {
        Space space = spaceRepository.findById(createPageDto.getSpaceId())
                .orElseThrow(() -> new EntityNotFoundException("Space with id " + createPageDto.getSpaceId() + " not found"));
        checkAccessAllowed(user, space, SpaceRole.OWNER, SpaceRole.CONTRIBUTOR);

        Page page = PageMapper.INSTANCE.map(createPageDto);
        page.setSpace(space);
        if (createPageDto.getParentPageId() != null) {
            Page parent = pageRepository.findById(createPageDto.getParentPageId())
                    .orElseThrow(() -> new EntityNotFoundException("Page with id " + createPageDto.getParentPageId() + " not found"));
            page.setParent(parent);
        }
        page.setVisibility(PageVisibility.PUBLIC);
        return createPage(page, user);
    }

    public Page createPage(Page page, User user) {
        if (page.getSpace() == null) {
            throw new IllegalArgumentException("Space must be set");
        }
        page = pageRepository.save(page);
        addAccountToPage(page, user, PageRole.OWNER);
        return page;
    }

    public PageUser addAccountToPage(Page page, User user, PageRole pageRole) {
        PageUser pageUser = new PageUser();
        pageUser.setPage(page);
        pageUser.setAccount(user);
        pageUser.setRole(pageRole);
        return pageAccountRepository.save(pageUser);
    }

    public Optional<Page> getPageById(long id) {
        return pageRepository.findById(id);
    }

    public List<Page> getPagesTreeBySpaceId(Long spaceId) {
        List<Page> pages = pageRepository.findRootPagesBySpaceId(spaceId);
        for (Page page : pages) {
            loadSubPages(page);  // rekursive Methode, um SubPages zu laden
        }
        return pages;
    }

    // Rekursive Methode, um die SubPages zu laden
    private void loadSubPages(Page page) {
        Set<Page> subPages = page.getSubPages();
        for (Page subPage : subPages) {
            loadSubPages(subPage);  // rekursiv die SubPages laden
        }
    }
}
