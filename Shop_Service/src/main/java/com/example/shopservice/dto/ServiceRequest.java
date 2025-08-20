package com.example.shopservice.dto;

import jakarta.validation.constraints.*;

public record ServiceRequest(

    @NotBlank(message = "Service name is required")
    @Size(min = 3, max = 100, message = "Service name must be between 3 and 100 characters")
    String serviceName,

    @Size(max = 500, message = "Description must not exceed 500 characters")
    String description,

    @NotNull(message = "Price is required")
    @Positive(message = "Price must be greater than 0")
    Double price,

    @NotNull(message = "Duration is required")
    @Min(value = 10, message = "Duration must be at least 10 minutes")
    @Max(value = 480, message = "Duration cannot exceed 480 minutes")
    Integer duration
) {}
