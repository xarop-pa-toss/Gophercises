package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the quiz!!")
}

type QuizQuestion struct {
	firstNum  int
	operator  string
	secondNum int
	result    int
}

func readCSV(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open CSV file.")
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read lines from CSV file.")
	}

	return lines
}

func mapToQuizQuestionStruct(lines [][]string) []QuizQuestion {
	operators := []string{"+", "-", "*", "/"}
	var questionSlice []QuizQuestion

	// Get both parts of the CSV line, now separated by a space
	for _, line := range lines {
		var question QuizQuestion

		// Separate the members and find the operator
		for _, operator := range operators {
			splitMathStr := strings.Split(line[0], operator)

			if len(splitMathStr) == 2 {
				firstNum, err1 := strconv.Atoi(splitMathStr[0])
				operator := operator
				secondNum, err2 := strconv.Atoi(splitMathStr[1])
				result, err3 := strconv.Atoi(line[1])
				if err1 != nil || err2 != nil || err3 != nil {
					continue
				}

				question = QuizQuestion{
					firstNum:  firstNum,
					operator:  operator,
					secondNum: secondNum,
					result:    result,
				}
			}
		}
		questionSlice = append(questionSlice, question)
	}
	return questionSlice
}
