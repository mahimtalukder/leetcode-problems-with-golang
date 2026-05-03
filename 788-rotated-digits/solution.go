package main

func rotatedDigits(n int) int {
	totalGoodNumber := 0
	for i := 1; i <= n; i++ {
		num := i
		isValid := false
		isGood := false

		for num > 0 {
			digit := num % 10
			switch digit {
			case 2, 5, 6, 9:
				isValid = true
				isGood = true
			case 0, 1, 8:
				isValid = true
			default:
				isValid = false
			}
			
			if !isValid{
				break
			}

			num /= 10
		}

		if isValid && isGood{
			totalGoodNumber++
		}
	}
	return totalGoodNumber
}
