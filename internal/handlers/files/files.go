package files

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const Directory = "uploads/"

func renamePath(path string) string {
	_, err := os.Open(path)

	if errors.Is(err, os.ErrNotExist) {
		return path
	} else {
		index := 0
		for {
			index++
			_, err := os.Open(path + "_" + strconv.Itoa(index))
			if errors.Is(err, os.ErrNotExist) {
				break
			}
		}
		path = path + "_" + strconv.Itoa(index)
	}
	return path
}

func readAllFiles() []os.DirEntry {
	err := os.MkdirAll(Directory, os.ModePerm)
	if err != nil {
		panic(err)
	}

	files, err := os.ReadDir(Directory)
	if err != nil {
		panic(err)
	}

	return files
}

func ListFiles(c *gin.Context) {
	files := readAllFiles()

	type FileInfo struct {
		Name     string `json:"name"`
		Uploaded string `json:"uploaded"`
		Size     int64  `json:"size"`
	}

	filesInfo := make([]FileInfo, 0)
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		fileInfo := FileInfo{
			Name:     info.Name(),
			Uploaded: info.ModTime().Format("2006-01-02 15:04:05"),
			Size:     info.Size(),
		}
		filesInfo = append(filesInfo, fileInfo)
	}

	c.JSON(http.StatusOK, filesInfo)
}

func DeleteFile(c *gin.Context) {
	fileName := c.Param("name")
	err := os.Remove(Directory + fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func UploadFile(c *gin.Context) {
	uploadedFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	path := renamePath(Directory + uploadedFile.Filename)

	c.SaveUploadedFile(uploadedFile, path)

	c.JSON(http.StatusOK, nil)
}

func GetGCode(c *gin.Context) {
	fileName := c.Param("name")

	gcode, err := os.ReadFile(Directory + fileName)
	if os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	} else if err != nil {
		panic(err)
	}

	c.Data(http.StatusOK, "text/plain", gcode)
}
