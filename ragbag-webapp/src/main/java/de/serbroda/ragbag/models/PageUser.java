package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.keys.PageUserKey;
import de.serbroda.ragbag.models.shared.PageRole;
import jakarta.persistence.*;

@Entity
@Table(name = "pages_users")
public class PageUser {

    private PageUserKey id = new PageUserKey();
    private Page page;
    private User user;
    private PageRole role;

    @EmbeddedId
    public PageUserKey getId() {
        return id;
    }

    public void setId(PageUserKey id) {
        this.id = id;
    }

    @ManyToOne
    @MapsId("pageId")
    @JoinColumn(name = "page_id")
    public Page getPage() {
        return page;
    }

    public void setPage(Page page) {
        this.page = page;
    }

    @ManyToOne
    @MapsId("userId")
    @JoinColumn(name = "user_id")
    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    @Enumerated(EnumType.STRING)
    public PageRole getRole() {
        return role;
    }

    public void setRole(PageRole role) {
        this.role = role;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj instanceof PageUser) {
            return super.equals(obj);
        }
        return false;
    }
}
