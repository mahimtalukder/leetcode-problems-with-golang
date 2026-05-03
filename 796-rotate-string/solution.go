package main

func rotateString(s string, goal string) bool {
    n := len(s)
	if n != len(goal){
		return false
	}

	for range n{
		s = s[1:n] + s[0:1]
		if s == goal {
			return true
		}
	}
	return false
}