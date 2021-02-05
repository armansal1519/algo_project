package sortInt

type minHeap struct {
	arr []int32
}

func NewMinHeap(arr []int32) *minHeap {
	mh := &minHeap{
		arr: arr,
	}
	return mh
}

func (m *minHeap) leftchildIndex(index int32) int32 {
	return 2*index + 1
}

func (m *minHeap) rightchildIndex(index int32) int32 {
	return 2*index + 2
}

func (m *minHeap) swap(first, second int32) {
	temp := m.arr[first]
	m.arr[first] = m.arr[second]
	m.arr[second] = temp
}

func (m *minHeap) leaf(index int32, size int32) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (m *minHeap) downHeapify(current int32, size int32) {
	if m.leaf(current, size) {
		return
	}
	smallest := current
	leftChildIndex := m.leftchildIndex(current)
	rightRightIndex := m.rightchildIndex(current)
	if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < size && m.arr[rightRightIndex] < m.arr[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest, size)
	}
	return
}

func (m *minHeap) buildMinHeap(size int32) {
	for index := (size / 2) - 1; index >= 0; index-- {
		m.downHeapify(index, size)
	}
}

func (m *minHeap) sort(size int32) {
	m.buildMinHeap(size)
	for i := size - 1; i > 0; i-- {
		// Move current root to end
		m.swap(0, i)
		m.downHeapify(0, i)
	}
}

