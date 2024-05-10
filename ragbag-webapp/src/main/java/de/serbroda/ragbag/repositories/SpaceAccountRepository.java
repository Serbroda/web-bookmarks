package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.PageAccount;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceAccount;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SpaceAccountRepository extends JpaRepository<SpaceAccount, Long> {

    Optional<SpaceAccount> findBySpaceAndAccount(Space space, Account account);
}
