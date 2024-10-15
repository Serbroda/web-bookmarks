package de.serbroda.ragbag.services;

import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.mappers.SpaceMapper;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceUser;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.shared.PageVisibility;
import de.serbroda.ragbag.models.shared.SpaceRole;
import de.serbroda.ragbag.repositories.SpaceAccountRepository;
import de.serbroda.ragbag.repositories.SpaceRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class SpaceService {

    private final SpaceRepository spaceRepository;
    private final SpaceAccountRepository spaceAccountRepository;
    private final PageService pageService;

    public SpaceService(SpaceRepository spaceRepository,
                        SpaceAccountRepository spaceAccountRepository, PageService pageService) {
        this.spaceRepository = spaceRepository;
        this.spaceAccountRepository = spaceAccountRepository;
        this.pageService = pageService;
    }

    public Space createSpace(CreateSpaceDto createSpaceDto, User user) {
        Space space = SpaceMapper.INSTANCE.map(createSpaceDto);
        return createSpace(space, user);
    }

    public Space createSpace(Space space, User user) {
        space = spaceRepository.save(space);
        addAccountToSpace(space, user, SpaceRole.OWNER);

        Page defaultPage = new Page();
        defaultPage.setSpace(space);
        defaultPage.setName("Default");
        defaultPage.setVisibility(PageVisibility.PUBLIC);
        pageService.createPage(defaultPage, user);
        return space;
    }

    public List<Space> getSpaces(User user) {
        return spaceRepository.findAllByAccount(user);
    }

    public Optional<Space> getSpaceById(long id) {
        return spaceRepository.findById(id);
    }

    public SpaceUser addAccountToSpace(Space space, User user, SpaceRole role) {
        SpaceUser spaceUser = new SpaceUser();
        spaceUser.setSpace(space);
        spaceUser.setAccount(user);
        spaceUser.setRole(role);
        return spaceAccountRepository.save(spaceUser);
    }

    public void removeAccountFromSpace(Space space, User user) {
        Optional<SpaceUser> spaceAccount = spaceAccountRepository.findBySpaceAndAccount(space, user);
        spaceAccount.ifPresent(value -> spaceAccountRepository.delete(value));
    }
}
