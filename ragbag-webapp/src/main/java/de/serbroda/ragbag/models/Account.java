package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.JoinTable;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.Table;
import jakarta.persistence.Transient;
import java.util.ArrayList;
import java.util.Collection;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;

@Entity
@Table(name = "account")
public class Account extends AbstractBaseEntity {

    private String username;
    private String password;
    private boolean active = true;
    private Set<AccountRole> accountRoles = new HashSet<>();
    private Set<Space> spaces = new HashSet<>();

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Override
    public Long getId() {
        return doGetId();
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public boolean isActive() {
        return active;
    }

    public void setActive(boolean active) {
        this.active = active;
    }

    @JoinTable(name = "account_accountrole", joinColumns = {
        @JoinColumn(name = "account_id", referencedColumnName = "id")}, inverseJoinColumns = {
        @JoinColumn(name = "role_id", referencedColumnName = "id")
    })
    @ManyToMany(fetch = FetchType.EAGER)
    public Set<AccountRole> getAccountRoles() {
        return accountRoles;
    }

    public void setAccountRoles(Set<AccountRole> accountRoles) {
        this.accountRoles = accountRoles;
    }

    @JoinTable(name = "space_account", joinColumns = {
        @JoinColumn(name = "account_id", referencedColumnName = "id")}, inverseJoinColumns = {
        @JoinColumn(name = "space_id", referencedColumnName = "id")
    })
    @ManyToMany(fetch = FetchType.EAGER)
    public Set<Space> getSpaces() {
        return spaces;
    }

    public void setSpaces(Set<Space> spaces) {
        this.spaces = spaces;
    }

    @Transient
    public Collection<? extends GrantedAuthority> getAuthorities() {
        List<GrantedAuthority> authorities = new ArrayList<>();
        for (AccountRole r : accountRoles) {
            authorities.add(new SimpleGrantedAuthority("ROLE_" + r.getName()));
        }
        return authorities;
    }
}
