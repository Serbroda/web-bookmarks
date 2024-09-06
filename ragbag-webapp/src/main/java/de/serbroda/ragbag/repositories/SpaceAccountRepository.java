package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceAccount;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface SpaceAccountRepository extends JpaRepository<SpaceAccount, Long> {

    Optional<SpaceAccount> findBySpaceAndAccount(Space space, Account account);
}