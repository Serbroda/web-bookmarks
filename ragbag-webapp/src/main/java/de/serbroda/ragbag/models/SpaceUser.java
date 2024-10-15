package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.keys.SpaceUserKey;
import de.serbroda.ragbag.models.shared.SpaceRole;
import jakarta.persistence.*;

@Entity
@Table(name = "spaces_users")
public class SpaceUser {

    private SpaceUserKey id = new SpaceUserKey();
    private Space space;
    private User user;
    private SpaceRole role;

    @EmbeddedId
    public SpaceUserKey getId() {
        return id;
    }

    public void setId(SpaceUserKey id) {
        this.id = id;
    }

    @ManyToOne(fetch = FetchType.LAZY)
    @MapsId("spaceId")
    @JoinColumn(name = "space_id")
    public Space getSpace() {
        return space;
    }

    public void setSpace(Space space) {
        this.space = space;
    }

    @ManyToOne(fetch = FetchType.EAGER)
    @MapsId("userId")
    @JoinColumn(name = "user_id")
    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    @Enumerated(EnumType.STRING)
    public SpaceRole getRole() {
        return role;
    }

    public void setRole(SpaceRole role) {
        this.role = role;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj instanceof SpaceUser) {
            return super.equals(obj);
        }
        return false;
    }
}
