package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name       string
	SecondName string
	LastName   string
	Age        uint8
	AvgScore   float32
	Expelled   bool
}

var RandomNames = []string{"Алексей", "Петр", "Генри", "Ким", "Виталя"}
var RandomSecondNames = []string{"Кротов", "Опалюк", "Сердюков", "Шифонер", "Булков"}
var RandomLastNames = []string{"Евгеньевич", "Дмитриевич", "Александрович", "Иммануилович", "Саныч"}
var students []Student

func main() {
}

func AddStudentRandom() {
	name := RandomNames[rand.Intn(len(RandomNames))]
	secondName := RandomSecondNames[rand.Intn(len(RandomSecondNames))]
	lastName := RandomLastNames[rand.Intn(len(RandomLastNames))]

	age := uint8(rand.Intn(9) + 17)
	avgScore := float32(rand.Intn(21)+30) / 10
	expelled := rand.Intn(5) == 0

	student := Student{
		Name:       name,
		SecondName: secondName,
		LastName:   lastName,
		Age:        age,
		AvgScore:   avgScore,
		Expelled:   expelled,
	}

	students = append(students, student)

	fmt.Printf("✅ Студент %s %s добавлен!\n", name, secondName)
}
func AddStudentsManually() {
	var name string
	var secondName string
	var lastName string
	var age uint8
	var avgScore float32
	var expelled bool
	fmt.Print("Введите имя студента: ")
	fmt.Scan(&name)
	fmt.Print("Введите фамилию студента: ")
	fmt.Scan(&secondName)
	fmt.Print("Введите отчество студента: ")
	fmt.Scan(&lastName)
	fmt.Print("Введите возраст студента: ")
	fmt.Scan(&age)
	fmt.Print("Введите средний балл студента: ")
	fmt.Scan(&avgScore)
	fmt.Print("Укажите, был ли отчислен студент или нет(true или false): ")
	fmt.Scan(&expelled)
	student := Student{
		Name:       name,
		SecondName: secondName,
		LastName:   lastName,
		Age:        age,
		AvgScore:   avgScore,
		Expelled:   expelled,
	}
	students = append(students, student)
	fmt.Printf("✅ Студент %s %s добавлен!\n", name, secondName)
}
