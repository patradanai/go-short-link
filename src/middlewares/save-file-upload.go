package middlewares

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func SaveFileUpload(imageField string) gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

		files := form.File[imageField]
		if len(files) > 0 {
			for _, file := range files {
				getUUid, err := gonanoid.New(9)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
					return
				}

				fileName := getUUid + filepath.Ext(file.Filename)

				// Create Folder

				outPath := filepath.Join("./public", imageField)
				if _, err := os.Stat(outPath); os.IsNotExist(err) {
					os.Mkdir(outPath, os.ModePerm)
				}

				destFile := filepath.Join(outPath, fileName)
				if err := c.SaveUploadedFile(file, destFile); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
					return
				}

				c.Set("file", map[string]string{"original": file.Filename, "filename": fileName, "path": destFile})

			}
		}
		c.Next()
	}
}
