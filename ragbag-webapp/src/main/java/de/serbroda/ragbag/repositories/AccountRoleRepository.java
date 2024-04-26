package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.AccountRole;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AccountRoleRepository extends JpaRepository<AccountRole, Long> {

    AccountRole findByNameIgnoreCase(String name);
}
