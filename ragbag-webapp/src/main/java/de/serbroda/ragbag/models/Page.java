package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import de.serbroda.ragbag.models.shared.PageVisibility;
import jakarta.persistence.*;

import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "page")
public class Page extends AbstractBaseEntity {

    private String name;
    private Space space;
    private Page parent;
    private PageVisibility visibility = PageVisibility.PUBLIC;
    private Set<Page> subPages = new HashSet<>();
    private Set<Bookmark> bookmarks = new HashSet<>();
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
    public Set<Bookmark> getBookmarks() {
        return bookmarks;
    }

    public void setBookmarks(Set<Bookmark> bookmarks) {
        this.bookmarks = bookmarks;
    }

    @OneToMany(mappedBy = "page")
    public Set<PageAccount> getAccounts() {
        return accounts;
    }

    public void setAccounts(Set<PageAccount> accounts) {
        this.accounts = accounts;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj instanceof Page) {
            return super.equals(obj);
        }
        return false;
    }
}
