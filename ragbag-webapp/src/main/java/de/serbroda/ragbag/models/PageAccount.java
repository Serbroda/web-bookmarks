package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.keys.PageAccountKey;
import de.serbroda.ragbag.models.shared.PageRole;
import jakarta.persistence.EmbeddedId;
import jakarta.persistence.Entity;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.MapsId;
import jakarta.persistence.Table;

@Entity
@Table(name = "page_account")
public class PageAccount {

    private PageAccountKey id;
    private Page page;
    private Account account;
    private PageRole role;

    @EmbeddedId
    public PageAccountKey getId() {
        return id;
    }

    public void setId(PageAccountKey id) {
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
    @MapsId("accountId")
    @JoinColumn(name = "account_id")
    public Account getAccount() {
        return account;
    }

    public void setAccount(Account account) {
        this.account = account;
    }

    @Enumerated(EnumType.STRING)
    public PageRole getRole() {
        return role;
    }

    public void setRole(PageRole role) {
        this.role = role;
    }
}
