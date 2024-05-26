package main

// Part 1 https://courses.calhoun.io/lessons/les_goph_01
// User must be able to specify file path using a flag in the terminal command

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var filePath string

// Is run before Main
func init() {
	flag.StringVar(&filePath, "f", "", "Path to the CSV file.")
}

func main() {
	flag.Parse()
	fmt.Println("Welcome to the quiz!!\n===================\n")

	if filePath == "" {
		filePath = "./problems.csv"
	}
	var questionsAndAnswers []QuizQuestion = readCSV(filePath)

	// Start Quiz
	var userScore int
	startSequence()

	for i := range questionsAndAnswers {
		questionData := questionsAndAnswers[i]

		var answer string
		fmt.Printf("%s = ", questionData.question)
		fmt.Scanln(&answer)

		if answer == questionData.answer {
			userScore++
		}
	}

	fmt.Printf("FINAL SCORE IS: %d out of %d!", userScore, len(questionsAndAnswers))
}

type QuizQuestion struct {
	question string
	answer   string
}

func readCSV(filePath string) []QuizQuestion {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open CSV file.")
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read lines from CSV file.")
	}

	//operators := []string{"+", "-", "*", "/"}
	var questionSlice []QuizQuestion

	// Get both parts of the CSV line, now separated by a space
	for _, line := range lines {
		var questionData QuizQuestion
		questionData.question = line[0]
		questionData.answer = line[1]

		questionSlice = append(questionSlice, questionData)

		// Separate the members and find the operator
		//for _, operator := range operators {
		// 	splitMathStr := strings.Split(line[0], operator)

		// 	if len(splitMathStr) == 2 {
		// 		firstNum, err1 := strconv.Atoi(splitMathStr[0])
		// 		operator := operator
		// 		secondNum, err2 := strconv.Atoi(splitMathStr[1])
		// 		result, err3 := strconv.Atoi(line[1])
		// 		if err1 != nil || err2 != nil || err3 != nil {
		// 			continue
		// 		}

		// 		question = QuizQuestion{
		// 			firstNum:  firstNum,
		// 			operator:  operator,
		// 			secondNum: secondNum,
		// 			result:    result,
		// 		}
		// 	}
		// }
		// questionSlice = append(questionSlice, question)
	}
	return questionSlice
}

func startSequence() {
	fmt.Println("QUIZ STARTING IN")
	fmt.Println("3...")
	time.Sleep(time.Second)
	fmt.Println("2..")
	time.Sleep(time.Second)
	fmt.Println("1.")
	time.Sleep(time.Second)
	fmt.Println("GO!!\n")
}
