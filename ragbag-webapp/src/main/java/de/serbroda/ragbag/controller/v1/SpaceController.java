package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.mappers.PageMapper;
import de.serbroda.ragbag.mappers.SpaceMapper;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.services.AuthorizationService;
import de.serbroda.ragbag.services.PageService;
import de.serbroda.ragbag.services.SpaceService;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

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
        Space space = spaceService.createSpace(dto, AuthorizationService.getAuthenticatedUserRequired());
        return ResponseEntity.ok(SpaceMapper.INSTANCE.map(space));
    }

    @GetMapping
    public ResponseEntity getSpaces() {
        return ResponseEntity.ok(SpaceMapper.INSTANCE.mapAll(spaceService.getSpaces(AuthorizationService.getAuthenticatedUserRequired())));
    }

    @PreAuthorize("@authorizationService.hasAccessToSpace(authentication, #spaceId)")
    @GetMapping("/{spaceId}")
    public ResponseEntity getSpaceById(@PathVariable("spaceId") long spaceId) {
        return spaceService.getSpaceById(spaceId)
                .map(SpaceMapper.INSTANCE::map)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    @PreAuthorize("@authorizationService.hasAccessToSpace(authentication, #spaceId)")
    @GetMapping("/{spaceId}/pages")
    public ResponseEntity getSpacePages(@PathVariable("spaceId") long spaceId) {
        return spaceService.getSpaceById(spaceId)
                .map(s -> pageService.getPagesTreeBySpaceId(s.getId()))
                .map(PageMapper.INSTANCE::mapAll)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
}