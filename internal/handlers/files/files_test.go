package files_test

import (
	"bytes"
	"encoding/json"
	"florianbgt/medusa/internal/handlers/files"
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/test"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func cleanup() {
	os.RemoveAll(files.Directory)
}

func uploadFileRequestHelper(w *httptest.ResponseRecorder, api *gin.Engine, token string) {
	route := "/api/files"

	file, _ := os.Open("testfile.gcode")
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "testfile.gcode")
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", route, body)
	req.Header = http.Header{
		"Authorization": []string{"Bearer " + token},
		"Content-Type":  []string{writer.FormDataContentType()},
	}
	api.ServeHTTP(w, req)
}

func TestListFilesRoute(t *testing.T) {
	api := test.SetupApi()
	route := "/api/files"

	configs := test.SetupConfigs()

	cleanup()

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

		cleanup()
	})

	t.Run("list files success", func(t *testing.T) {
		w := httptest.NewRecorder()

		os.MkdirAll(files.Directory, os.ModePerm)
		testFile, _ := os.Open("testfile.gcode")
		defer testFile.Close()
		uploadedFile, _ := os.Create(files.Directory + "testfile.gcode")
		defer uploadedFile.Close()
		uploadedFile.ReadFrom(testFile)

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)
		req, _ := http.NewRequest("GET", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		response := make([]map[string]interface{}, 0)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 1, len(response))

		fileInfo, _ := uploadedFile.Stat()

		assert.Equal(t, "testfile.gcode", response[0]["name"])
		assert.Equal(t, fileInfo.ModTime().Format("2006-01-02 15:04:05"), response[0]["uploaded"])
		assert.Equal(t, float64(fileInfo.Size()), response[0]["size"])

		cleanup()
	})
}

func TestDeleteFileRoute(t *testing.T) {
	api := test.SetupApi()

	configs := test.SetupConfigs()

	cleanup()

	t.Run("delete a file success", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		uploadFileRequestHelper(w, api, token_pair.Access)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)

		uploaded_files, _ := os.ReadDir(files.Directory)
		assert.Equal(t, 1, len(uploaded_files))

		route := "/api/files/" + uploaded_files[0].Name()

		req, _ := http.NewRequest("DELETE", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		response = make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)

		uploaded_files, _ = os.ReadDir(files.Directory)
		assert.Equal(t, 0, len(uploaded_files))

		cleanup()
	})

	t.Run("delete a file 404", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		route := "/api/files/idonotexist.gcode"

		req, _ := http.NewRequest("DELETE", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusNotFound, w.Code)

		cleanup()
	})
}

func TestUploadFileRoute(t *testing.T) {
	api := test.SetupApi()

	configs := test.SetupConfigs()

	cleanup()

	t.Run("upload a file success", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		uploadFileRequestHelper(w, api, token_pair.Access)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)

		uploaded_files, _ := os.ReadDir(files.Directory)
		assert.Equal(t, 1, len(uploaded_files))
		assert.Equal(t, "testfile.gcode", uploaded_files[0].Name())

		cleanup()
	})

	t.Run("upload same file twice success", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		uploadFileRequestHelper(w, api, token_pair.Access)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)

		uploaded_files, _ := os.ReadDir(files.Directory)
		assert.Equal(t, 1, len(uploaded_files))
		assert.Equal(t, "testfile.gcode", uploaded_files[0].Name())

		uploadFileRequestHelper(w, api, token_pair.Access)

		response = make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)

		uploaded_files, _ = os.ReadDir(files.Directory)
		assert.Equal(t, 2, len(uploaded_files))
		assert.Equal(t, "testfile_1.gcode", uploaded_files[1].Name())

		cleanup()
	})
}

func TestGetcodeRoute(t *testing.T) {
	api := test.SetupApi()

	configs := test.SetupConfigs()

	cleanup()

	t.Run("get a gcode success", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		uploadFileRequestHelper(w, api, token_pair.Access)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)

		uploaded_files, _ := os.ReadDir(files.Directory)
		assert.Equal(t, 1, len(uploaded_files))

		route := "/api/files/testfile.gcode/gcode"

		req, _ := http.NewRequest("GET", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		cleanup()
	})

	t.Run("get gcode 404", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		route := "/api/files/idonotexist.gcode/gcode"

		req, _ := http.NewRequest("GET", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestGetGCodeInfoRoute(t *testing.T) {
	api := test.SetupApi()

	configs := test.SetupConfigs()

	cleanup()

	t.Run("get gcode info success", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		uploadFileRequestHelper(w, api, token_pair.Access)

		uploaded_files, _ := os.ReadDir(files.Directory)
		assert.Equal(t, 1, len(uploaded_files))

		route := "/api/files/testfile.gcode/gcode/info"

		req, _ := http.NewRequest("GET", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		dec := json.NewDecoder(w.Body)
		var v interface{}
		dec.Decode(&v)
		var response map[string]string
		dec.Decode(&response)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, "7860", response["total_time"])
		assert.Equal(t, "1.43339m", response["filament_used"])
		assert.Equal(t, "0.1", response["layer_height"])
		assert.Equal(t, "478", response["layer_count"])
		assert.Equal(t, "200", response["nozzle_temp"])
		assert.Equal(t, "0", response["bed_temp"])

		cleanup()
	})

	t.Run("get gcode info 404", func(t *testing.T) {
		w := httptest.NewRecorder()

		token_pair, _ := helpers.GenerateTokenPair(configs.API_KEY)

		route := "/api/files/idonotexist.gcode/gcode/info"

		req, _ := http.NewRequest("GET", route, nil)
		req.Header = http.Header{
			"Authorization": []string{"Bearer " + token_pair.Access},
		}
		api.ServeHTTP(w, req)

		response := make(map[string]string)
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusNotFound, w.Code)

		cleanup()
	})
}
