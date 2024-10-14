package de.serbroda.ragbag.services;

import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.mappers.SpaceMapper;
import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceAccount;
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

    public Space createSpace(CreateSpaceDto createSpaceDto, Account account) {
        Space space = SpaceMapper.INSTANCE.map(createSpaceDto);
        return createSpace(space, account);
    }

    public Space createSpace(Space space, Account account) {
        space = spaceRepository.save(space);
        addAccountToSpace(space, account, SpaceRole.OWNER);

        Page defaultPage = new Page();
        defaultPage.setSpace(space);
        defaultPage.setName("Default");
        defaultPage.setVisibility(PageVisibility.PUBLIC);
        pageService.createPage(defaultPage, account);
        return space;
    }

    public List<Space> getSpaces(Account account) {
        return spaceRepository.findAllByAccount(account);
    }

    public Optional<Space> getSpaceById(long id) {
        return spaceRepository.findById(id);
    }

    public SpaceAccount addAccountToSpace(Space space, Account account, SpaceRole role) {
        SpaceAccount spaceAccount = new SpaceAccount();
        spaceAccount.setSpace(space);
        spaceAccount.setAccount(account);
        spaceAccount.setRole(role);
        return spaceAccountRepository.save(spaceAccount);
    }

    public void removeAccountFromSpace(Space space, Account account) {
        Optional<SpaceAccount> spaceAccount = spaceAccountRepository.findBySpaceAndAccount(space, account);
        spaceAccount.ifPresent(value -> spaceAccountRepository.delete(value));
    }
}
