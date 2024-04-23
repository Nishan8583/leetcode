func dailyTemperatures(temperatures []int) []int {

	return_values := make([]int, len(temperatures))
	return_values[len(temperatures)-1] = 0

	for i := len(temperatures) - 2; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				return_values[i] = j - i
				break
			}
		}

	}

	return return_values
}