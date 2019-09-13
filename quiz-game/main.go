package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var quizBook []Quiz
var t1 time.Timer
var right int
var wrong int
var done = make(chan bool)

//Quiz holds the question and answer for the quiz game
type Quiz struct {
	question string
	answer   string
}

func main() {
	readtoQuizBook("problem.csv")
	fmt.Printf("All Question printed %v\n", quizBook)
	fmt.Println("Welcome to the math quiz ..Press Enter to Start..")

	t1 := time.NewTimer(time.Minute * 10)
	go startTest()
	select {
	case <-done:
	case <-t1.C:
		fmt.Println("Time's Up!")
	}

	fmt.Println("Total Question Answered : ", right+wrong)
	fmt.Println("Total Right Answer : ", right)
	fmt.Println("Total Wrong Answer : ", wrong)
}

//read the quiz from cv to an array
func readtoQuizBook(filename string) {
	csvfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	quizBook = make([]Quiz, 0)
	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		q := Quiz{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
		quizBook = append(quizBook, q)
	}
}

func loadNextQuestion(next int) {
	qz := quizBook[next]
	fmt.Printf("What is %v :  ", qz.question)
}

func startTest() {
	l := len(quizBook)
	fmt.Printf("Number of quiz :  %v\n", l)
	for i := 0; i < l; i++ {

		fmt.Printf("Next counter from loop %v\n", i)
		fmt.Printf("Scores for right answer : %v\n", right)
		fmt.Printf("Scores for wrong answer : %v\n", wrong)

		total := right + wrong
		result := checkEndofTest(total, l)

		if !result {
			for {
				loadNextQuestion(i)
				reader := bufio.NewReader(os.Stdin)
				input, _ := reader.ReadString('\n')

				input = input[:len(input)-2]

				if input != "" {
					fmt.Printf("Current counter %v\n", i)
					ans := quizBook[i].answer
					if ans != input {
						wrong++
					} else {
						right++
					}
					break
				}
			}
		}
	}
	done <- true
}
func checkEndofTest(a, b int) bool {
	if a <= b {
		return false
	}
	return true
}
func play() {
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
	fmt.Println("Playing")
}
