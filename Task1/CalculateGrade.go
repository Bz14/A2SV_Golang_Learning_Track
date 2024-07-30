package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Function to calculate average */
func calculateAverage(subjects map[string]float64)float64{
	n := float64(len(subjects))
	var total float64
	for _, grade := range subjects{
		total += grade
	}
	return total / n
}

/* Function to display formatted output */
func displayOutput(studentName string, subjects map[string]float64)string{
	st := "=============================================\n"
	st += fmt.Sprintf("\nStudent Name: %s\n", studentName)
	st += fmt.Sprint("Subjects:\n")
	for key, value := range subjects{
		st += fmt.Sprintf("%-6s %s = %.2f\n", " " , key, value)
	}
	st += fmt.Sprintf("\nAverage Grade: %.2f\n", calculateAverage(subjects))
	st += "=============================================\n"
	return st
}

/* Function to accept input from user */
func getInput(prompt string, r *bufio.Reader)(string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return input, err
}

func main() {
	subjects := map[string]float64{}
	reader := bufio.NewReader(os.Stdin)
	studentName, _ := getInput("Enter Your Name: ", reader)
	numOfSubjects, _ := getInput("Enter the number of subjects: ", reader)
	subCount, _ := strconv.ParseInt(strings.TrimSpace(numOfSubjects), 10, 64)
	for subCount > 0{
		subject, _ := getInput("Enter the name of subject: ", reader)
		_, isExists := subjects[subject]
		if isExists{
			fmt.Println("Subject already exists, enter another: ")
			continue
		}
		grade, _ := getInput("Enter your grade: ", reader)
		gradeInt, _ := strconv.ParseFloat(strings.TrimSpace(grade), 64)
		if (gradeInt < 0 || gradeInt > 100){
			fmt.Println("Invalid Grade Range, insert between (0-100).")
			continue
		}
		subjects[strings.TrimSpace(subject)] = gradeInt
		subCount--
	}
	fmt.Println(displayOutput(studentName, subjects))
}
