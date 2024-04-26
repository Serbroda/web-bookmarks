package de.serbroda.ragbag.controller.v1;

import de.serbroda.ragbag.controller.v1.api.UserApi;
import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.mappers.UserMapper;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.repositories.UserRepository;
import java.util.List;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1")
public class UserController implements UserApi {

    private final UserRepository userRepository;

    public UserController(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @Override
    public ResponseEntity<List<UserDto>> getUsers() {
        List<User> users = userRepository.findAll();
        return ResponseEntity.ok(UserMapper.INSTANCE.mapAll(users));
    }
}
