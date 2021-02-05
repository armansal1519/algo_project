package sortFloat

func BubbleSort(numbers []float64) []float64 {

	for i := len(numbers); i > 0; i-- {

		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}             }
	}
	return numbers
}
