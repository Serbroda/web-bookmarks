package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import de.serbroda.ragbag.models.shared.PageVisibility;
import jakarta.persistence.Entity;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

@Entity
@Table(name = "page")
public class Page extends AbstractBaseEntity {

    private String name;
    private Space space;
    private Page parent;
    private PageVisibility visibility;
    private Set<Page> subPages = new HashSet<>();
    private Set<Link> links = new HashSet<>();
    private Set<PageAccount> accounts = new HashSet<>();

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

    @ManyToOne(fetch = FetchType.EAGER)
    public Space getSpace() {
        return space;
    }

    public void setSpace(Space space) {
        this.space = space;
    }

    @ManyToOne(fetch = FetchType.EAGER)
    public Page getParent() {
        return parent;
    }

    public void setParent(Page parent) {
        this.parent = parent;
    }

    @Enumerated(EnumType.STRING)
    public PageVisibility getVisibility() {
        return visibility;
    }

    public void setVisibility(PageVisibility visibility) {
        this.visibility = visibility;
    }

    @OneToMany(fetch = FetchType.EAGER, mappedBy = "parent")
    public Set<Page> getSubPages() {
        return subPages;
    }

    public void setSubPages(Set<Page> subPages) {
        this.subPages = subPages;
    }

    @OneToMany(fetch = FetchType.EAGER, mappedBy = "page")
    public Set<Link> getLinks() {
        return links;
    }

    public void setLinks(Set<Link> links) {
        this.links = links;
    }

    @OneToMany(mappedBy = "page")
    public Set<PageAccount> getAccounts() {
        return accounts;
    }

    public void setAccounts(Set<PageAccount> accounts) {
        this.accounts = accounts;
    }
}
