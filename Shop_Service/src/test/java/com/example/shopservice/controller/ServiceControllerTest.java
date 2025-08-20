package com.example.shopservice.controller;

import com.example.shopservice.dto.ServiceRequest;
import com.example.shopservice.dto.StylishRequest;
import com.example.shopservice.dto.StylishIdRequest;
import com.example.shopservice.model.ServiceEntity;
import com.example.shopservice.model.StylishEntity;
import com.example.shopservice.service.ServiceService;
import com.example.shopservice.service.StylishService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.http.ResponseEntity;
import java.util.Arrays;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class ServiceControllerTest {
    @Mock
    private ServiceService serviceService;
    @Mock
    private StylishService stylishService;
    @InjectMocks
    private ServiceController serviceController;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void testCreateService() {
        ServiceRequest request = new ServiceRequest(
            "Haircut", // serviceName
            "Basic haircut service", // description
            20.0, // price
            30 // duration
        );
        ServiceEntity entity = new ServiceEntity();
        when(serviceService.addService(request)).thenReturn(entity);
        ResponseEntity<ServiceEntity> response = serviceController.createService(request);
        assertEquals(entity, response.getBody());
        verify(serviceService).addService(request);
    }

    @Test
    void testAddOrUpdateStylish() {
        StylishRequest request = new StylishRequest();
        request.setStylishId(1L);
        request.setExperience(5);
        request.setDescription("desc");
        request.setImageUrls(Arrays.asList("img1.jpg", "img2.jpg"));
        StylishEntity entity = new StylishEntity();
        when(stylishService.addOrUpdateStylish(request)).thenReturn(entity);
        ResponseEntity<StylishEntity> response = serviceController.addOrUpdateStylish(request);
        assertEquals(entity, response.getBody());
        verify(stylishService).addOrUpdateStylish(request);
    }

    @Test
    void testViewStylishById() {
        StylishIdRequest request = new StylishIdRequest();
        request.setStylishId(1L);
        StylishEntity entity = new StylishEntity();
        when(stylishService.getStylishById(1L)).thenReturn(entity);
        ResponseEntity<?> response = serviceController.viewStylishById(request);
        assertEquals(entity, response.getBody());
        verify(stylishService).getStylishById(1L);
    }
}
