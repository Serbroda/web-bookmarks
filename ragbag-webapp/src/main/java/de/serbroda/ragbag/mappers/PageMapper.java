package de.serbroda.ragbag.mappers;

import de.serbroda.ragbag.dtos.page.CreatePageDto;
import de.serbroda.ragbag.dtos.page.PageDto;
import de.serbroda.ragbag.models.Page;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

import java.util.List;

@Mapper
public interface PageMapper {

    PageMapper INSTANCE = Mappers.getMapper(PageMapper.class);

    Page map(CreatePageDto source);

    PageDto map(Page source);

    List<PageDto> mapAll(List<Page> sources);

}
