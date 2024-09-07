package main

import "fmt"

type student struct {
	name   string
	gender int8
}

func searchStudent(name string, studentMap map[string]student) *student {
	for _, student := range studentMap {
		if student.name == name {
			return &student
		}
	}
	return nil
}

func addStudent(key string, name string, gender int8, studentMap map[string]student) {
	student := student{name: name, gender: gender}
	studentMap[key] = student
}

func updateStudent(key string, name string, gender int8, studentMap map[string]student) {
	student := student{name: name, gender: gender}
	studentMap[key] = student
}

func deleteStudent(key string, studentMap map[string]student) {
	delete(studentMap, key)
}

func genderClassification(studentMap map[string]student) {
	var males []student
	var females []student

	for _, student := range studentMap {
		if student.gender == 1 {
			males = append(males, student)
		} else if student.gender == 0 {
			females = append(females, student)
		}
	}

	fmt.Println("Male Students:")
	for _, male := range males {
		fmt.Println(male.name)
	}

	fmt.Println("\nFemale Students:")
	for _, female := range females {
		fmt.Println(female.name)
	}
}

func main() {
	var studentMap = map[string]student{
		"sv0": {
			name:   "Student1",
			gender: 1,
		},
		"sv2": {
			name:   "Student2",
			gender: 0,
		},
		"sv3": {
			name:   "Student3",
			gender: 1,
		},
		"sv4": {
			name:   "Student4",
			gender: 0,
		},
	}

	foundStudent := searchStudent("Student7", studentMap)
	if foundStudent != nil {
		fmt.Println(*foundStudent)
	} else {
		fmt.Println("Student not found")
	}

	addStudent("sv5", "Student5", 0, studentMap)
	updateStudent("sv5", "Student5", 1, studentMap)
	deleteStudent("sv5", studentMap)
	genderClassification(studentMap)
}
