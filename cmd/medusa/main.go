package main

import (
	"context"
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/db"
	"florianbgt/medusa/internal/gcode"
	"florianbgt/medusa/internal/handlers/stream"
	"florianbgt/medusa/internal/routing"
)

func main() {
	configs := configs.SetupConfigs()

	db_path := "medusa.db"
	db := db.SetupDb(db_path, configs)

	if configs.ENABLE_CAMERA {
		camera := stream.SetupCamera(configs.CAMERA_NAME)

		defer camera.Close()

		if err := camera.Start(context.TODO()); err != nil {
			panic(err)
		}

		stream.Frames = camera.GetOutput()
	}

	gcodeSender := gcode.NewGCodeSender(configs.PRINTER_NAME, configs.PRINTER_BAUD)

	router := routing.SetupRouter(db, configs, gcodeSender)

	router.Run(":" + configs.PORT)
}
