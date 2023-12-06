package main

func totalMoney(n int) int {
	sum := 0
	weekDay := 0
	weekCounter := 1
	for day := 1; day <= n; day++ {
		dollars := weekCounter + weekDay
		if day%7 == 0 {
			weekCounter++
			weekDay = 0
			sum += dollars
			continue
		}
		weekDay++
		sum += dollars
	}
	return sum
}
