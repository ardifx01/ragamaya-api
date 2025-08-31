package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type CategoryRes struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type QuizRes struct{}