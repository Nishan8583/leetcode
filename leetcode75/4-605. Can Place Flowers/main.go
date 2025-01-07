package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i, v := range flowerbed {
		if v == 1 {
			continue
		}

		empty_left := (i == 0) || (flowerbed[i-1] == 0)
		empty_right := (i == (len(flowerbed) - 1)) || (flowerbed[i+1] == 0)

		if empty_left && empty_right {
			flowerbed[i] = 1
			count += 1
			if count >= n {
				return true
			}
		}
	}

	return count >= n
}
