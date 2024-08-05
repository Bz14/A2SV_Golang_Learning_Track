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
func displayOutput(studentName string, subjects map[string]float64){
	fmt.Println("=============================================")
	fmt.Printf("Student Name: %s", studentName)
	fmt.Println("=============================================")
	fmt.Printf("%s\t\t%s\n", "Subjects", "Grades") 
	fmt.Println("=============================================")
	for key, value := range subjects{
		fmt.Printf("%s\t\t\t%.2f\n", key, value)
	}
	fmt.Println("=============================================")
	fmt.Printf("Average Grade: %.2f\n", calculateAverage(subjects))
	fmt.Println("=============================================")
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
	inValidName := true
	inValidSubject := true
	flag := true
	for flag{
		for inValidName{
			studentName, _ := getInput("Enter Your Name: ", reader)
			if strings.TrimSpace(studentName ) == ""{
				fmt.Println("Invalid input. Please enter again.")
				continue
			}
			inValidName = false
			for inValidSubject{
				numOfSubjects, _ := getInput("Enter the number of subjects: ", reader)
				subCount, err := strconv.ParseInt(strings.TrimSpace(numOfSubjects), 10, 64)
				if err != nil{
					fmt.Println("Invalid input. Please enter again.")
					continue
				}
				inValidSubject = false
				for subCount > 0{
					subject, _ := getInput("Enter the name of subject: ", reader)
					if strings.TrimSpace(subject) == ""{
						fmt.Println("Invalid input. Please enter again.")
						continue
					}
					_, isExists := subjects[subject]
					if isExists{
						fmt.Println("Subject already exists, enter another: ")
						continue
					}
					grade, _ := getInput("Enter your grade: ", reader)
					gradeInt, err := strconv.ParseFloat(strings.TrimSpace(grade), 64)
					if err != nil{
						fmt.Println("Invalid input. Please enter again.")
						continue
					}
					if (gradeInt < 0 || gradeInt > 100){
						fmt.Println("Invalid Grade Range, insert between (0-100).")
						continue
					}
					subjects[strings.TrimSpace(subject)] = gradeInt
					subCount--
				}
			}
			flag = false
			displayOutput(studentName, subjects)
		}
	}
}
