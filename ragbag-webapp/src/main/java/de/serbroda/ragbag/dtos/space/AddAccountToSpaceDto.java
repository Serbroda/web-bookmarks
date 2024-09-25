package de.serbroda.ragbag.dtos.space;

import com.fasterxml.jackson.annotation.JsonProperty;
import de.serbroda.ragbag.models.shared.SpaceRole;
import jakarta.validation.constraints.NotNull;

public class AddAccountToSpaceDto {

    @NotNull
    @JsonProperty("accountId")
    private long accountId;

    @NotNull
    @JsonProperty("role")
    private SpaceRole role;

    public long getAccountId() {
        return accountId;
    }

    public void setAccountId(long accountId) {
        this.accountId = accountId;
    }

    public SpaceRole getRole() {
        return role;
    }

    public void setRole(SpaceRole role) {
        this.role = role;
    }
}
