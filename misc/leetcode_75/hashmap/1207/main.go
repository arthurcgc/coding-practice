package main

func uniqueOccurrences(arr []int) bool {
    // occurrences : values
    countNums := map[int]int{}
    for _, v := range arr {
        if _, ok := countNums[v]; !ok {
            countNums[v] = 1
        } else {
            countNums[v]++
        }
    }

    uniqueOccurences := map[int]struct{}{}
    for _, occurence := range countNums {
        if _, ok := uniqueOccurences[occurence]; !ok {
            uniqueOccurences[occurence] = struct{}{}
        } else {
            return false
        }
    }

    return true
}
