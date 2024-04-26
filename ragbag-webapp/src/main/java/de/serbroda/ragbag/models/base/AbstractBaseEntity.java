package de.serbroda.ragbag.models.base;

import jakarta.persistence.Column;
import jakarta.persistence.MappedSuperclass;
import jakarta.persistence.PrePersist;
import jakarta.persistence.PreUpdate;
import jakarta.persistence.Temporal;
import jakarta.persistence.TemporalType;
import jakarta.persistence.Version;
import java.util.Date;
import java.util.Objects;

@MappedSuperclass
public abstract class AbstractBaseEntity implements BaseEntity {

    private Long id;
    private int version;
    private Date created;
    private Date lastModified;

    protected Long doGetId() {
        return id;
    }

    @Override
    public void setId(long id) {
        this.id = id;
    }

    @Version
    @Override
    public int getVersion() {
        return version;
    }

    @Override
    public void setVersion(int version) {
        this.version = version;
    }

    @Override
    @Temporal(TemporalType.TIMESTAMP)
    @Column(updatable = false)
    public Date getCreated() {
        return created != null ? new Date(created.getTime()) : null;
    }

    @Override
    public void setCreated(Date created) {
        this.created = created != null ? new Date(created.getTime()) : null;
    }

    @Override
    @Temporal(TemporalType.TIMESTAMP)
    @Column(name = "last_modified")
    public Date getLastModified() {
        return lastModified != null ? new Date(lastModified.getTime()) : null;
    }

    @Override
    public void setLastModified(Date lastModified) {
        this.lastModified = lastModified != null ? new Date(lastModified.getTime()) : null;
    }

    @PrePersist
    protected void setCreatedAndLastModifiedOnCreate() {
        Date now = new Date();
        created = now;
        lastModified = now;
    }

    @PreUpdate
    protected void setLastModifiedOnUpdate() {
        lastModified = new Date();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        AbstractBaseEntity that = (AbstractBaseEntity) o;
        return Objects.equals(id, that.id);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id);
    }
}
