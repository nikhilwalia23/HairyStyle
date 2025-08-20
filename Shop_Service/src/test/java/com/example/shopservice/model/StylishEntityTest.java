package com.example.shopservice.model;

import org.junit.jupiter.api.Test;
import java.util.Arrays;
import static org.junit.jupiter.api.Assertions.*;

class StylishEntityTest {
    @Test
    void testGettersAndSetters() {
        StylishEntity entity = new StylishEntity();
        entity.setStylishId(1L);
        entity.setExperience(5);
        entity.setDescription("desc");
        entity.setImageUrls(Arrays.asList("img1.jpg", "img2.jpg"));
        assertEquals(1L, entity.getStylishId());
        assertEquals(5, entity.getExperience());
        assertEquals("desc", entity.getDescription());
        assertEquals(Arrays.asList("img1.jpg", "img2.jpg"), entity.getImageUrls());
    }
}
