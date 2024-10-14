package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import de.serbroda.ragbag.models.shared.SpaceVisibility;
import jakarta.persistence.*;

import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "spaces")
public class Space extends AbstractBaseEntity {

    private String name;
    private SpaceVisibility visibility = SpaceVisibility.PRIVATE;
    private Set<Page> pages = new HashSet<>();
    private Set<SpaceUser> accounts = new HashSet<>();

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

    @Enumerated(EnumType.STRING)
    public SpaceVisibility getVisibility() {
        return visibility;
    }

    public void setVisibility(SpaceVisibility visibility) {
        this.visibility = visibility;
    }

    @OneToMany(fetch = FetchType.EAGER, mappedBy = "space")
    public Set<Page> getPages() {
        return pages;
    }

    public void setPages(Set<Page> pages) {
        this.pages = pages;
    }

    @OneToMany(mappedBy = "space")
    public Set<SpaceUser> getAccounts() {
        return accounts;
    }

    public void setAccounts(Set<SpaceUser> accounts) {
        this.accounts = accounts;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj instanceof Space) {
            return super.equals(obj);
        }
        return false;
    }
}
