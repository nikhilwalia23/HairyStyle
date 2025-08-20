package com.example.shopservice.dto;

public class StylishRequest {
    private Long stylishId;
    private int experience;
    private String description;
    private java.util.List<String> imageUrls;

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
    public java.util.List<String> getImageUrls() {
        return imageUrls;
    }
    public void setImageUrls(java.util.List<String> imageUrls) {
        this.imageUrls = imageUrls;
    }
}
