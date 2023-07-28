package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shriker95/notes-backend/models"
)

func GetNotes(c *gin.Context) {
	var notes []models.Note
	models.DB.Find(&notes)

	c.JSON(http.StatusOK, gin.H{"data": notes})
}

type CreateNoteInput struct {
	Text string `json:"text" binding:"required"`
}

func CreateNote(c *gin.Context) {
	var input CreateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := models.Note{TEXT: input.Text}
	models.DB.Create(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})
}

func GetNote(c *gin.Context) { // Get model if exist
	var note models.Note

	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Note with this ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": note})
}

type UpdateNoteInput struct {
	Text string `json:"text"`
}

func UpdateNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Note not found!"})
		return
	}

	var input UpdateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&note).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": note})
}

func DeleteNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Note not found!"})
		return
	}

	models.DB.Delete(&note)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
