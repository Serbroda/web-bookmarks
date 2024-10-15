package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface SpaceRepository extends JpaRepository<Space, Long> {

    @Query("select s from Space s left join s.users u where u.user in ?1")
    List<Space> finaAllByUser(User user);
}
