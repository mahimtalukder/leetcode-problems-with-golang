package main

func romanToInt(s string) int {
    romainNumbersWithValue := map[rune]int{
		'I' : 1,
		'V' : 5,
		'X' : 10,
		'L' : 50,
		'C' : 100,
		'D' : 500,
		'M' : 1000,
	}
	charSlice := []rune(s)
	value := 0
	for i := 0; i < len(charSlice); i++{
		if i == len(charSlice) - 1{
			value += romainNumbersWithValue[charSlice[i]]
		}else if romainNumbersWithValue[charSlice[i]] == 1 && romainNumbersWithValue[charSlice[i+1]] == 5{
			value += 4
			i++
		} else if romainNumbersWithValue[charSlice[i]] == 1 && romainNumbersWithValue[charSlice[i+1]] == 10{
			value += 9
			i++
		} else if romainNumbersWithValue[charSlice[i]] == 10 && romainNumbersWithValue[charSlice[i+1]] == 50{
			value += 40
			i++
		} else if romainNumbersWithValue[charSlice[i]] == 10 && romainNumbersWithValue[charSlice[i+1]] == 100{
			value += 90
			i++
		} else if romainNumbersWithValue[charSlice[i]] == 100 && romainNumbersWithValue[charSlice[i+1]] == 500{
			value += 400
			i++
		} else if romainNumbersWithValue[charSlice[i]] == 100 && romainNumbersWithValue[charSlice[i+1]] == 1000{
			value += 900
			i++
		}else{
			value += romainNumbersWithValue[charSlice[i]]
		}
	}
	return value
}