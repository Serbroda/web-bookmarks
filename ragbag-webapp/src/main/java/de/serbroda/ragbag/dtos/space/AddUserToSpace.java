package de.serbroda.ragbag.dtos.space;

import com.fasterxml.jackson.annotation.JsonProperty;
import de.serbroda.ragbag.models.shared.SpaceRole;
import jakarta.validation.constraints.NotNull;

public class AddUserToSpace {

    @NotNull
    @JsonProperty("userId")
    private long userId;

    @NotNull
    @JsonProperty("role")
    private SpaceRole role;

    public long getUserId() {
        return userId;
    }

    public void setUserId(long userId) {
        this.userId = userId;
    }

    public SpaceRole getRole() {
        return role;
    }

    public void setRole(SpaceRole role) {
        this.role = role;
    }
}
