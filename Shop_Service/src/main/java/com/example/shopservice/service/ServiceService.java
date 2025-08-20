package com.example.shopservice.service;

import com.example.shopservice.dto.ServiceRequest;
import com.example.shopservice.model.ServiceEntity;
import com.example.shopservice.repository.ServiceRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class ServiceService {

    private final ServiceRepository serviceRepository;

    public ServiceService(ServiceRepository serviceRepository) {
        this.serviceRepository = serviceRepository;
    }

    @Transactional
    public ServiceEntity addService(ServiceRequest request) {
        serviceRepository.findByServiceName(request.serviceName()).ifPresent(s -> {
            throw new IllegalArgumentException("Service with this name already exists!");
        });

        ServiceEntity service = ServiceEntity.builder()
                .serviceName(request.serviceName())
                .description(request.description())
                .price(request.price())
                .duration(request.duration())
                .build();

        return serviceRepository.save(service);
    }
}
