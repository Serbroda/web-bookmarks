package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.AccountPermission;
import de.serbroda.ragbag.models.AccountRole;
import java.util.List;
import java.util.Set;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AccountPermissionRepository extends JpaRepository<AccountPermission, Long> {

    AccountPermission findByNameIgnoreCase(String name);

    Set<AccountPermission> findByNameIn(List<String> names);
}
