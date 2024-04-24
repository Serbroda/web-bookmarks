package de.serbroda.ragbag.configuration;

import de.serbroda.ragbag.services.DataInitializer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AppConfiguration {

    @Autowired
    public void initialize(DataInitializer dataInitializer) {
        dataInitializer.initialize();
    }
}
