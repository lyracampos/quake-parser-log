package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/lyracampos/quake-log-parser/internal/logger"
	"github.com/lyracampos/quake-log-parser/internal/usecases"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name               string
		filePath           string
		expectedStatusCode int
	}{
		{
			name:               "should return ok when parsing log successfully",
			filePath:           "../../static/logs/test_valid.log",
			expectedStatusCode: 200,
		},
		{
			name:               "should return band request when the file log type is invalid",
			filePath:           "../../static/logs/test_invalid_type.html",
			expectedStatusCode: 400,
		},
		{
			name:               "should return band request when the file log is invalid",
			filePath:           "../../static/logs/test_empty_file.log",
			expectedStatusCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Open the file
			file, err := os.Open(tt.filePath)
			if err != nil {
				t.Fatalf("Error opening file: %v", err)
			}
			defer file.Close()

			logger.InitLog()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("qgames", filepath.Base(file.Name()))
			if err != nil {
				t.Fatalf("Error creating form file: %v", err)
			}

			_, err = io.Copy(part, file)
			if err != nil {
				t.Fatalf("Error copying file to part: %v", err)
			}
			writer.Close()

			// Create a request
			req, err := http.NewRequest("POST", "/parse", body)
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}
			req.Header.Add("Content-Type", writer.FormDataContentType())

			// Create a ResponseRecorder
			rr := httptest.NewRecorder()

			parseLogUseCase := usecases.NewParseLogUseCase()
			h := NewHandler(parseLogUseCase)
			handler := http.HandlerFunc(h.Parse)

			// Send the request to the handler
			handler.ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatusCode, rr.Code)
		})
	}
}
