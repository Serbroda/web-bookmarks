package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import java.util.ArrayList;
import java.util.List;

@Entity
@Table(name = "page")
public class Page extends AbstractBaseEntity {

    private String name;
    private Space space;
    private Page parent;
    private List<Page> subPages = new ArrayList<>();
    private List<Link> links = new ArrayList<>();

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

    @OneToMany(fetch = FetchType.EAGER, mappedBy = "parent")
    public List<Page> getSubPages() {
        return subPages;
    }

    public void setSubPages(List<Page> subPages) {
        this.subPages = subPages;
    }

    @OneToMany(fetch = FetchType.EAGER, mappedBy = "page")
    public List<Link> getLinks() {
        return links;
    }

    public void setLinks(List<Link> links) {
        this.links = links;
    }
}
