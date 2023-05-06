package control

import (
	"florianbgt/medusa/internal/gcode"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Move(c *gin.Context, gcodeSender *gcode.GCodeSender) {
	var payload struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	gcodeSender.RelativeMove(payload.X, payload.Y, payload.Z)
	c.JSON(http.StatusOK, nil)
}

func Home(c *gin.Context, gcodeSender *gcode.GCodeSender) {
	gcodeSender.Send("G28")
	c.JSON(http.StatusOK, nil)
}

func Level(c *gin.Context, gcodeSender *gcode.GCodeSender) {
	gcodeSender.Send("G29")
	c.JSON(http.StatusOK, nil)
}

func Extrude(c *gin.Context, gcodeSender *gcode.GCodeSender) {
	var payload struct {
		E float64 `json:"e"`
	}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	gcodeSender.Send("G91") // set relative mode
	gcode := fmt.Sprintf("G1 E%v", strconv.FormatFloat(payload.E, 'f', 2, 64))
	gcodeSender.Send(gcode)
	gcodeSender.Send("G90") // set absolute mode
	c.JSON(http.StatusOK, nil)
}

func GetTemperatures(c *gin.Context, gcodeSender *gcode.GCodeSender) {
	out := gcodeSender.Send("M105")
	fmt.Println(out)
	regex := regexp.MustCompile(`T:(-?[0-9.]+) /(-?[0-9.]+) B:(-?[0-9.]+) /(-?[0-9.]+)`)
	fmt.Println(regex.FindStringSubmatch(out))
	hotEndTemp, _ := strconv.ParseFloat(regex.FindStringSubmatch(out)[1], 64)
	bedTemp, _ := strconv.ParseFloat(regex.FindStringSubmatch(out)[3], 64)

	c.JSON(http.StatusOK, gin.H{
		"hot_temp": hotEndTemp,
		"bed_temp":    bedTemp,
	})
}

func SetExtruderTemp(c *gin.Context, gcodeSender *gcode.GCodeSender) {
	var payload struct {
		Temperature float64 `json:"temperature" binding:"required"`
	}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	gcode := fmt.Sprintf("M104 S%v", strconv.FormatFloat(payload.Temperature, 'f', 2, 64))
	gcodeSender.Send(gcode)

	c.JSON(http.StatusOK, nil)
}
