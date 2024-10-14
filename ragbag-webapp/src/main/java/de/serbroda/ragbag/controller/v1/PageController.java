package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.dtos.page.CreatePageDto;
import de.serbroda.ragbag.dtos.page.PageDto;
import de.serbroda.ragbag.mappers.PageMapper;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.services.PageService;
import de.serbroda.ragbag.security.AuthorizationService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.nio.file.AccessDeniedException;
import java.util.Set;

@RestController
@RequestMapping("/api/v1/pages")
public class PageController {

    private final PageService pageService;

    public PageController(PageService pageService) {
        this.pageService = pageService;
    }

    @PostMapping
    public ResponseEntity<PageDto> createPage(@RequestBody CreatePageDto dto) throws AccessDeniedException {
        Page entity = pageService.createPage(dto, AuthorizationService.getAuthenticatedAccountRequired());
        return ResponseEntity.ok(PageMapper.INSTANCE.map(entity));
    }


}
