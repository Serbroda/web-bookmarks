package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.mappers.UserMapper;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.repositories.UserRepository;
import de.serbroda.ragbag.services.AuthorizationService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1/users")
public class UserController {

    private final UserRepository userRepository;

    public UserController(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @PreAuthorize("hasAnyRole('ADMIN')")
    @GetMapping
    public ResponseEntity<List<UserDto>> getUsers() {
        List<User> users = userRepository.findAll();
        return ResponseEntity.ok(UserMapper.INSTANCE.mapAll(users));
    }

    @GetMapping("/me")
    public ResponseEntity<UserDto> authenticatedUser() {
        return AuthorizationService.getAuthenticatedUser()
                .map(UserMapper.INSTANCE::map)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.status(HttpStatus.UNAUTHORIZED).build());
    }

}
