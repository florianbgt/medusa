package files

import (
	"bufio"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const Directory = "uploads/"

type FileInfo struct {
	FileName     string `json:"file_name"`
	TotalTime    string `json:"total_time"`
	FilamentUsed string `json:"filament_used"`
	LayerHeight  string `json:"layer_height"`
	LayerCount   string `json:"layer_count"`
	NozzleTemp   string `json:"nozzle_temp"`
	BedTemp      string `json:"bed_temp"`
}

func renameFile(name string) string {
	extension := filepath.Ext(name)
	baseName := strings.TrimSuffix(name, extension)

	_, err := os.Open(Directory + name)
	if errors.Is(err, os.ErrNotExist) {
		return name
	}

	index := 0
	for {
		index++
		newName := baseName + "_" + strconv.Itoa(index) + extension
		_, err := os.Open(Directory + newName)
		if errors.Is(err, os.ErrNotExist) {
			return newName
		}
	}
}

func getGCodeInfo(name string) (FileInfo, error) {
	var fileInfo = FileInfo{
		FileName: name,
	}

	file, err := os.Open(Directory + name)
	if os.IsNotExist(err) {
		return fileInfo, errors.New("file_not_found")
	} else if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		valuesMap := map[string]*string{
			";TIME:":           &fileInfo.TotalTime,
			";Filament used: ": &fileInfo.FilamentUsed,
			";Layer height: ":  &fileInfo.LayerHeight,
			";LAYER_COUNT:":    &fileInfo.LayerCount,
			"M104 S":           &fileInfo.NozzleTemp,
			"M140 S":           &fileInfo.BedTemp,
		}

		for key, value := range valuesMap {
			if *value == "" && strings.HasPrefix(line, key) {
				*value = strings.TrimPrefix(line, key)
			}
		}

		check := 0
		for _, value := range valuesMap {
			if *value != "" {
				check += 1
			}
		}
		if check == len(valuesMap) {
			break
		}
	}

	return fileInfo, nil
}

func ListFiles(c *gin.Context) {
	err := os.MkdirAll(Directory, os.ModePerm)
	if err != nil {
		panic(err)
	}

	files, err := os.ReadDir(Directory)
	if err != nil {
		panic(err)
	}

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
	if os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	if err != nil {
		panic(err)
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

	name := renameFile(uploadedFile.Filename)

	c.SaveUploadedFile(uploadedFile, Directory+name)

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

func GetGCodeInfo(c *gin.Context) {
	fileName := c.Param("name")

	fileInfo, err := getGCodeInfo(fileName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, fileInfo)
}
