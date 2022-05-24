package models

const (
	QuestionEasy   = iota // 0
	QuestionMedium = iota // 1
	QuestionHard   = iota // 2
)

type Question struct {
	ID           int16
	Statement    string
	InputFormat  string
	OutputFormat string
	Level        int  // easy, medium or hard
	SampleIO     []IO // examples
}
