package com.example.shopservice.service;

import com.example.shopservice.dto.StylishRequest;
import com.example.shopservice.model.StylishEntity;
import com.example.shopservice.repository.StylishRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import java.util.Optional;
import java.util.Arrays;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class StylishServiceTest {
    @Mock
    private StylishRepository stylishRepository;
    @InjectMocks
    private StylishService stylishService;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void testAddOrUpdateStylish_New() {
        StylishRequest request = new StylishRequest();
        request.setStylishId(1L);
        request.setExperience(5);
        request.setDescription("desc");
        request.setImageUrls(Arrays.asList("img1.jpg", "img2.jpg"));
        when(stylishRepository.findById(1L)).thenReturn(Optional.empty());
        StylishEntity entity = new StylishEntity();
        when(stylishRepository.save(any())).thenReturn(entity);
        StylishEntity result = stylishService.addOrUpdateStylish(request);
        assertEquals(entity, result);
        verify(stylishRepository).save(any());
    }

    @Test
    void testAddOrUpdateStylish_Update() {
        StylishRequest request = new StylishRequest();
        request.setStylishId(2L);
        request.setExperience(10);
        request.setDescription("updated");
        request.setImageUrls(Arrays.asList("img3.jpg"));
        StylishEntity existing = new StylishEntity();
        when(stylishRepository.findById(2L)).thenReturn(Optional.of(existing));
        when(stylishRepository.save(existing)).thenReturn(existing);
        StylishEntity result = stylishService.addOrUpdateStylish(request);
        assertEquals(existing, result);
        verify(stylishRepository).save(existing);
    }
}
