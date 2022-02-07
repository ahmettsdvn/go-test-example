package service

import(
	"strings"
	"sort"
	"math"
	"reflect"
)

func SortString(t string) string {
	s := strings.Split(t, "")
    sort.Strings(s)
    return strings.Join(s, "")
}


func MaskText(t string, i uint) string {
	var s []string
	s = strings.Split(t, "")
	len_string := uint(len(s))

	if t == "" || len(t) == 1{
		s = []string{"*",}
	
	}else{
		if i>=len_string{
			for x := uint(0);  x<len_string; x++ {
				s[x] = "*"
			}
		}else{
			for x := i;  x<len_string; x++ {
				s[x] = "*"
			}
		}
	}
	
    return strings.Join(s, "")
}

func WordsSplit(arr [2]string) string {
	words := strings.Split(arr[1], ",")
	var v1,v2 bool
	var s1,s2 string
	for _, values := range words{
		v1= strings.HasPrefix(arr[0], values)
		if v1{
			s1 = values
		}
	}

	for _, values := range words{
		v2= strings.HasSuffix(arr[0], values)
		if v2{
			s2 = values
		}
	}
	if s1 == "" || s2 == ""{
		return "not possible"
	}
	return s1 + "," + s2
}

func CeilNumber(f float64) float64 {
	return (math.Ceil(f*4))/4
}

func AddUint32(x, y uint32) (uint32, bool) {
	c := x + y
	max := uint32(2147483647)
    if (x < max && y < max){
        return c, false
    }else {
		return c, true}
}

func VariadicSet(i interface{}) []interface{} {
	var m []interface{}
	rv := reflect.ValueOf(uniqueInterfaceSlice(i))
	if rv.Kind() == reflect.Slice {
        for i := 0; i < rv.Len(); i++ {
            m = append(m, rv.Index(i).Interface())
        }
    }

	return m
}

func uniqueInterfaceSlice(src interface{}) interface{} {
    srcv := reflect.ValueOf(src)
    dstv := reflect.MakeSlice(srcv.Type(), 0, 0)
    visited := make(map[interface{}]struct{})
    for i := 0; i < srcv.Len(); i++ {
        elemv := srcv.Index(i)
        if _, ok := visited[elemv.Interface()]; ok {
            continue
        }
        visited[elemv.Interface()] = struct{}{}
        dstv = reflect.Append(dstv, elemv)
    }
    return dstv.Interface()
}