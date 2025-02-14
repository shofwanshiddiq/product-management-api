package models

type CreateFileRequest struct {
	FileId  string `json:"file_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type DownloadFileRequest struct {
	FileId string `json:"file_id" binding:"required"`
}
