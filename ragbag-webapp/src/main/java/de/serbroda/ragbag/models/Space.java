package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import jakarta.persistence.*;

import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "space")
public class Space extends AbstractBaseEntity {

    private String name;
    private Set<Page> pages = new HashSet<>();
    private Set<SpaceAccount> accounts = new HashSet<>();

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Override
    public Long getId() {
        return doGetId();
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    @OneToMany(fetch = FetchType.EAGER, mappedBy = "space")
    public Set<Page> getPages() {
        return pages;
    }

    public void setPages(Set<Page> pages) {
        this.pages = pages;
    }

    @OneToMany(mappedBy = "space")
    public Set<SpaceAccount> getAccounts() {
        return accounts;
    }

    public void setAccounts(Set<SpaceAccount> accounts) {
        this.accounts = accounts;
    }
}
