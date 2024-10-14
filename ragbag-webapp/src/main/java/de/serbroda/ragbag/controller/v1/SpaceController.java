package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.mappers.PageMapper;
import de.serbroda.ragbag.mappers.SpaceMapper;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.security.AuthorizationService;
import de.serbroda.ragbag.services.PageService;
import de.serbroda.ragbag.services.SpaceService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.nio.file.AccessDeniedException;
import java.util.Optional;

@RestController
@RequestMapping("/api/v1/spaces")
public class SpaceController {

    private final SpaceService spaceService;
    private final PageService pageService;

    public SpaceController(SpaceService spaceService, PageService pageService) {
        this.spaceService = spaceService;
        this.pageService = pageService;
    }

    @PostMapping
    public ResponseEntity createSpace(@RequestBody CreateSpaceDto dto) {
        Space space = spaceService.createSpace(dto, AuthorizationService.getAuthenticatedAccountRequired());
        return ResponseEntity.ok(SpaceMapper.INSTANCE.map(space));
    }

    @GetMapping
    public ResponseEntity getSpaces() {
        return ResponseEntity.ok(SpaceMapper.INSTANCE.mapAll(spaceService.getSpaces(AuthorizationService.getAuthenticatedAccountRequired())));
    }

    @GetMapping("/{spaceId}")
    public ResponseEntity getSpaceById(@PathVariable("spaceId") long spaceId) throws AccessDeniedException {
        Optional<Space> space = spaceService.getSpaceById(spaceId);
        if (!space.isPresent()) {
            return ResponseEntity.notFound().build();
        }
        AuthorizationService.checkAccessAllowed(space.get());
        return ResponseEntity.ok(SpaceMapper.INSTANCE.map(space.get()));
    }

    @GetMapping("/{spaceId}/pages")
    public ResponseEntity getSpacePages(@PathVariable("spaceId") long spaceId) throws AccessDeniedException {
        Optional<Space> space = spaceService.getSpaceById(spaceId);
        if (!space.isPresent()) {
            return ResponseEntity.notFound().build();
        }
        AuthorizationService.checkAccessAllowed(space.get());

        return ResponseEntity.ok(PageMapper.INSTANCE.mapAll(pageService.getPagesTreeBySpaceId(spaceId)));
    }
}