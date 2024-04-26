package de.serbroda.ragbag.mappers;

import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.models.Account;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class AccountMapperTest {

    @Test
    public void itShouldMapUser() {
        Account from = new Account();
        from.setUsername("Max");

        UserDto to = UserMapper.INSTANCE.map(from);
        Assertions.assertNotNull(to);
        Assertions.assertEquals(from.getUsername(), to.getUsername());
    }
}
