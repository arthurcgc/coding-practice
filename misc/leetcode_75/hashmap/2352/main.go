package main

func row2String(grid [][]int, rowIndex int) string {
	return fmt.Sprintf("%v", grid[rowIndex])
}

func column2String(grid [][]int, columnIndex int) string {
	cs := "["
	for i := 0; i < len(grid[0]); i++ {
		if i+1 < len(grid[0]) {
			cs += strconv.Itoa(grid[i][columnIndex]) + " "
		} else {
			cs += strconv.Itoa(grid[i][columnIndex]) + "]"
		}
	}
	return cs
}

func add2Map(m map[string]int, s string) {
    if _, ok := m[s]; ok {
        m[s]++
    } else {
        m[s] = 1
    }
}
func countPair(s string, counter *int, m1 map[string]int, m2 map[string]int) {
    if _, ok := m1[s]; ok {
        times := m1[s]*m2[s]
        *counter += times
        delete(m1, s)
        delete(m2, s)
    }
}

func equalPairs(grid [][]int) int {
    rows := map[string]int{}
    columns := map[string]int{}
    equalPairs := 0
    for i, _ := range grid {
        rowString := row2String(grid, i)
        columnString := column2String(grid, i)
        add2Map(rows, rowString)
        add2Map(columns, columnString)
    }

    for r, _ := range rows {
        countPair(r, &equalPairs, columns, rows)
    }
    return equalPairs
}
