package de.serbroda.ragbag.dtos.page;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

public class CreatePageDto {

    @NotNull(message = "spaceId must be set")
    @JsonProperty("spaceId")
    private long spaceId;

    @JsonProperty("parentPageId")
    private Long parentPageId;

    @NotBlank(message = "name must be set")
    @JsonProperty("name")
    private String name;

    @NotNull(message = "spaceId must be set")
    public long getSpaceId() {
        return spaceId;
    }

    public void setSpaceId(long spaceId) {
        this.spaceId = spaceId;
    }

    public Long getParentPageId() {
        return parentPageId;
    }

    public void setParentPageId(Long parentPageId) {
        this.parentPageId = parentPageId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
