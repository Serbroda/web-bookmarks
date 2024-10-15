package de.serbroda.ragbag.models;

import de.serbroda.ragbag.models.base.AbstractBaseEntity;
import jakarta.persistence.*;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.*;

@Entity
@Table(name = "users")
public class User extends AbstractBaseEntity implements UserDetails {

    private String username;
    private String password;
    private boolean active = true;
    private Set<UserRole> userRoles = new HashSet<>();
    private Set<SpaceUser> spaces = new HashSet<>();
    private Set<PageUser> pages = new HashSet<>();

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Override
    public Long getId() {
        return doGetId();
    }

    @Override
    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    @Override
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

    @JoinTable(name = "users_role", joinColumns = {
            @JoinColumn(name = "user_id", referencedColumnName = "id")}, inverseJoinColumns = {
            @JoinColumn(name = "role_id", referencedColumnName = "id")
    })
    @ManyToMany(fetch = FetchType.EAGER)
    public Set<UserRole> getUserRoles() {
        return userRoles;
    }

    public void setUserRoles(Set<UserRole> userRoles) {
        this.userRoles = userRoles;
    }

//    @JoinTable(name = "space_account", joinColumns = {
//        @JoinColumn(name = "account_id", referencedColumnName = "id")}, inverseJoinColumns = {
//        @JoinColumn(name = "space_id", referencedColumnName = "id")
//    })
//    @ManyToMany(fetch = FetchType.EAGER)
//    public Set<SpaceAccount> getSpaces() {
//        return spaces;
//    }

    @OneToMany(mappedBy = "user", fetch = FetchType.EAGER)
    public Set<SpaceUser> getSpaces() {
        return spaces;
    }

    public void setSpaces(Set<SpaceUser> spaces) {
        this.spaces = spaces;
    }

    @OneToMany(mappedBy = "user")
    public Set<PageUser> getPages() {
        return pages;
    }

    public void setPages(Set<PageUser> pages) {
        this.pages = pages;
    }

    @Transient
    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        List<GrantedAuthority> authorities = new ArrayList<>();
        for (UserRole r : userRoles) {
            authorities.add(new SimpleGrantedAuthority("ROLE_" + r.getName()));
        }
        return authorities;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj instanceof User) {
            return super.equals(obj);
        }
        return false;
    }
}
