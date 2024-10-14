package de.serbroda.ragbag.services;

import de.serbroda.ragbag.dtos.auth.LoginUserDto;
import de.serbroda.ragbag.dtos.auth.RegisterUserDto;
import de.serbroda.ragbag.models.User;
import de.serbroda.ragbag.repositories.AccountRepository;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class AuthenticationService {

    private final PasswordEncoder passwordEncoder;
    private final AccountRepository accountRepository;
    private final AuthenticationManager authenticationManager;

    public AuthenticationService(PasswordEncoder passwordEncoder,
                                 AccountRepository accountRepository,
                                 AuthenticationManager authenticationManager) {
        this.passwordEncoder = passwordEncoder;
        this.accountRepository = accountRepository;
        this.authenticationManager = authenticationManager;
    }

    public User signup(RegisterUserDto input) {
        User user = new User();
        user.setUsername(input.getEmail());
        user.setPassword(passwordEncoder.encode(input.getPassword()));
        return accountRepository.save(user);
    }

    public User authenticate(LoginUserDto input) {
        authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(
                        input.getEmail(),
                        input.getPassword()
                )
        );

        return accountRepository.findByUsernameIgnoreCase(input.getEmail())
                .orElseThrow();
    }
}
