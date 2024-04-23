package de.serbroda.ragbag.models.base;

import java.io.Serializable;
import java.util.Date;

public interface BaseEntity extends Serializable {

    Long getId();

    void setId(long id);

    int getVersion();

    void setVersion(int version);

    Date getCreated();

    void setCreated(Date created);

    Date getLastModified();

    void setLastModified(Date lastModified);
}
