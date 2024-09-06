package de.serbroda.ragbag.controller;

import io.jsonwebtoken.ExpiredJwtException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ProblemDetail;
import org.springframework.security.authentication.AccountStatusException;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.nio.file.AccessDeniedException;
import java.security.SignatureException;

@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler({UsernameNotFoundException.class, BadCredentialsException.class})
    public ProblemDetail handleBadCredentialsException(Exception exception) {
        return createProblemDetails(
                exception,
                HttpStatus.UNAUTHORIZED,
                "The username or password is incorrect"
        );
    }

    @ExceptionHandler(AccountStatusException.class)
    public ProblemDetail handleAccountStatusException(Exception exception) {
        return createProblemDetails(
                exception,
                HttpStatus.FORBIDDEN,
                "The account is locked"
        );
    }

    @ExceptionHandler(AccessDeniedException.class)
    public ProblemDetail handleAccessDeniedException(Exception exception) {
        return createProblemDetails(
                exception,
                HttpStatus.FORBIDDEN,
                "You are not authorized to access this resource"
        );
    }

    @ExceptionHandler(SignatureException.class)
    public ProblemDetail handleSignatureException(Exception exception) {
        return createProblemDetails(
                exception,
                HttpStatus.FORBIDDEN,
                "The JWT signature is invalid"
        );
    }

    @ExceptionHandler(ExpiredJwtException.class)
    public ProblemDetail handleExpiredJwtException(Exception exception) {
        return createProblemDetails(
                exception,
                HttpStatus.FORBIDDEN,
                "The JWT token has expired"
        );
    }

    @ExceptionHandler(Exception.class)
    public ProblemDetail handleException(Exception exception) {
        return createProblemDetails(
                exception,
                HttpStatus.INTERNAL_SERVER_ERROR,
                "Unknown internal server error"
        );
    }

    private ProblemDetail createProblemDetails(Exception exception, HttpStatus status, String description) {
        exception.printStackTrace();

        ProblemDetail errorDetail = ProblemDetail.forStatusAndDetail(status, exception.getMessage());
        errorDetail.setProperty("description", description);
        return errorDetail;
    }
}
