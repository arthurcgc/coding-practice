package main

import "sort"

func kidsWithCandies(candies []int, extraCandies int) []bool {
    sorted := make([]int, len(candies))
    copy(sorted, candies)
    sort.Ints(sorted)
    highest := sorted[len(sorted)-1]
    result := make([]bool, len(candies))
    for i := range candies {
        sum := candies[i] + extraCandies
        if sum >= highest {
            result[i] = true
        }
    }
    return result
}