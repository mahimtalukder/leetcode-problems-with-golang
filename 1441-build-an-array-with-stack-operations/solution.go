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

func buildArray(target []int, n int) []string {
	ans := []string{}
	const (
		push = "Push"
		pop  = "Pop"
	)
	var stack Stack
	steamCurrent := 1
	i := 0
	lenTerget := len(target)
	for steamCurrent <= n && i < lenTerget  {
		if len(stack) == 0 {
			stack.Push(steamCurrent)
			steamCurrent++
			ans = append(ans, push)
		} else if i<lenTerget && steamCurrent <= n && stack.Top() == target[i]{
			stack.Push(steamCurrent)
			steamCurrent++
			i++
			ans = append(ans, push)
		}else if stack.Top() != target[i]{
			stack.Pop()
			if i != 0{
				i--
			}
			ans = append(ans, pop)
		}

		if stack.Top() == target[lenTerget-1]{
			break
		}

	}

	return ans
}