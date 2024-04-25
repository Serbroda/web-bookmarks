package de.serbroda.ragbag.configuration;

import de.serbroda.ragbag.services.DataInitializer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Profile;

@Configuration
public class AppConfiguration {

    @Profile("!test")
    @Configuration
    public static class DefaultAppConfiguration {

        @Autowired
        public void initialize(DataInitializer dataInitializer) {
            dataInitializer.initialize();
        }
    }
}
