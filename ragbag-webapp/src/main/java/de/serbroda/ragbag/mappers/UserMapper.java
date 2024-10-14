package de.serbroda.ragbag.mappers;

import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.models.User;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

import java.util.List;

@Mapper
public interface UserMapper {

    UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);

    UserDto map(User user);

    List<UserDto> mapAll(List<User> users);
}
