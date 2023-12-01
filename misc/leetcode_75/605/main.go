package main

func canPlaceFlower(flowerbed []int, index int) bool {
    left := index - 1
    right := index + 1
    if left < 0 {
        if right >= len(flowerbed) {
            return flowerbed[index] == 0
        }
        return flowerbed[index] == 0 && flowerbed[right] == 0
    }

    if right >= len(flowerbed) {
        return flowerbed[index] == 0 && flowerbed[left] == 0
    }

    return flowerbed[index] == 0 && flowerbed[left] == 0 && flowerbed[right] == 0
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    canPlaceCount := 0
    for i := range flowerbed {
        if canPlaceFlower(flowerbed, i) {
            flowerbed[i] = 1
            canPlaceCount++
        }
    }
    return canPlaceCount >= n
}