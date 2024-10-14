package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.PageUser;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface PageAccountRepository extends JpaRepository<PageUser, Long> {

}
