package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.SpaceUser;
import de.serbroda.ragbag.models.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface SpaceUserRepository extends JpaRepository<SpaceUser, Long> {

    Optional<SpaceUser> findBySpaceAndUser(Space space, User user);
}
