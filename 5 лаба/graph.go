package main

// ============================================================
// ЗАДАЧА 1: Course Schedule (Расписание курсов)
// ============================================================
// Проверяем, можно ли пройти все курсы, учитывая зависимости.
// Это проверка графа на наличие циклов (топологическая сортировка).
// Решение: алгоритм Кана (BFS по уровням входящих степеней).

func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Строим граф: для каждого курса список тех, кто от него зависит
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses) // Входящая степень (сколько предков)

	for _, pre := range prerequisites {
		ai, bi := pre[0], pre[1]
		// Чтобы пройти ai, нужно сначала bi. То есть ребро bi -> ai
		graph[bi] = append(graph[bi], ai)
		inDegree[ai]++
	}

	// В очередь кладем все курсы, которые можно пройти сразу (без предков)
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	visitedCount := 0

	for len(queue) > 0 {
		// Берем курс из начала очереди
		current := queue[0]
		queue = queue[1:]
		visitedCount++

		// Убираем зависимость от текущего курса у всех его "потомков"
		for _, next := range graph[current] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// Если прошли все курсы — цикла нет
	return visitedCount == numCourses
}

// ============================================================
// ЗАДАЧА 2: Network Delay Time (Задержка в сети)
// ============================================================
// Находим минимальное время, за которое сигнал из узла k дойдет до всех.
// Это классический алгоритм Дейкстры на ориентированном графе.

func NetworkDelayTime(times [][]int, n int, k int) int {
	// Строим граф смежности: graph[u] = список пар {v, w}
	type edge struct {
		to, weight int
	}
	graph := make([][]edge, n+1) // узлы от 1 до n

	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		graph[u] = append(graph[u], edge{v, w})
	}

	// dist[i] — минимальное время до узла i. Изначально бесконечность.
	const INF = 1 << 30
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = INF
	}
	dist[k] = 0

	// Приоритетная очередь (реализуем через срез, сортируя вручную для простоты)
	// В Go стандартно используют container/heap, но для простоты используем срез
	type item struct {
		node, time int
	}
	pq := []item{{k, 0}}

	for len(pq) > 0 {
		// Извлекаем элемент с минимальным временем (простая реализация)
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i].time < pq[minIdx].time {
				minIdx = i
			}
		}
		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)

		u := current.node
		if current.time > dist[u] {
			continue // Устаревшая запись
		}

		// Расслабляем ребра
		for _, e := range graph[u] {
			newTime := dist[u] + e.weight
			if newTime < dist[e.to] {
				dist[e.to] = newTime
				pq = append(pq, item{e.to, newTime})
			}
		}
	}

	// Находим максимальное время среди всех узлов
	maxTime := 0
	for i := 1; i <= n; i++ {
		if dist[i] == INF {
			return -1 // До какого-то узла дойти невозможно
		}
		if dist[i] > maxTime {
			maxTime = dist[i]
		}
	}

	return maxTime
}
