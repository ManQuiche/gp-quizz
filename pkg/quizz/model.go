package quizz

type Question struct {
	Desc   string
	Answer string
}

type quizz struct {
	questions  []Question
	current    int
	score      int
	terminated bool
}

type Quizz interface {
	Stop()

	Terminated() bool
	Questions() []Question
	Current() string
	Score() int
	Check(answer string) bool
}
