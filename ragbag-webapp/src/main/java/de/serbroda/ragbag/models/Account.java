package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import jakarta.persistence.*;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;

import java.util.*;

@Entity
@Table(name = "account")
public class Account extends AbstractBaseEntity {

    private String username;
    private String password;
    private boolean active = true;
    private Set<AccountRole> accountRoles = new HashSet<>();
    private Set<SpaceAccount> spaces = new HashSet<>();
    private Set<PageAccount> pages = new HashSet<>();

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

//    @JoinTable(name = "space_account", joinColumns = {
//        @JoinColumn(name = "account_id", referencedColumnName = "id")}, inverseJoinColumns = {
//        @JoinColumn(name = "space_id", referencedColumnName = "id")
//    })
//    @ManyToMany(fetch = FetchType.EAGER)
//    public Set<SpaceAccount> getSpaces() {
//        return spaces;
//    }

    @OneToMany(mappedBy = "account")
    public Set<SpaceAccount> getSpaces() {
        return spaces;
    }

    public void setSpaces(Set<SpaceAccount> spaces) {
        this.spaces = spaces;
    }

    @OneToMany(mappedBy = "account")
    public Set<PageAccount> getPages() {
        return pages;
    }

    public void setPages(Set<PageAccount> pages) {
        this.pages = pages;
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
