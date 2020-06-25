package controllers

import (
    "gqlgen-todos/Test/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

type CreateAuthorInput struct {
    Name  string `json:"name" binding:"required"`
    Book string `json:"book" binding:"required"`
}

type UpdateAuthorInput struct {
    Name  string `json:"title"`
    Book string `json:"book"`
}

// GET /Authors
// Find all Authors
func FindAuthors(c *gin.Context) {
    var Authors []models.Author
    models.DB.Find(&Authors)

    c.JSON(http.StatusOK, gin.H{"data": Authors})
}

// GET /Authors/:id
// Find a Author
func FindAuthor(c *gin.Context) {
    // Get model if exist
    var Author models.Author
    if err := models.DB.Where("id = ?", c.Param("id")).First(&Author).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": Author})
}

// POST /Authors
// Create new Author
func CreateAuthor(c *gin.Context) {
    // Validate input
    var input CreateAuthorInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Author
    Author := models.Author{Name: input.Name, Book: input.Book}
    models.DB.Create(&Author)

    c.JSON(http.StatusOK, gin.H{"data": Author})
}

//// PATCH /Authors/:id
//// Update a Author
//func UpdateAuthor(c *gin.Context) {
//    // Get model if exist
//    var Author models.Author
//    if err := models.DB.Where("id = ?", c.Param("id")).First(&Author).Error; err != nil {
//        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
//        return
//    }
//
//    // Validate input
//    var input UpdateAuthorInput
//    if err := c.ShouldBindJSON(&input); err != nil {
//        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//        return
//    }
//
//    models.DB.Model(&Author).Updates(input)
//
//    c.JSON(http.StatusOK, gin.H{"data": Author})
//}
//
//// DELETE /Authors/:id
//// Delete a Author
//func DeleteAuthor(c *gin.Context) {
//    // Get model if exist
//    var Author models.Author
//    if err := models.DB.Where("id = ?", c.Param("id")).First(&Author).Error; err != nil {
//        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
//        return
//    }
//
//    models.DB.Delete(&Author)
//
//    c.JSON(http.StatusOK, gin.H{"data": true})
//}
