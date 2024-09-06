package de.serbroda.ragbag.mappers;

import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.dtos.space.SpaceDto;
import de.serbroda.ragbag.models.Space;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

@Mapper
public interface SpaceMapper {

    SpaceMapper INSTANCE = Mappers.getMapper(SpaceMapper.class);

    Space map(CreateSpaceDto source);

    SpaceDto map(Space source);
}
