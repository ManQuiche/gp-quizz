package ui

import (
	"bufio"
	"fmt"
	"gp-quizz/pkg/quizz"
	"log"
	"os"
	"time"
)

type UI struct {
	quizz   quizz.Quizz
	timeout int
}

func NewUI(q quizz.Quizz, timeout int) UI {
	return UI{quizz: q, timeout: timeout}
}

func (ui UI) Run() {
	select {
	case <-ui.askMany():
		fmt.Printf("\n\nQuizz done ! Score: %d", ui.quizz.Score())
	case <-time.After(time.Duration(ui.timeout) * time.Second):
		fmt.Printf("\n\nQuizz timed out ! Try again ... Score: %d", ui.quizz.Score())
	}

	return
}

func (ui UI) askMany() <-chan bool {
	done := make(chan bool)

	go func() {
		for !ui.quizz.Terminated() {
			ui.ask()
		}

		done <- true
	}()

	return done
}

func (ui UI) ask() {
	reader := bufio.NewReader(os.Stdin)

	correct := false

	for !correct {
		fmt.Printf("%s:\n", ui.quizz.Current())

		read, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}

		if ui.quizz.Check(read) {
			fmt.Printf("Correct answer !\n\n")
			correct = true
		} else {
			fmt.Printf("Wrong ! Try again...\n\n")
		}
	}
}
