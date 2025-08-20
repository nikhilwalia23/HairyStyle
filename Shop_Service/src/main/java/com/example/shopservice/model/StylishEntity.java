package com.example.shopservice.model;

import jakarta.persistence.*;
import java.util.List;

@Entity
@Table(name = "stylish")
public class StylishEntity {
    @Id
    private Long stylishId; // Provided by another microservice
    private int experience;
    private String description;
    @ElementCollection
    private List<String> imageUrls;

    public Long getStylishId() {
        return stylishId;
    }
    public void setStylishId(Long stylishId) {
        this.stylishId = stylishId;
    }
    public int getExperience() {
        return experience;
    }
    public void setExperience(int experience) {
        this.experience = experience;
    }
    public String getDescription() {
        return description;
    }
    public void setDescription(String description) {
        this.description = description;
    }
    public List<String> getImageUrls() {
        return imageUrls;
    }
    public void setImageUrls(List<String> imageUrls) {
        this.imageUrls = imageUrls;
    }
}
