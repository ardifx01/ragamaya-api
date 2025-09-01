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

type QuizRes struct {
	UUID           string    `json:"uuid"`
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Level          QuizLevel `json:"level"`
	Estimate       int       `json:"estimate"`
	MinimumScore   int       `json:"minimum_score"`
	TotalQuestions int       `json:"total_questions"`

	Category *CategoryRes `json:"category"`
}

type QuizQuestionRes struct {
	Question    string   `json:"question" validate:"required"`
	Options     []string `json:"options" validate:"required"`
	AnswerIndex int      `json:"answer_index" validate:"number,gte=0"`
}
