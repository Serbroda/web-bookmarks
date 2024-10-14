package de.serbroda.ragbag.configuration;

import de.serbroda.ragbag.repositories.AccountRepository;
import de.serbroda.ragbag.services.DataInitializer;
import de.serbroda.ragbag.services.DataInitializerTest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Profile;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.authentication.dao.DaoAuthenticationProvider;
import org.springframework.security.config.annotation.authentication.configuration.AuthenticationConfiguration;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

@Configuration
public class AppConfiguration {

    private final AccountRepository accountRepository;

    public AppConfiguration(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    @Profile({"prod"})
    @Configuration
    public static class DefaultAppConfiguration {

        @Autowired
        public void initialize(DataInitializer dataInitializer) {
            dataInitializer.initialize();
        }
    }

    @Profile("!prod")
    @Configuration
    public static class TestAppConfiguration {

        @Autowired
        public void initialize(DataInitializerTest dataInitializer) {
            dataInitializer.initialize();
        }
    }

    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }

    @Bean
    public UserDetailsService userDetailsService() {
        return username -> accountRepository.findByUsernameIgnoreCase(username)
                .orElseThrow(() -> new UsernameNotFoundException("User not found"));
    }

    @Bean
    public AuthenticationManager authenticationManager(AuthenticationConfiguration config) throws Exception {
        return config.getAuthenticationManager();
    }

    @Bean
    AuthenticationProvider authenticationProvider() {
        DaoAuthenticationProvider authProvider = new DaoAuthenticationProvider();

        authProvider.setUserDetailsService(userDetailsService());
        authProvider.setPasswordEncoder(passwordEncoder());

        return authProvider;
    }
}
