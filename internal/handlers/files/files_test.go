package files_test

import (
	"encoding/json"
	"florianbgt/medusa/internal/handlers/files"
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/internal/models/password_model"
	"florianbgt/medusa/test"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFilesRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/files"
	db := test.Setupdb()

	var passwordInstance password_model.Password

	configs := test.SetupConfigs()

	passwordInstance.Setup(db, configs.DEFAULT_PASSWORD, configs.API_KEY)

	os.RemoveAll(files.Directory)

	t.Run("list files empty", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)
		req, _ := http.NewRequest("GET", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 0, len(response))

		os.RemoveAll(files.Directory)
	})

	// t.Run("list files properties", func(t *testing.T) {
	// 	w := httptest.NewRecorder()

	// 	os.MkdirAll(files.Directory, os.ModePerm)
	// 	os.Create(files.Directory + "/test.gcode")

	// 	token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)
	// 	req, _ := http.NewRequest("GET", route, nil)
	// 	req.Header = http.Header{
	// 		"Authorization": []string{"Bearer " + token_pair.Access},
	// 	}
	// 	api.ServeHTTP(w, req)

	// 	response := make(map[string]string)
	// 	json.Unmarshal(w.Body.Bytes(), &response)

	// 	assert.Equal(t, http.StatusOK, w.Code)
	// 	assert.Equal(t, 1, len(response))

	// 	// os.RemoveAll(files.Directory)
	// })
}

// func TestUploadFileRoute(t *testing.T) {
// 	api := test.SetupApi()
// 	route := "/api/files"
// 	db := test.Setupdb()

// 	var passwordInstance password_model.Password

// 	configs := test.SetupConfigs()

// 	passwordInstance.Setup(db, configs.DEFAULT_PASSWORD, configs.API_KEY)

// 	os.RemoveAll(files.Directory)

// 	t.Run("upload a file", func(t *testing.T) {
// 		w := httptest.NewRecorder()

// 		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

// 		body := new(bytes.Buffer)
// 		writer := multipart.NewWriter(body)
// 		part, _ := writer.CreateFormFile("file", "testfile.gcode")
// 		sample, _ := os.Open("testfile.gcode")
// 		io.Copy(part, sample)

// 		req, _ := http.NewRequest("POST", route, body)
// 		req.Header = http.Header{
// 			"Authorization": []string{"Bearer " + token_pair.Access},
// 			"Content-Type":  []string{writer.FormDataContentType()},
// 		}
// 		api.ServeHTTP(w, req)

// 		response := make(map[string]string)
// 		json.Unmarshal(w.Body.Bytes(), &response)

// 		fmt.Println(response)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		// assert.Equal(t, 0, len(response))
// 	})
// }
