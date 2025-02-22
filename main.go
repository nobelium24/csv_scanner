package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Student struct {
	MatricNum     string `csv:"MATRIC_NUM"`
	JambNum       string `csv:"JAMB_NUM"`
	Surname       string `csv:"SURNAME"`
	FirstName     string `csv:"FIRST_NAME"`
	MiddleName    string `csv:"MIDDLE_NAME"`
	State         string `csv:"STATE"`
	LGA           string `csv:"LGA"`
	EntryMode     string `csv:"ENTRY_MODE"`
	Sex           string `csv:"SEX"`
	DOB           string `csv:"DOB"`
	Email         string `csv:"EMAIL"`
	MaritalStatus string `csv:"MARITAL_STATUS"`
	NextOfKin     string `csv:"NEXT_OF_KIN"`
	NextOfKinAddr string `csv:"NEXT_OF_KIN_ADDRESS"`
}

func readFile(filename string) ([]Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("no data found")
	}

	var students []Student
	for _, record := range records[1:] {
		student := Student{
			MatricNum:     record[0],
			JambNum:       record[1],
			Surname:       record[2],
			FirstName:     record[3],
			MiddleName:    record[4],
			State:         record[5],
			LGA:           record[6],
			EntryMode:     record[7],
			Sex:           record[8],
			DOB:           record[9],
			Email:         record[10],
			MaritalStatus: record[11],
			NextOfKin:     record[12],
			NextOfKinAddr: record[13],
		}
		students = append(students, student)
	}
	return students, nil
}

func sortStudents(students []Student) map[string][]Student {
	stateMap := make(map[string][]Student)

	for _, student := range students {
		stateMap[student.State] = append(stateMap[student.State], student)
	}

	for state, students := range stateMap {
		sort.Slice(students, func(i, j int) bool {
			return students[i].MatricNum < students[j].MatricNum
		})
		stateMap[state] = students
	}

	return stateMap
}

func writeFile(statesMap map[string][]Student) error {
	var failedStates []string
	invalidFilenameChars := regexp.MustCompile(`[<>:"/\\|?*]`)

	for state, students := range statesMap {
		if strings.TrimSpace(state) == "" { // Skip empty state values
			fmt.Println("Skipping entry with empty state name")
			continue
		}

		// Replace invalid filename characters
		safeState := invalidFilenameChars.ReplaceAllString(state, "_")

		fileName := fmt.Sprintf("%s.csv", safeState)
		newFile, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", fileName, err)
			failedStates = append(failedStates, state)
			continue // Skip this state but continue with others
		}
		defer newFile.Close()

		fileWriter := csv.NewWriter(newFile)
		header := []string{"MATRIC_NUM", "JAMB_NUM", "SURNAME", "FIRST_NAME", "MIDDLE_NAME",
			"STATE", "LGA", "ENTRY_MODE", "SEX", "DOB", "EMAIL", "MARITAL_STATUS", "NEXT_OF_KIN", "NEXT_OF_KIN_ADDRESS"}
		fileWriter.Write(header)

		for _, student := range students {
			record := []string{
				student.MatricNum, student.JambNum, student.Surname, student.FirstName, student.MiddleName, student.State, student.LGA, student.EntryMode,
				student.Sex, student.DOB, student.Email, student.MaritalStatus, student.NextOfKin, student.NextOfKinAddr,
			}
			fileWriter.Write(record)
		}
		fileWriter.Flush()
		fmt.Printf("Created file: %s\n", fileName)
	}

	if len(failedStates) > 0 {
		return fmt.Errorf("failed to create files for states: %v", failedStates)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <csv_file>")
		return
	}

	csvFile := os.Args[1]
	students, err := readFile(csvFile)
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return
	}

	stateMap := sortStudents(students)

	err = writeFile(stateMap)
	if err != nil {
		fmt.Printf("Error writing CSV files: %v\n", err)
		return
	}

	fmt.Println("Processing complete.")
}
