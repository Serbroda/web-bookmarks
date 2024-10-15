package de.serbroda.ragbag.services;

import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.mappers.SpaceMapper;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceUser;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.shared.PageVisibility;
import de.serbroda.ragbag.models.shared.SpaceRole;
import de.serbroda.ragbag.repositories.SpaceUserRepository;
import de.serbroda.ragbag.repositories.SpaceRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class SpaceService {

    private final SpaceRepository spaceRepository;
    private final SpaceUserRepository spaceUserRepository;
    private final PageService pageService;

    public SpaceService(SpaceRepository spaceRepository,
                        SpaceUserRepository spaceUserRepository, PageService pageService) {
        this.spaceRepository = spaceRepository;
        this.spaceUserRepository = spaceUserRepository;
        this.pageService = pageService;
    }

    public Space createSpace(CreateSpaceDto createSpaceDto, User user) {
        Space space = SpaceMapper.INSTANCE.map(createSpaceDto);
        return createSpace(space, user);
    }

    public Space createSpace(Space space, User user) {
        space = spaceRepository.save(space);
        addUserToSpace(space, user, SpaceRole.OWNER);

        Page defaultPage = new Page();
        defaultPage.setSpace(space);
        defaultPage.setName("Default");
        defaultPage.setVisibility(PageVisibility.PUBLIC);
        pageService.createPage(defaultPage, user);
        return space;
    }

    public List<Space> getSpaces(User user) {
        return spaceRepository.finaAllByUser(user);
    }

    public Optional<Space> getSpaceById(long id) {
        return spaceRepository.findById(id);
    }

    public SpaceUser addUserToSpace(Space space, User user, SpaceRole role) {
        SpaceUser spaceUser = new SpaceUser();
        spaceUser.setSpace(space);
        spaceUser.setUser(user);
        spaceUser.setRole(role);
        return spaceUserRepository.save(spaceUser);
    }

    public void removeUserFromSpace(Space space, User user) {
        Optional<SpaceUser> spaceUser = spaceUserRepository.findBySpaceAndUser(space, user);
        spaceUser.ifPresent(value -> spaceUserRepository.delete(value));
    }
}
