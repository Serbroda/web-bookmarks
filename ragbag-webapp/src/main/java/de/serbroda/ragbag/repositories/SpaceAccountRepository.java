package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceUser;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface SpaceAccountRepository extends JpaRepository<SpaceUser, Long> {

    Optional<SpaceUser> findBySpaceAndAccount(Space space, User user);
}
