package de.serbroda.ragbag.models.keys;

import jakarta.persistence.Column;
import jakarta.persistence.Embeddable;

@Embeddable
public class SpaceAccountKey {

    private Long spaceId;
    private Long accountId;

    public SpaceAccountKey() {
    }

    public SpaceAccountKey(Long spaceId, Long accountId) {
        this.spaceId = spaceId;
        this.accountId = accountId;
    }

    @Column(name = "space_id")
    public Long getSpaceId() {
        return spaceId;
    }

    public void setSpaceId(Long spaceId) {
        this.spaceId = spaceId;
    }

    @Column(name = "account_id")
    public Long getAccountId() {
        return accountId;
    }

    public void setAccountId(Long accountId) {
        this.accountId = accountId;
    }
}
