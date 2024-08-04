package main

import (
	"fmt"
)

func main() {
	stu_grade_calc := StudentGradeCalculator{}
	fmt.Println("Welcome to my non-AI-powered grade calculator")
	fmt.Print("Your Name Please: ")
	fmt.Scan(&stu_grade_calc.studentName)
	fmt.Print("Number of subjects you took: ")
	fmt.Scan(&stu_grade_calc.numberOfSubjects)
	for i := 0; i < stu_grade_calc.numberOfSubjects; i++ {
		subject := Subject{}
		fmt.Printf("Subject %v\n", i + 1)

		fmt.Print("\tEnter Name: ")
		fmt.Scan(&subject.subject_name)
		
		fmt.Print("\tEnter Grade (0 - 100): ")
		fmt.Scan(&subject.grade)
		for !valid(subject.grade) {
			fmt.Print("\tPlease Enter a Valid Grade between 0 and 100: ")
			fmt.Scan(&subject.grade)
		}
		stu_grade_calc.subjects = append(stu_grade_calc.subjects, subject)
	}
	fmt.Printf("Your average grade is %2f.", stu_grade_calc.average_grade())
}

func valid(grade float64) bool {
	return grade >= 0 && grade <= 100
}

type Subject struct {
	subject_name string
	grade        float64
}

type StudentGradeCalculator struct {
	studentName      string
	numberOfSubjects int
	subjects         []Subject
}

func (SGC StudentGradeCalculator) average_grade() float64 {
	total := 0.0
	for _, subject := range SGC.subjects {
		total += subject.grade
	}

	average_grade := total / float64(SGC.numberOfSubjects)
	return average_grade
}
