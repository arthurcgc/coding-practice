package main

func uniques(arr []int) map[int]struct{} {
    uniqueInts := map[int]struct{}{}
    for _, v := range arr {
        if _, ok := uniqueInts[v]; !ok {
            uniqueInts[v] = struct{}{}
        }
    }
    return uniqueInts
}

func generateUniqueList(uniqueMap1 map[int]struct{}, uniqueMap2 map[int]struct{}) []int {
    ans := []int{}
    for k1, _ := range uniqueMap1 {
        if _, ok := uniqueMap2[k1]; ok {
            continue
        } else {
            ans = append(ans, k1)
        }
    }
    return ans
}

func findDifference(nums1 []int, nums2 []int) [][]int {
    un1 := uniques(nums1)
    un2 := uniques(nums2)

    ans := make([][]int, 2)
    ans[0] = generateUniqueList(un1, un2)
    ans[1] = generateUniqueList(un2, un1)

    return ans
}
