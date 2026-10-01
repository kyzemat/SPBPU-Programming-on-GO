package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
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

	name := readString("Введите имя студента: ")
	secondName := readString("Введите фамилию студента: ")
	lastName := readString("Введите отчество студента: ")
	age := readUint8("Введите возраст студента(14 - 120): ", 14, 120)
	avgScore := readFloat32("Введите средний балл студента (0 - 5): ", 0, 5)
	expelled := readBool("Укажите, был ли отчислен (true/false): ")

	students = append(students, Student{
		Name: name, SecondName: secondName, LastName: lastName,
		Age: age, AvgScore: avgScore, Expelled: expelled,
	})
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

		var choice uint8
		fmt.Print("Выберите действие: ")
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("❌ Введите число от 0 до 6")
			var dump string
			fmt.Scanln(&dump)
			continue
		}

		switch choice {
		case 1:
			AddStudentsManually()
		case 2:
			AddStudentRandom()
		case 3:
			var count uint8
			fmt.Print("Сколько студентов добавить? ")
			if _, err := fmt.Scan(&count); err != nil {
				fmt.Println("❌ Введите число от 0 до 255")
				var dump string
				fmt.Scanln(&dump)
				continue
			}
			for i := uint8(0); i < count; i++ {
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

func readUint8(prompt string, min, max uint8) uint8 {
	var v uint8
	for {
		fmt.Print(prompt)
		if _, err := fmt.Scan(&v); err != nil {
			fmt.Println("❌ Введите целое число")
			var dump string
			fmt.Scanln(&dump) // чистим буфер от мусора
			continue
		}
		if v < min || v > max {
			fmt.Printf("❌ Число должно быть от %d до %d\n", min, max)
			continue
		}
		return v
	}
}

func readFloat32(prompt string, min, max float32) float32 {
	var v float32
	for {
		fmt.Print(prompt)
		if _, err := fmt.Scan(&v); err != nil {
			fmt.Println("❌ Введите число, например 4.5")
			var dump string
			fmt.Scanln(&dump)
			continue
		}
		if v < min || v > max {
			fmt.Printf("❌ Число должно быть от %.2f до %.2f\n", min, max)
			continue
		}
		return v
	}
}

func readBool(prompt string) bool {
	for {
		fmt.Print(prompt)
		var s string
		fmt.Scan(&s)
		switch strings.ToLower(s) {
		case "true", "1", "да", "yes", "y":
			return true
		case "false", "0", "нет", "no", "n":
			return false
		default:
			fmt.Println("❌ Введите true или false")
		}
	}
}

func readString(prompt string) string {
	var s string
	for {
		fmt.Print(prompt)
		if _, err := fmt.Scan(&s); err == nil {
			return s
		}
		fmt.Println("❌ Поле не может быть пустым")
	}
}
