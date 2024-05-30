package main

// https://courses.calhoun.io/lessons/les_goph_01
// Part 1 - User must be able to specify file path using a flag in the terminal command
// Part 2 - A timer must run from the start and stop after X amount of time.

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var flagFilePath string
var flagTimer uint
var flagRandomize bool

// Is run before Main
func init() {
	flag.StringVar(&flagFilePath, "f", "", "Path to the CSV file.")
	flag.UintVar(&flagTimer, "t", 30, "Time in seconds for quiz.")
	flag.BoolVar(&flagRandomize, "r", false, "Randomize the question order.")
}

func main() {

	flag.Parse()
	handleFlagFilePath()
	handleFlagTimer()

	var questionsAndAnswers []QuizQuestion = readCSV(flagFilePath)
	questionIndexOrder := handleFlagRandomize(questionsAndAnswers)

	fmt.Println("Welcome to the quiz!!\n===================\n")
	// Start Quiz
	var userScore int
	printStartSequence()

	for i := range questionsAndAnswers {
		questionData := questionsAndAnswers[questionIndexOrder[i]]

		var answer string
		fmt.Printf("%s = ", questionData.question)
		fmt.Scanln(&answer)

		// turn all to lower case and trim whitespace
		answer = strings.ToLower((answer))
		answer = strings.Trim(answer, " ")

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
		questionData.answer = strings.ToLower(strings.Trim(line[1], " "))

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

func handleFlagFilePath() {
	if flagFilePath == "" {
		flagFilePath = "./problems.csv"
	}
}

func handleFlagTimer() {
	if flagTimer < 0 {
		fmt.Println("Timer cannot be a negative number!")
		flagTimer = 30
		os.Exit(1)
	}
}

func handleFlagRandomize(questions []QuizQuestion) []int {
	indexes := make([]int, len(questions))
	for i := range indexes {
		indexes[i] = i
	}

	if flagRandomize {
		// Shuffle
		rand.Shuffle(len(indexes), func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i] })
	}

	return indexes
}

func printStartSequence() {
	fmt.Println("QUIZ STARTING IN")
	fmt.Println("3...")
	time.Sleep(time.Second)
	fmt.Println("2..")
	time.Sleep(time.Second)
	fmt.Println("1.")
	time.Sleep(time.Second)
	fmt.Println("GO!!\n")
}
