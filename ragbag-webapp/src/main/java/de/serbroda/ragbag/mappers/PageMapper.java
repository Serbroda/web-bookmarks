package de.serbroda.ragbag.mappers;

import de.serbroda.ragbag.dtos.page.CreatePageDto;
import de.serbroda.ragbag.dtos.page.PageDto;
import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.dtos.space.SpaceDto;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.Space;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

@Mapper
public interface PageMapper {

    PageMapper INSTANCE = Mappers.getMapper(PageMapper.class);

    Page map(CreatePageDto source);

    PageDto map(Page source);
}
