package dto

type FilesInput struct {
	OriginalFileName string
	FileBuffer       []byte
	Size             string
	Extension        string
	MimeType         string
	MimeSubType      string
	Meta             string
}
