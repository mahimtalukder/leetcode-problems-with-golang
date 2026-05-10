package main

const MAXN = 1000005

// isComposite[x] == false means x is prime.
// isComposite[x] == true means x is not prime.
var isComposite [MAXN]bool

func init() {
	isComposite[0] = true
	isComposite[1] = true

	for i := 2; i*i < MAXN; i++ {
		if !isComposite[i] {
			for j := i * i; j < MAXN; j += i {
				isComposite[j] = true
			}
		}
	}
}

func minJumps(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	// Find maximum value from nums.
	// We only need arrays up to maxVal.
	maxVal := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	// head[value] stores the first linked-list node for that value.
	//
	// Important trick:
	// We store index + 1 instead of index.
	// So:
	// 0 means empty
	// actual index = storedValue - 1
	head := make([]int, maxVal+1)

	// next[index] stores the next linked-list node.
	// It also stores index + 1.
	next := make([]int, n)

	// Build value -> indexes linked list.
	for i, value := range nums {
		next[i] = head[value]
		head[value] = i + 1
	}

	// visited[index] tells whether this index is already reached by BFS.
	visited := make([]bool, n)
	visited[0] = true

	// Fixed-size queue.
	// At most n indexes can be pushed.
	queue := make([]int, n)
	front, back := 0, 0

	queue[back] = 0
	back++

	// seenPrime[p] means teleportation for prime p was already processed.
	seenPrime := make([]bool, maxVal+1)

	jumps := 0

	for front < back {
		// Process one BFS level at a time.
		// All nodes in this level have the same jump count.
		levelSize := back - front

		for levelSize > 0 {
			levelSize--

			current := queue[front]
			front++

			if current == n-1 {
				return jumps
			}

			// -----------------------------
			// Move 1: Adjacent right
			// -----------------------------
			right := current + 1
			if right < n && !visited[right] {
				if right == n-1 {
					return jumps + 1
				}

				visited[right] = true
				queue[back] = right
				back++
			}

			// -----------------------------
			// Move 2: Adjacent left
			// -----------------------------
			left := current - 1
			if left >= 0 && !visited[left] {
				if left == n-1 {
					return jumps + 1
				}

				visited[left] = true
				queue[back] = left
				back++
			}

			// -----------------------------
			// Move 3: Prime teleportation
			// -----------------------------
			value := nums[current]

			// Teleport only if nums[current] is prime
			// and we have not processed this prime before.
			if !isComposite[value] && !seenPrime[value] {
				seenPrime[value] = true

				// Check every multiple of this prime value.
				for multiple := value; multiple <= maxVal; multiple += value {
					// Visit all indexes where nums[index] == multiple.
					for node := head[multiple]; node != 0; node = next[node-1] {
						idx := node - 1

						if !visited[idx] {
							if idx == n-1 {
								return jumps + 1
							}

							visited[idx] = true
							queue[back] = idx
							back++
						}
					}

					// Clear this value bucket after processing.
					// This prevents repeated scanning later.
					head[multiple] = 0
				}
			}
		}

		// After finishing this BFS level, jump count increases by 1.
		jumps++
	}

	return -1
}