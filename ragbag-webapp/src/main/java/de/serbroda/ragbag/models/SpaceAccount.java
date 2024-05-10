package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import de.serbroda.ragbag.models.keys.SpaceAccountKey;
import de.serbroda.ragbag.models.shared.SpaceRole;
import jakarta.persistence.Column;
import jakarta.persistence.EmbeddedId;
import jakarta.persistence.Entity;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.MapsId;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import java.util.ArrayList;
import java.util.List;

@Entity
@Table(name = "space_account")
public class SpaceAccount {

    private SpaceAccountKey id = new SpaceAccountKey();
    private Space space;
    private Account account;
    private SpaceRole role;

    @EmbeddedId
    public SpaceAccountKey getId() {
        return id;
    }

    public void setId(SpaceAccountKey id) {
        this.id = id;
    }

    @ManyToOne
    @MapsId("spaceId")
    @JoinColumn(name = "space_id")
    public Space getSpace() {
        return space;
    }

    public void setSpace(Space space) {
        this.space = space;
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
    public SpaceRole getRole() {
        return role;
    }

    public void setRole(SpaceRole role) {
        this.role = role;
    }
}
