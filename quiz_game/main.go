package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// Holds the questions/answers from read CSV file
var questions [][]string
// Quiz time limit
var timeLimit int

func main() {
	readCSV()
	initializeQuiz()
	startQuiz(timeLimit)
}

func readCSV() {
	file, fileErr := os.ReadFile("./problems.csv")
	if fileErr != nil {
		fmt.Printf("Error while reading the csv file: %s\n", fileErr)
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(string(file)))
	var readErr error
	questions, readErr = r.ReadAll()

	if readErr != nil {
		fmt.Printf("Error while reading the csv file: %s\n", readErr)
		os.Exit(1)
	}
}

func initializeQuiz() {
	fmt.Printf("\n\nWelcome to math quiz!\n\n")
	fmt.Printf("Answer to questions one by one by typing the answer and then press 'enter'-key\n\n")
	fmt.Printf("\nAre you ready? [Y/N]: ")

	var playerReady string
	fmt.Scanln(&playerReady)

	if playerReady != "Y" {
		fmt.Println("Try quiz again next time!")
		os.Exit(0)
	}

	fmt.Printf("\nPlease provide timelimit for quiz (seconds): ")
	fmt.Scanln(&timeLimit)
}

func promptRetry() {
	var retry string
	fmt.Printf("\nDo you want to try again? [Y/N]  \n")
	fmt.Scanln(&retry)
	if retry == "Y" {
		startQuiz(timeLimit)
	} else {
		fmt.Println("\n\nThank you for trying out the quiz! See you again!")
		os.Exit(0)
	}
}

func startQuiz(timeLimit int) {
	var answers = make(map[string]string)
	var timeRunout bool

	// Fairly simple implementation of timer - set end time for quiz
	endTime := time.Now().Add(time.Second * time.Duration(timeLimit)).Unix()

	for i := 0; i < len(questions); i++ {
		// Compare current time to endTime - stop quiz if time limit reached
		// Not strict implementation as one could possibly answer to question
		// even if time has passed over timelimit until next question is asked
		if time.Now().Unix() > endTime {
			timeRunout = true
			break
		}

		var question = questions[i][0]
		var answer string
		fmt.Printf("\nQuestion %d:\n%v=  ", i+1, question)
		fmt.Scanln(&answer)

		// Save to collection without possible extra empty spaces
		answers[question] = strings.ReplaceAll(answer, " ", "")
	}

	if timeRunout {
		fmt.Printf("\n\nEND OF QUIZ!\nUNFORTUNATELY TIME LIMIT WAS REACHED\nHERE ARE THE RESULTS: \n\n")
	} else {
		fmt.Printf("\n\nEND OF QUIZ!\nHERE ARE THE RESULTS: \n\n")
	}

	quizResults(answers)
	promptRetry()
}

func quizResults(answers map[string]string) {
	var correctAnswers int

	for i := 0; i < len(questions); i++ {
		question := questions[i][0]
		// Remove possible empty spaces
		answer := strings.ReplaceAll(answers[question], " ", "")
		correctAnswer := strings.ReplaceAll(questions[i][1], " ", "")

		if answer == correctAnswer {
			correctAnswers++
		}
	}

	totalCorrect := fmt.Sprintf(string("\033[32m")+"%d"+string("\033[37m")+"/%d", correctAnswers, len(questions))

	fmt.Printf("YOUR RESULT: %s\n\nRESULTS QUESTION BY QUESTION:\n", totalCorrect)

	for i := 0; i < len(questions); i++ {
		question := questions[i][0]
		answer := answers[question]
		correctAnswer := questions[i][1]
		var result string
		if answer == correctAnswer {
			result = string("\033[32m") + "CORRECT" + string("\033[37m")
			answer = string("\033[32m") + answer + string("\033[37m")
			correctAnswer = string("\033[32m") + correctAnswer + string("\033[37m")
		} else {
			result = string("\033[31m") + "INCORRECT" + string("\033[37m")
			answer = string("\033[31m") + answer + string("\033[37m")
			correctAnswer = string("\033[32m") + correctAnswer + string("\033[37m")
		}

		time.Sleep(time.Second * 1)
		fmt.Printf("\nQUESTION %d :: %v\nQuestion: %s=?\nYour answer: %s\nCorrect answer: %s\n\n", i+1, result, question, answer, correctAnswer)
	}
}
