package dto

import "time"

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
	Desc           string    `json:"desc"`
	Level          QuizLevel `json:"level"`
	Estimate       int       `json:"estimate"`
	MinimumScore   int       `json:"minimum_score"`
	TotalQuestions int       `json:"total_questions"`

	Category *CategoryRes `json:"category"`
}

type QuizDetailRes struct {
	UUID           string    `json:"uuid"`
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Desc           string    `json:"desc"`
	Level          QuizLevel `json:"level"`
	Estimate       int       `json:"estimate"`
	MinimumScore   int       `json:"minimum_score"`
	TotalQuestions int       `json:"total_questions"`

	Questions []QuizQuestionRes `json:"questions"`
	Category  *CategoryRes      `json:"category"`
}

type QuizQuestionRes struct {
	Question    string   `json:"question" validate:"required"`
	Options     []string `json:"options" validate:"required"`
	AnswerIndex int      `json:"answer_index" validate:"number,gte=0"`
}

type QuizPublicDetailRes struct {
	UUID           string    `json:"uuid"`
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Desc           string    `json:"desc"`
	Level          QuizLevel `json:"level"`
	Estimate       int       `json:"estimate"`
	MinimumScore   int       `json:"minimum_score"`
	TotalQuestions int       `json:"total_questions"`

	Questions []QuizPublicQuestionRes `json:"questions"`
	Category  *CategoryRes      `json:"category"`
}

type QuizPublicQuestionRes struct {
	Question    string   `json:"question" validate:"required"`
	Options     []string `json:"options" validate:"required"`
}

type AnalyzeStatus string

const (
	Success AnalyzeStatus = "success"
	Failed  AnalyzeStatus = "failed"
)

type AnalyzeRes struct {
	Score  float32       `json:"score"`
	Status AnalyzeStatus `json:"match"`

	Certificate *CertificateRes `json:"certificate,omitempty"`
}

type CertificateRes struct {
	UUID           string    `json:"uuid"`
	Score          float32   `json:"score"`
	CertificateURL string    `json:"certificate_url"`
	CreatedAt      time.Time `json:"created_at"`
}
