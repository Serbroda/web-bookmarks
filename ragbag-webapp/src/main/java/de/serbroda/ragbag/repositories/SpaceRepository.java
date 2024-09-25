package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Space;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface SpaceRepository extends JpaRepository<Space, Long> {

    @Query("select s from Space s left join s.accounts accs where accs.account in ?1")
    List<Space> findAllByAccount(Account account);
}
