package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Long> {
}
