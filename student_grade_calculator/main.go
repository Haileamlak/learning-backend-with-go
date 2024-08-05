package student_grade_calculator

import (
	"fmt"
)

type Subject struct {
	name  string
	grade float64
}

type StudentGradeCalculator struct {
	studentName      string
	numberOfSubjects int
	subjects         []Subject
}

func main() {
	stuGradeCalc := StudentGradeCalculator{}

	fmt.Println(" --- Welcome to my non-AI-powered average grade calculator --- ")

	fmt.Print("Your Name Please: ")
	fmt.Scan(&stuGradeCalc.studentName)

	fmt.Print("Number of subjects you took: ")
	fmt.Scan(&stuGradeCalc.numberOfSubjects)

	for stuGradeCalc.numberOfSubjects <= 0 {
		fmt.Print("Please Enter a Valid Number of Subjects: ")
		fmt.Scan(&stuGradeCalc.numberOfSubjects)
	}

	for i := 0; i < stuGradeCalc.numberOfSubjects; i++ {
		subject := Subject{}

		fmt.Printf("Subject %v:\n", i+1)

		fmt.Print("\tEnter Name: ")
		fmt.Scan(&subject.name)

		fmt.Print("\tEnter Grade (0 - 100): ")
		fmt.Scan(&subject.grade)

		for !isValidGrade(subject.grade) {
			fmt.Print("\tPlease Enter a Valid Grade between 0 and 100: ")
			fmt.Scan(&subject.grade)
		}
		stuGradeCalc.subjects = append(stuGradeCalc.subjects, subject)
	}

	// clear the console
	fmt.Print("\033[H\033[2J")

	fmt.Println("\n --- Your Grades --- ")
	for _, subject := range stuGradeCalc.subjects {
		fmt.Println(subject.name, " -> ", grade_calculator(subject.grade))
	}

	fmt.Printf("\nYour average grade is %v.", grade_calculator(stuGradeCalc.calculateAverageGrade()))
}

func isValidGrade(grade float64) bool {
	return grade >= 0 && grade <= 100
}

func grade_calculator(score float64) string {
	if score >= 85 {
		return "A"
	}
	if score >= 80 {
		return "A-"
	}
	if score >= 75 {
		return "B+"
	}
	if score >= 68 {
		return "B"
	}
	if score >= 65 {
		return "B-"
	}
	if score >= 60 {
		return "C+"
	}
	if score >= 50 {
		return "C"
	}
	if score >= 45 {
		return "C-"
	}
	if score >= 40 {
		return "D"
	}
	return "F"
}

func (sgc StudentGradeCalculator) calculateAverageGrade() float64 {
	total := 0.0

	for _, subject := range sgc.subjects {
		total += subject.grade
	}

	averageGrade := total / float64(sgc.numberOfSubjects)
	return averageGrade
}
