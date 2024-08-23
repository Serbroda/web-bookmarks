package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import jakarta.persistence.*;

@Entity
@Table(name = "accountrole")
public class AccountRole extends AbstractBaseEntity {

    private String name;

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Override
    public Long getId() {
        return doGetId();
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

}
