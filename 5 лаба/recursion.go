package main

// ============================================================
// ЗАДАЧА 5: Merge Two Sorted Lists (Слияние двух отсортированных списков)
// ============================================================
// Нужно объединить два односвязных списка, отсортированных по
// неубыванию, в один отсортированный список. Новые узлы не создаём:
// перенастраиваем указатели Next у существующих узлов.

// ListNode — узел односвязного списка.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists сливает два отсортированных списка в один.
func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	// Фиктивный узел избавляет от особых случаев для головы результата.
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		// Берём меньший из двух текущих узлов.
		// <= сохраняет порядок равных элементов (устойчивое слияние).
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Один из списков закончился — присоединяем остаток другого целиком.
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

// NewList строит список из переданных значений (вспомогательная функция для демо).
func NewList(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// ListToSlice превращает список в срез (для вывода).
func ListToSlice(head *ListNode) []int {
	result := []int{}
	for ; head != nil; head = head.Next {
		result = append(result, head.Val)
	}
	return result
}
