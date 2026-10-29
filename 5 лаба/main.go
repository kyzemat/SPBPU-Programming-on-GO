package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n========== МЕНЮ: АЛГОРИТМЫ ==========")
		fmt.Println("1 — Course Schedule (Расписание курсов)")
		fmt.Println("2 — Network Delay Time (Задержка в сети)")
		fmt.Println("3 — Binary Tree Level Order Traversal (Обход по уровням)")
		fmt.Println("4 — Populating Next Right Pointers (Связывание указателей)")
		fmt.Println("5 — Merge Two Sorted Lists (Слияние отсортированных списков)")
		fmt.Println("6 — Same Tree (Одинаковые деревья)")
		fmt.Println("0 — Выход")
		fmt.Print("Ваш выбор: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("❌ Ошибка: введите число (0-4).")
			continue
		}

		switch choice {
		case 1:
			demoCourseSchedule()
		case 2:
			demoNetworkDelay()
		case 3:
			demoLevelOrder()
		case 4:
			demoConnect()
		case 5:
			demoMergeLists()
		case 6:
			demoSameTree()
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("❌ Нет такого пункта. Введите 0-4.")
		}
	}
}

// ============================================================
// ДЕМОНСТРАЦИИ
// ============================================================

func demoCourseSchedule() {
	fmt.Println("\n--- Задача 1: Course Schedule ---")
	fmt.Println("Вход: numCourses=2, prerequisites=[[1,0]]")
	fmt.Println("Выход:", CanFinish(2, [][]int{{1, 0}}))

	fmt.Println("\nВход: numCourses=2, prerequisites=[[1,0],[0,1]]")
	fmt.Println("Выход:", CanFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func demoNetworkDelay() {
	fmt.Println("\n--- Задача 2: Network Delay Time ---")
	fmt.Println("Вход: times=[[2,1,1],[2,3,1],[3,4,1]], n=4, k=2")
	fmt.Println("Выход:", NetworkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))

	fmt.Println("\nВход: times=[[1,2,1]], n=2, k=1")
	fmt.Println("Выход:", NetworkDelayTime([][]int{{1, 2, 1}}, 2, 1))

	fmt.Println("\nВход: times=[[1,2,1]], n=2, k=2")
	fmt.Println("Выход:", NetworkDelayTime([][]int{{1, 2, 1}}, 2, 2))
}

func demoLevelOrder() {
	fmt.Println("\n--- Задача 3: Binary Tree Level Order Traversal ---")
	// Строим дерево [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Дерево: [3,9,20,null,null,15,7]")
	fmt.Println("Выход:", LevelOrder(root))

	// Пустое дерево
	fmt.Println("\nПустое дерево:")
	fmt.Println("Выход:", LevelOrder(nil))
}

func demoConnect() {
	fmt.Println("\n--- Задача 4: Populating Next Right Pointers ---")
	// Строим идеальное дерево [1,2,3,4,5,6,7]
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}

	Connect(root)

	// Выводим результат в формате: 1->nil, 2->3->nil, ...
	fmt.Println("Дерево: [1,2,3,4,5,6,7]")
	fmt.Println("Результат связывания Next-указателей:")

	leftmost := root
	for leftmost != nil {
		cur := leftmost
		for cur != nil {
			if cur.Next != nil {
				fmt.Printf("%d -> ", cur.Val)
			} else {
				fmt.Printf("%d -> nil\n", cur.Val)
			}
			cur = cur.Next
		}
		leftmost = leftmost.Left
	}
}

func demoMergeLists() {
	fmt.Println("\n--- Задача 5: Merge Two Sorted Lists ---")

	fmt.Println("Вход: list1=[1,2,4], list2=[1,3,4]")
	fmt.Println("Выход:", ListToSlice(MergeTwoLists(NewList(1, 2, 4), NewList(1, 3, 4))))

	fmt.Println("\nВход: list1=[], list2=[]")
	fmt.Println("Выход:", ListToSlice(MergeTwoLists(NewList(), NewList())))

	fmt.Println("\nВход: list1=[], list2=[0]")
	fmt.Println("Выход:", ListToSlice(MergeTwoLists(NewList(), NewList(0))))
}

func demoSameTree() {
	fmt.Println("\n--- Задача 6: Same Tree ---")

	// [1,2,3] и [1,2,3]
	a := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	b := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println("Вход: p=[1,2,3], q=[1,2,3]")
	fmt.Println("Выход:", IsSameTree(a, b))

	// [1,2] и [1,null,2] — разная форма
	c := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	d := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println("\nВход: p=[1,2], q=[1,null,2]")
	fmt.Println("Выход:", IsSameTree(c, d))

	// [1,2,1] и [1,1,2] — разные значения
	e := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	f := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	fmt.Println("\nВход: p=[1,2,1], q=[1,1,2]")
	fmt.Println("Выход:", IsSameTree(e, f))
}
