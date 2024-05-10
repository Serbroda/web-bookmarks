package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceAccount;
import de.serbroda.ragbag.models.shared.SpaceRole;
import de.serbroda.ragbag.repositories.SpaceAccountRepository;
import de.serbroda.ragbag.repositories.SpaceRepository;
import java.util.Optional;
import javax.management.relation.Role;
import org.springframework.stereotype.Service;

@Service
public class SpaceService {

    private SpaceRepository spaceRepository;
    private SpaceAccountRepository spaceAccountRepository;

    public SpaceService(SpaceRepository spaceRepository,
        SpaceAccountRepository spaceAccountRepository) {
        this.spaceRepository = spaceRepository;
        this.spaceAccountRepository = spaceAccountRepository;
    }

    public Space createSpace(Space space, Account account) {
        space = spaceRepository.save(space);
        addAccountToSpace(space, account, SpaceRole.OWNER);
        return space;
    }

    public SpaceAccount addAccountToSpace(Space space, Account account, SpaceRole role) {
        SpaceAccount spaceAccount = new SpaceAccount();
        spaceAccount.setSpace(space);
        spaceAccount.setAccount(account);
        spaceAccount.setRole(role);
        return spaceAccountRepository.save(spaceAccount);
    }

    private void removeAccountFromSpace(Space space, Account account) {
        Optional<SpaceAccount> spaceAccount = spaceAccountRepository.findBySpaceAndAccount(space, account);
        spaceAccount.ifPresent(value -> spaceAccountRepository.delete(value));
    }
}
