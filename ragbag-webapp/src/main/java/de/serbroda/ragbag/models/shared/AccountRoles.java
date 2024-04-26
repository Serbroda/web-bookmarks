package de.serbroda.ragbag.models.shared;

import static de.serbroda.ragbag.models.shared.AccountPermissions.DELETE_SPACES;
import static de.serbroda.ragbag.models.shared.AccountPermissions.EDIT_SPACES;
import static de.serbroda.ragbag.models.shared.AccountPermissions.VIEW_SPACES;

import java.util.Set;

public enum AccountRoles {
    ADMIN(VIEW_SPACES, EDIT_SPACES, DELETE_SPACES),
    USER(VIEW_SPACES);

    private final Set<AccountPermissions> permissions;

    AccountRoles(AccountPermissions... permissions) {
        this.permissions = Set.of(permissions);
    }

    public Set<AccountPermissions> getPermissions() {
        return permissions;
    }
}
