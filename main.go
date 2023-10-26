package main

import (
	"flag"
	"gp-quizz/pkg/csv"
	"gp-quizz/pkg/quizz"
	"gp-quizz/pkg/ui"
	"log"
)

func main() {
	//tmr := time.NewTimer(2 * time.Minute)
	//
	//select {
	//case <-tmr.C:
	//	// timeout
	//}

	timeout := flag.Int("t", 30, "Quizz timeout in seconds.")
	flag.Parse()

	questions, err := csv.ReadQuestions("./problems.csv")
	if err != nil {
		log.Panic(err)
	}

	qz := quizz.NewQuizz(questions)
	qui := ui.NewUI(qz, *timeout)

	qui.Run()
}
