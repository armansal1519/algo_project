package structs

type IntArr []int32

func (a IntArr) Len() int           { return len(a) }
func (a IntArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntArr) Less(i, j int) bool { return a[i] < a[j] }

type FloatArr []float64
func (a FloatArr) Len() int           { return len(a) }
func (a FloatArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FloatArr) Less(i, j int) bool { return a[i] < a[j] }


type StringArr []string
func (a StringArr) Len() int           { return len(a) }
func (a StringArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StringArr) Less(i, j int) bool { return a[i] < a[j] }

