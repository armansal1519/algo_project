package sortInt

func merge(ldata []int32, rdata []int32) (result []int32) {
	result = make([]int32, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return
}

func IntSingleMergeSort(data []int32) []int32 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return merge(IntSingleMergeSort(data[:middle]), IntSingleMergeSort(data[middle:]))
}

func multiMergeSort(data []int32, res chan []int32) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []int32)
	rightChan := make(chan []int32)
	middle := len(data) / 2

	go multiMergeSort(data[:middle], leftChan)
	go multiMergeSort(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- merge(ldata, rdata)
	return
}

func IntRunMultiMergeSort(data []int32) (multiResult []int32) {
	res := make(chan []int32)
	go multiMergeSort(data, res)
	multiResult = <-res
	return
}

