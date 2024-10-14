package de.serbroda.ragbag.models.keys;

import jakarta.persistence.Column;
import jakarta.persistence.Embeddable;

@Embeddable
public class SpaceUserKey {

    private Long spaceId;
    private Long userId;

    public SpaceUserKey() {
    }

    public SpaceUserKey(Long spaceId, Long userId) {
        this.spaceId = spaceId;
        this.userId = userId;
    }

    @Column(name = "space_id")
    public Long getSpaceId() {
        return spaceId;
    }

    public void setSpaceId(Long spaceId) {
        this.spaceId = spaceId;
    }

    @Column(name = "account_id")
    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long accountId) {
        this.userId = accountId;
    }
}
