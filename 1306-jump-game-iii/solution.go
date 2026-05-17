package main

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return -1
	}

	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func canReach(arr []int, start int) bool {
	n := len(arr)
	found := false
	graph := make(map[int][]int)

	i := start
	leftIdx := i + arr[i]
	rightIdx := i - arr[i]

	edges := []int{}
	if leftIdx >= 0 && leftIdx < n && arr[leftIdx] == 0 {
		return true
	} else if leftIdx >= 0 && leftIdx < n {
		edges = append(edges, leftIdx)
	}

	if rightIdx >= 0 && rightIdx < n && arr[rightIdx] == 0 {
		return true
	} else if rightIdx >= 0 && rightIdx < n {
		edges = append(edges, rightIdx)
	}

	graph[i] = edges
	var stack Stack
	stack = append(stack, edges...)

	for !found && len(stack) != 0 {
		j, _ := stack.Pop()
		_, ok := graph[j]
		if !ok {
			i = j
			leftIdx = i + arr[i]
			rightIdx = i - arr[i]

			edges := []int{}
			if leftIdx >= 0 && leftIdx < n && arr[leftIdx] == 0 {
				return true
			} else if leftIdx >= 0 && leftIdx < n {
				edges = append(edges, leftIdx)
				if _, ok := graph[leftIdx]; !ok{
					stack.Push(leftIdx)
				}
			}

			if rightIdx >= 0 && rightIdx < n && arr[rightIdx] == 0 {
				return true
			} else if rightIdx >= 0 && rightIdx < n {
				edges = append(edges, rightIdx)
				if _, ok := graph[rightIdx]; !ok{
					stack.Push(rightIdx)
				}
			}
			graph[i] = edges
		}
	}
	return false
}
