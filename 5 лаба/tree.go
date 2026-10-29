package main

// ============================================================
// Определения структур деревьев
// ============================================================

// TreeNode — определение для обычного бинарного дерева (Задача 3)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node — определение для идеального дерева с указателем next (Задача 4)
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ============================================================
// ЗАДАЧА 3: Binary Tree Level Order Traversal
// ============================================================
// Обход дерева по уровням (BFS). Возвращает [][]int, где каждый
// подмассив — значения узлов на одном уровне.

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		// Обрабатываем все узлы текущего уровня
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// ============================================================
// ЗАДАЧА 4: Populating Next Right Pointers in Each Node
// ============================================================
// Для идеального бинарного дерева связываем узлы одного уровня
// через указатель Next. Используем уже установленные Next-указатели
// на предыдущем уровне для перехода к следующему (O(1) памяти).

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// leftmost — самый левый узел текущего уровня
	leftmost := root

	for leftmost.Left != nil {
		// cur — текущий узел на уровне
		cur := leftmost

		for cur != nil {
			// 1. Левый ребенок указывает на правого
			cur.Left.Next = cur.Right

			// 2. Правый ребенок указывает на левого ребенка следующего узла
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}

			// Переходим к следующему узлу на том же уровне
			cur = cur.Next
		}

		// Спускаемся на уровень ниже
		leftmost = leftmost.Left
	}

	return root
}
