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

// -----------------------------------------------------------------------------------------------------

func SortStudentsByLessScore() {
	sort.Slice(students, func(i, j int) bool {
		return students[i].AvgScore < students[j].AvgScore
	})
}

func SortStudentsByMoreScore() {
	sort.Slice(students, func(i, j int) bool {
		return students[i].AvgScore > students[j].AvgScore
	})
}

func PrintStudents() {
	if len(students) == 0 {
		fmt.Println("❌ Список студентов пуст.")
		return
	}
	fmt.Println("\n========== СПИСОК СТУДЕНТОВ ==========")
	for i, student := range students {
		fmt.Printf(
			"\n№ %d\nИмя: %s %s %s\nВозраст: %d\nСредний балл: %.1f\nОтчислен: %t\n",
			i+1,
			student.Name,
			student.SecondName,
			student.LastName,
			student.Age,
			student.AvgScore,
			student.Expelled)
	}
	fmt.Println("\n=======================================")
}

func main() {
	for {
		fmt.Println("\n========== МЕНЮ ==========")
		fmt.Println("1. Добавить студента вручную")
		fmt.Println("2. Добавить случайного студента")
		fmt.Println("3. Добавить несколько случайных студентов")
		fmt.Println("4. Показать всех студентов")
		fmt.Println("5. Сортировать по возрастанию балла")
		fmt.Println("6. Сортировать по убыванию балла")
		fmt.Println("0. Выход")
		fmt.Println("===========================")

		var choice int
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			AddStudentsManually()
		case 2:
			AddStudentRandom()
		case 3:
			var count int
			fmt.Print("Сколько студентов добавить? ")
			fmt.Scan(&count)
			for i := 0; i < count; i++ {
				AddStudentRandom()
			}
		case 4:
			PrintStudents()
		case 5:
			SortStudentsByLessScore()
			fmt.Println("✅ Студенты отсортированы по возрастанию балла.")
		case 6:
			SortStudentsByMoreScore()
			fmt.Println("✅ Студенты отсортированы по убыванию балла.")
		case 0:
			fmt.Println("👋 Программа завершена.")
			return
		default:
			fmt.Println("❌ Такого пункта меню нет.")
		}
	}
}
