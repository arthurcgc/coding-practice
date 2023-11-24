func calcRealValue(currentValue, i int, s string, romanValues map[byte]int) (int, int) {
    j := i - 1
    if j < 0 {
        return i, currentValue
    }
    nextValue := romanValues[s[j]]
    if nextValue < currentValue {
        i--
        return i, currentValue - nextValue
    }
    return i, currentValue
}

func romanToInt(s string) int {
    romanValues := map[byte]int{'I':1, 'V': 5, 'X':10, 'L':50, 'C': 100, 'D':500, 'M':1000}
    sum := 0
    for i := len(s) - 1; i >= 0; i-- {
        currentValue := romanValues[s[i]]
        var realValue int
        i, realValue = calcRealValue(currentValue, i, s, romanValues)
        sum += realValue
    }

    return sum
}
