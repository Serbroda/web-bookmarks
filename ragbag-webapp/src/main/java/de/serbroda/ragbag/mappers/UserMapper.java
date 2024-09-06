package de.serbroda.ragbag.mappers;

import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.models.Account;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

import java.util.List;

@Mapper
public interface UserMapper {

    UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);

    UserDto map(Account account);

    List<UserDto> mapAll(List<Account> accounts);
}