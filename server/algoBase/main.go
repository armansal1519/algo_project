package algoBase

import (
	"server/structs"
	"sort"
)

func DefaultSortInt(arr []int32) structs.IntArr {
	i:=structs.IntArr(arr)
	sort.Sort(i)
	return i
}

func DefaultSortFloat(arr []float64) []float64 {
	sort.Float64s(arr)
	return arr
}

func SortString(arr []string) []string {
	sort.Strings(arr)
	return arr
}
