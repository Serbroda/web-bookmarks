package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.dtos.space.AddAccountToSpaceDto;
import de.serbroda.ragbag.dtos.space.CreateSpaceDto;
import de.serbroda.ragbag.mappers.SpaceMapper;
import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.shared.SpaceRole;
import de.serbroda.ragbag.security.AuthorizationService;
import de.serbroda.ragbag.services.SpaceService;
import de.serbroda.ragbag.services.UserService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.nio.file.AccessDeniedException;
import java.util.Optional;

@RestController
@RequestMapping("/api/v1/spaces")
public class SpaceController {

    private final SpaceService spaceService;
    private final UserService userService;

    public SpaceController(SpaceService spaceService, UserService userService) {
        this.spaceService = spaceService;
        this.userService = userService;
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

    @PutMapping("/{spaceId}/accounts/add")
    public ResponseEntity addAccountToSpace(@PathVariable("spaceId") long spaceId, @RequestBody AddAccountToSpaceDto request) throws AccessDeniedException {
        Optional<Space> space = spaceService.getSpaceById(spaceId);
        if (!space.isPresent()) {
            return ResponseEntity.notFound().build();
        }

        AuthorizationService.checkAccessAllowed(space.get(), SpaceRole.OWNER, SpaceRole.CONTRIBUTOR);

        Optional<Account> account = userService.getUserById(request.getAccountId());
        if (!account.isPresent()) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Account not found");
        }
        spaceService.addAccountToSpace(space.get(), account.get(), request.getRole());
        return ResponseEntity.ok(SpaceMapper.INSTANCE.map(space.get()));
    }
}