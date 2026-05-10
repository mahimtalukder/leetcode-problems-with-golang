package main

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 1. Pre-allocate slice with the exact size needed
	dp := make([]int, n)
	
	// 2. Bound checking elimination hint for Go compiler
	// This tells the compiler that i and j will never exceed bounds,
	// removing safety branches from the compiled assembly loop.
	_ = nums[n-1]
	_ = dp[n-1]

	// 3. Unroll initialization loop locally
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < n; i++ {
		maxJumps := -1
		valI := nums[i]

		// 4. Tight inner loop using direct register values
		for j := 0; j < i; j++ {
			prevDP := dp[j]
			if prevDP != -1 {
				diff := valI - nums[j]
				if diff >= -target && diff <= target {
					if prevDP > maxJumps {
						maxJumps = prevDP
					}
				}
			}
		}

		if maxJumps != -1 {
			dp[i] = maxJumps + 1
		}
	}

	return dp[n-1]
}
