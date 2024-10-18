package controller

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/avran02/decoplan/files/internal/config"
	"github.com/avran02/decoplan/files/internal/dto"
	"github.com/avran02/decoplan/files/internal/service"
	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type FilesController interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	DownloadFile(w http.ResponseWriter, r *http.Request)
	DeleteFile(w http.ResponseWriter, r *http.Request)
}

type filesController struct {
	Service service.FilesService
}

func (c filesController) DownloadFile(w http.ResponseWriter, r *http.Request) {
	slog.Info("filesController.DownloadFile")
	ctx := r.Context()
	fileID := chi.URLParam(r, "id")
	data, err := c.Service.DownloadFile(ctx, fileID)
	if err != nil {
		err = fmt.Errorf("failed to download file: %w", err)
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer data.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileID))

	if _, err = io.Copy(w, data); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c filesController) UploadFile(w http.ResponseWriter, r *http.Request) {
	slog.Info("filesController.UploadFile")

	ctx := r.Context()
	if err := r.ParseMultipartForm(config.StreamChunkSize); err != nil {
		err = fmt.Errorf("failed to parse multipart form: %w", err)
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("content")
	if err != nil {
		err = fmt.Errorf("failed to get file: %w", err)
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileID, err := c.Service.UploadFile(ctx, file, fileHeader.Filename)
	if err != nil {
		err = fmt.Errorf("failed to upload file: %w", err)
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.UploadFileResponse{ID: fileID}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c filesController) DeleteFile(w http.ResponseWriter, r *http.Request) {
	slog.Info("filesController.DeleteFile")

	ctx := r.Context()
	fileID := chi.URLParam(r, "id")

	if err := c.Service.DeleteFile(ctx, fileID); err != nil {
		err = fmt.Errorf("failed to delete file: %w", err)
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.DeleteFileResponse{Ok: true}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func (c filesController) asyncSendFile(stream pb.FileService_DownloadFileServer, file io.ReadCloser, streamErrChan chan error) {
// 	defer close(streamErrChan)
// 	defer file.Close()
// 	buf := make([]byte, config.StreamChunkSize)

// 	for {
// 		n, err := file.Read(buf)
// 		if err != nil {
// 			if errors.Is(err, io.EOF) {
// 				if n == 0 {
// 					break
// 				}
// 			} else {
// 				err = fmt.Errorf("failed to read file: %w", err)
// 				slog.Error(err.Error())
// 				streamErrChan <- err
// 			}
// 		}

// 		if err = stream.Send(&pb.DownloadFileResponse{
// 			Content: buf[:n],
// 		}); err != nil {
// 			streamErrChan <- fmt.Errorf("failed to send download file response: %w", err)
// 		}
// 	}

// 	if err := stream.Send(&pb.DownloadFileResponse{
// 		Success: true,
// 	}); err != nil {
// 		streamErrChan <- fmt.Errorf("failed to send download file response: %w", err)
// 	}
// }

// func (c filesController) asyncGetFileFromGrpcStream(stream pb.FileService_UploadFileServer, requestDTO *dto.UploadFileStreamRequest, streamErrChan chan error) {
// 	defer close(streamErrChan)
// 	defer requestDTO.CloseWriter()

// 	for {
// 		req, err := stream.Recv()
// 		if err != nil {
// 			if errors.Is(err, io.EOF) {
// 				break
// 			}

// 			err = fmt.Errorf("failed to receive upload file request: %w", err)
// 			slog.Error(err.Error())
// 			streamErrChan <- err
// 			return
// 		}

// 		_, err = requestDTO.Write(req.Content)
// 		if err != nil {
// 			err = fmt.Errorf("failed to write upload file request: %w", err)
// 			slog.Error(err.Error())
// 			streamErrChan <- err
// 			return
// 		}
// 	}
// }

func New(service service.FilesService) FilesController {
	slog.Info("initializing controller")
	return filesController{
		Service: service,
	}
}
