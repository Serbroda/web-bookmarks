package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Group;
import de.serbroda.ragbag.models.Space;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface GroupRepository extends JpaRepository<Group, Long> {

}
