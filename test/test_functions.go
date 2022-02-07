package assignment

import (
	services "testexample/service"
)

func AddUint32(x, y uint32) (uint32, bool) {

	return services.AddUint32(x, y)
}

func CeilNumber(f float64) float64 {
	return services.CeilNumber(f)
}

func AlphabetSoup(s string) string {
	return services.SortString(s)
}

func StringMask(s string, n uint) string {
	return services.MaskText(s, n)
}

func WordSplit(arr [2]string) string {
	return services.WordsSplit(arr)
}

func VariadicSet(i interface{}) []interface{} {
	return services.VariadicSet(i)
}
