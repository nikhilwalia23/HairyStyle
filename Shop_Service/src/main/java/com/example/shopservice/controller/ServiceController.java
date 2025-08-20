package com.example.shopservice.controller;

import com.example.shopservice.dto.ServiceRequest;
import com.example.shopservice.dto.StylishRequest;
import com.example.shopservice.dto.StylishIdRequest;
import com.example.shopservice.model.ServiceEntity;
import com.example.shopservice.model.StylishEntity;
import com.example.shopservice.service.ServiceService;
import com.example.shopservice.service.StylishService;
import jakarta.validation.Valid;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/services")
public class ServiceController {

    private final ServiceService serviceService;
    private final StylishService stylishService;

    public ServiceController(ServiceService serviceService, StylishService stylishService) {
        this.serviceService = serviceService;
        this.stylishService = stylishService;
    }

    @PostMapping("/create")
    public ResponseEntity<ServiceEntity> createService(@Valid @RequestBody ServiceRequest request) {
        ServiceEntity savedService = serviceService.addService(request);
        return ResponseEntity.ok(savedService);
    }

    @PostMapping("/stylish")
    public ResponseEntity<StylishEntity> addOrUpdateStylish(@Valid @RequestBody StylishRequest request) {
        StylishEntity entity = stylishService.addOrUpdateStylish(request);
        return ResponseEntity.ok(entity);
    }


    @PostMapping("/stylish/view")
    public ResponseEntity<?> viewStylishById(@RequestBody StylishIdRequest request) {
        StylishEntity entity = stylishService.getStylishById(request.getStylishId());
        if (entity == null) {
            return ResponseEntity.status(404).body("Stylish with id " + request.getStylishId() + " does not exist.");
        }
        return ResponseEntity.ok(entity);
    }
}
