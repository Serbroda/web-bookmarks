package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Account;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface AccountRepository extends JpaRepository<Account, Long> {

    Optional<Account> findByUsernameIgnoreCase(String username);
}
