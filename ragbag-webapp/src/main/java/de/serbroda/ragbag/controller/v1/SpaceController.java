package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.mappers.SpaceMapper;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.services.SpaceService;
import de.serbroda.ragbag.utils.SessionUtils;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/spaces")
public class SpaceController {

    private final SpaceService spaceService;

    public SpaceController(SpaceService spaceService) {
        this.spaceService = spaceService;
    }

    @PostMapping
    public ResponseEntity createSpace(@RequestBody CreateSpaceDto dto) {
        Space space = spaceService.createSpace(dto, SessionUtils.getAuthenticatedAccountRequired());
        return ResponseEntity.ok(SpaceMapper.INSTANCE.map(space));
    }
}
