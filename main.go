package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shriker95/notes-backend/controllers"
	"github.com/shriker95/notes-backend/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "v9000"})
	})

	r.GET("/notes", controllers.GetNotes)
	r.POST("/note", controllers.CreateNote)
	r.GET("/note/:id", controllers.GetNote)
	r.PATCH("/note/:id", controllers.UpdateNote)
	r.DELETE("/note/:id", controllers.DeleteNote)

	err := r.Run()
	if err != nil {
		return
	}
}
