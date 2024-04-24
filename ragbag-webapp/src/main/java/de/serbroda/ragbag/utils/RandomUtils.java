package de.serbroda.ragbag.utils;

import java.util.Random;

public class RandomUtils {

    private static final String ALPHANUMERIC_CHARACTERS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

    public static String randomString(int length) {
        return randomString(length, ALPHANUMERIC_CHARACTERS);
    }

    public static String randomString(int length, String characters) {
        StringBuilder sb = new StringBuilder(length);
        Random rnd = new Random();
        for (int i = 0; i < length; i++) {
            sb.append(characters.charAt(rnd.nextInt(characters.length())));
        }
        return sb.toString();
    }

    public static int randomInt(int min, int max) {
        return new Random().nextInt(max + 1 - min) + min;
    }
}
