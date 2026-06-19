package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type File struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

const filePath = "data/files.json"

// LOAD
func loadFiles() []File {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return []File{}
	}

	var files []File
	json.Unmarshal(data, &files)
	return files
}

// SAVE
func saveFiles(files []File) {
	data, _ := json.MarshalIndent(files, "", "  ")
	os.WriteFile(filePath, data, 0644)
}

func main() {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// GET
	r.GET("/files", func(c *gin.Context) {
		c.JSON(200, loadFiles())
	})

	// CREATE
	r.POST("/files", func(c *gin.Context) {
		files := loadFiles()

		var newFile File
		c.BindJSON(&newFile)

		newFile.ID = len(files) + 1
		newFile.Status = "uploading"
		newFile.CreatedAt = time.Now().Format("2006-01-02 15:04")

		files = append(files, newFile)
		saveFiles(files)

		go simulateProcessing(newFile.ID)

		c.JSON(201, newFile)
	})

	r.DELETE("/files/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		files := loadFiles()

		for i, f := range files {
			if f.ID == id {
				files[i].Status = "deleted"
			}
		}

		saveFiles(files)

		c.JSON(200, gin.H{
			"message": "moved to trash",
			"id":      id,
		})
	})

	r.Run(":8080")
}

func simulateProcessing(id int) {

	time.Sleep(2 * time.Second)

	files := loadFiles()

	for i, f := range files {
		if f.ID == id {
			if f.Status != "deleted" {
				files[i].Status = "processing"
			}
		}
	}

	saveFiles(files)

	time.Sleep(3 * time.Second)

	files = loadFiles()

	for i, f := range files {
		if f.ID == id {
			if f.Status != "deleted" {
				files[i].Status = "done"
			}
		}
	}

	saveFiles(files)
}
