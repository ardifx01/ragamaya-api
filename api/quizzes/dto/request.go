package dto

type QuizLevel string

const (
	Beginner     QuizLevel = "beginner"
	Intermediate QuizLevel = "intermediate"
	Advanced     QuizLevel = "advanced"
)

type QuizReq struct {
	Title        string    `json:"title" validate:"required"`
	Level        QuizLevel `json:"level" validate:"required,oneof=beginner intermediate advanced"`
	Category     string    `json:"category" validate:"required"`
	Estimate     int       `json:"estimate" validate:"required"`
	MinimumScore int       `json:"minimum_score" validate:"required,gte=0,lte=100"`

	Questions []QuizQuestionReq `json:"questions" validate:"required,dive"`
}

type QuizQuestionReq struct {
	Question    string   `json:"question" validate:"required"`
	Options     []string `json:"options" validate:"required"`
	AnswerIndex int      `json:"answer_index" validate:"number,gte=0"`
}
