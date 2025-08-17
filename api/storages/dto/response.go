package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type FilesOutput struct {
	UUID             string    `json:"uuid"`
	PublicURL        string    `json:"public_url"`
	OriginalFileName string    `json:"original_file_name"`
	Size             string    `json:"size"`
	Extension        string    `json:"extension"`
	MimeType         string    `json:"mime_type"`
	MimeSubType      string    `json:"mime_sub_type"`
	Meta             string    `json:"meta"`
	CreatedAt        time.Time `json:"created_at"`
}