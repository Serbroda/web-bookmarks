package de.serbroda.ragbag.controller.v1.api;

import de.serbroda.ragbag.dtos.UserDto;
import de.serbroda.ragbag.dtos.auth.RegisterUserDto;
import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.info.Info;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.List;

@Validated
@OpenAPIDefinition(info = @Info(
        title = "Users API",
        version = "0.0",
        description = "My API")
)
public interface UserApi {

    @Operation(
            summary = "Get a list of users",
            tags = {"users"}
    )
    @ApiResponses({@ApiResponse(
            responseCode = "200",
            description = "Successful operation",
            content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = UserDto.class)
            ))
    })
    @GetMapping(
            value = {"/users"},
            produces = {"application/json"}
    )
    default ResponseEntity<List<UserDto>> getUsers() {
        return ResponseEntity.status(HttpStatus.NOT_IMPLEMENTED).build();
    }

    @Operation(
            summary = "Creating a user",
            tags = {"users"}
    )
    @ApiResponses({@ApiResponse(
            responseCode = "200",
            description = "Successful operation",
            content = @Content(
                    mediaType = "application/json",
                    schema = @Schema(implementation = UserDto.class)
            )),
            @ApiResponse(
                    responseCode = "400",
                    description = "Bad Request"
            )
    })
    @PostMapping(
            value = {"/users"},
            produces = {"application/json"},
            consumes = {"application/json"}
    )
    default ResponseEntity<UserDto> createUser(
            @Parameter(required = true) @RequestBody @Valid RegisterUserDto createUserDto) {
        return ResponseEntity.status(HttpStatus.NOT_IMPLEMENTED).build();
    }
}
