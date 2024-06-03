package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lyracampos/quake-log-parser/internal/logger"
	"github.com/lyracampos/quake-log-parser/internal/usecases"
)

var ErrInvalidFileType = errors.New("invalid file type, expected .log")

type Handler struct {
	parseLogUseCase *usecases.ParseLogUseCase
}

func NewHandler(parseLogUseCase *usecases.ParseLogUseCase) *Handler {
	return &Handler{
		parseLogUseCase: parseLogUseCase,
	}
}

func (h *Handler) Parse(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("HandlerParse - parse begin")

	ctx := r.Context()

	file, err := h.readFileFromRequest(r)
	if err != nil {
		h.handlerError(w, fmt.Errorf("failed on reading file from request: %w", err))
		return
	}

	tempFile, err := h.createTemporaryFile(file)
	if err != nil {
		h.handlerError(w, fmt.Errorf("failed on creating temporary file: %w", err))
		return
	}

	games, err := h.parseLogUseCase.Execute(ctx, tempFile.Name())
	if err != nil {
		h.handlerError(w, fmt.Errorf("failed on executing parse log use case: %w", err))
		return
	}

	gamesJSON, err := json.Marshal(games)
	if err != nil {
		h.handlerError(w, fmt.Errorf("failed on marshalling games: %w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(gamesJSON)
	if err != nil {
		logger.Log.Errorf("Error on writing the error message: %v", err)
	}

	logger.Log.Info("HandlerParse - parse finished")
}

func (h *Handler) readFileFromRequest(r *http.Request) (multipart.File, error) {
	file, header, err := r.FormFile("qgames")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if filepath.Ext(header.Filename) != ".log" {
		return nil, ErrInvalidFileType
	}

	return file, nil
}

func (h *Handler) createTemporaryFile(file multipart.File) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "upload-*.log")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}

func (h *Handler) handlerError(w http.ResponseWriter, err error) {
	logger.Log.Error(err.Error())

	msgErr := ""
	statusCode := http.StatusInternalServerError

	switch {
	case errors.Is(err, ErrInvalidFileType):
		msgErr = err.Error()
		statusCode = http.StatusBadRequest
	default:
		msgErr = err.Error()
	}

	w.WriteHeader(statusCode)
	_, err = w.Write([]byte(msgErr))
	if err != nil {
		logger.Log.Errorf("Error on writing the error message: %v", err)
	}
}
