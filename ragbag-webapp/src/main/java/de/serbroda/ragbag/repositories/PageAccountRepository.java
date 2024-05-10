package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.PageAccount;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface PageAccountRepository extends JpaRepository<PageAccount, Long> {

}
