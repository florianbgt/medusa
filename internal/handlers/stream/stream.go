package stream

import (
	"florianbgt/medusa/internal/helpers"
	"fmt"
	"log"
	"mime/multipart"
	"net/textproto"

	"github.com/gin-gonic/gin"
	"github.com/vladimirvivien/go4vl/device"
	"github.com/vladimirvivien/go4vl/v4l2"
)

var (
	Frames <-chan []byte
)

func SetupCamera() *device.Device {
	camera, err := device.Open(
		"/dev/video0",
		device.WithPixFormat(v4l2.PixFormat{PixelFormat: v4l2.PixelFmtMJPEG, Width: 640, Height: 480}),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to open device: %s", err))
	}

	return camera
}

func Stream(c *gin.Context, api_key string) {

	query := c.Request.URL.Query()
	token := query.Get("token")
	if !helpers.IsTokenValid(token, api_key) {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	mimeWriter := multipart.NewWriter(c.Writer)
	c.Writer.Header().Set("Content-Type", fmt.Sprintf("multipart/x-mixed-replace; boundary=%s", mimeWriter.Boundary()))
	partHeader := make(textproto.MIMEHeader)
	partHeader.Add("Content-Type", "image/jpeg")

	var frame []byte
	for frame = range Frames {
		partWriter, err := mimeWriter.CreatePart(partHeader)
		if err != nil {
			log.Printf("failed to create multi-part writer: %s", err)
			return
		}

		if _, err := partWriter.Write(frame); err != nil {
			log.Printf("failed to write image: %s", err)
		}
	}
}
