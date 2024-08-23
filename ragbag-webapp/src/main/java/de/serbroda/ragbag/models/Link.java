package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import jakarta.persistence.*;

@Entity
@Table(name = "link")
public class Link extends AbstractBaseEntity {

    private String name;
    private String url;
    private Page page;

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

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    @ManyToOne(fetch = FetchType.EAGER)
    public Page getPage() {
        return page;
    }

    public void setPage(Page page) {
        this.page = page;
    }
}
