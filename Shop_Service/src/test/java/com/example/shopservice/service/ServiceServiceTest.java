package com.example.shopservice.service;

import com.example.shopservice.dto.ServiceRequest;
import com.example.shopservice.model.ServiceEntity;
import com.example.shopservice.repository.ServiceRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class ServiceServiceTest {
    @Mock
    private ServiceRepository serviceRepository;
    @InjectMocks
    private ServiceService serviceService;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void testAddService_Success() {
        ServiceRequest request = new ServiceRequest("Haircut", "Basic haircut", 20.0, 30);
        when(serviceRepository.findByServiceName("Haircut")).thenReturn(Optional.empty());
        ServiceEntity entity = ServiceEntity.builder()
                .serviceName("Haircut")
                .description("Basic haircut")
                .price(20.0)
                .duration(30)
                .build();
        when(serviceRepository.save(any(ServiceEntity.class))).thenReturn(entity);
        ServiceEntity result = serviceService.addService(request);
        assertEquals(entity.getServiceName(), result.getServiceName());
        assertEquals(entity.getDescription(), result.getDescription());
        assertEquals(entity.getPrice(), result.getPrice());
        assertEquals(entity.getDuration(), result.getDuration());
        verify(serviceRepository).findByServiceName("Haircut");
        verify(serviceRepository).save(any(ServiceEntity.class));
    }

    @Test
    void testAddService_DuplicateName() {
        ServiceRequest request = new ServiceRequest("Haircut", "Basic haircut", 20.0, 30);
        ServiceEntity existing = ServiceEntity.builder()
                .serviceName("Haircut")
                .description("Basic haircut")
                .price(20.0)
                .duration(30)
                .build();
        when(serviceRepository.findByServiceName("Haircut")).thenReturn(Optional.of(existing));
        assertThrows(IllegalArgumentException.class, () -> serviceService.addService(request));
        verify(serviceRepository).findByServiceName("Haircut");
        verify(serviceRepository, never()).save(any(ServiceEntity.class));
    }
}
