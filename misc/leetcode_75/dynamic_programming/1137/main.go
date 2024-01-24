package main

func tribWithMemo(n int, memo map[int]int) int {
    if v, ok := memo[n]; ok {
        return v
    }
    result := tribWithMemo(n-3, memo) + tribWithMemo(n-2, memo) + tribWithMemo(n-1, memo)
    memo[n] = result
    return result
}

func tribonacci(n int) int {
    // if n == 0 {
    //     return 0
    // }
    // if n == 1 || n == 2 {
    //     return 1
    // }
    // return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
    memo := map[int]int{0:0, 1:1, 2:1}
    return tribWithMemo(n, memo)
}
