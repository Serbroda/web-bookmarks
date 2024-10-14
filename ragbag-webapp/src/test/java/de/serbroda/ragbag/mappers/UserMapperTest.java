package de.serbroda.ragbag.mappers;

import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.models.User;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class UserMapperTest {

    @Test
    public void itShouldMapUser() {
        User from = new User();
        from.setUsername("Max");

        UserDto to = UserMapper.INSTANCE.map(from);
        Assertions.assertNotNull(to);
        Assertions.assertEquals(from.getUsername(), to.getUsername());
    }
}
