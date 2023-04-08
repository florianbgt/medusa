package main

import (
	"context"
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/db"
	"florianbgt/medusa/internal/handlers/stream"
	"florianbgt/medusa/internal/routing"
)

func main() {
	configs := configs.SetupConfigs()

	camera := stream.SetupCamera()

	defer camera.Close()

	if err := camera.Start(context.TODO()); err != nil {
		panic(err)
	}

	stream.Frames = camera.GetOutput()

	db_path := configs.SQLITE_FILE_PATH
	db := db.SetupDb(db_path, configs)

	router := routing.SetupRouter(db, configs)

	router.Run(":" + configs.PORT)
}
