package com.example.shopservice.service;

import com.example.shopservice.dto.StylishRequest;
import com.example.shopservice.model.StylishEntity;
import com.example.shopservice.repository.StylishRepository;
import org.springframework.stereotype.Service;

@Service
public class StylishService {
    private final StylishRepository stylishRepository;

    public StylishService(StylishRepository stylishRepository) {
        this.stylishRepository = stylishRepository;
    }

    public StylishEntity addOrUpdateStylish(StylishRequest request) {
        StylishEntity entity = stylishRepository.findById(request.getStylishId())
            .orElse(new StylishEntity());
        entity.setStylishId(request.getStylishId());
        entity.setExperience(request.getExperience());
        entity.setDescription(request.getDescription());
        entity.setImageUrls(request.getImageUrls());
        return stylishRepository.save(entity);
    }

    public StylishEntity getStylishById(Long stylishId) {
        return stylishRepository.findById(stylishId).orElse(null);
    }
}
