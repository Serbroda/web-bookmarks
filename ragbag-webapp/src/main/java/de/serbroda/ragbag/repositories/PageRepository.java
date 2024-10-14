package de.serbroda.ragbag.repositories;

import de.serbroda.ragbag.models.Page;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Set;

@Repository
public interface PageRepository extends JpaRepository<Page, Long> {

    @Query("SELECT p FROM Page p WHERE p.space.id = :spaceId AND p.parent IS NULL")
    List<Page> findRootPagesBySpaceId(@Param("spaceId") Long spaceId);
}
