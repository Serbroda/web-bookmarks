package de.serbroda.ragbag.models.keys;

import jakarta.persistence.Column;
import jakarta.persistence.Embeddable;

@Embeddable
public class PageAccountKey {

    private Long pageId;
    private Long accountId;

    @Column(name = "page_id")
    public Long getPageId() {
        return pageId;
    }

    public void setPageId(Long pageId) {
        this.pageId = pageId;
    }

    @Column(name = "account_id")
    public Long getAccountId() {
        return accountId;
    }

    public void setAccountId(Long accountId) {
        this.accountId = accountId;
    }
}
