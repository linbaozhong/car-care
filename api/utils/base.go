package utils

type CompareFunc func(interface{}, interface{}) bool

func IndexOf(a []interface{}, e interface{}, cmp CompareFunc) int {
	n := len(a)
	for i := 0; i < n; i++ {
		if cmp(e, a[i]) {
			return i;
		}
	}
	return -1;
}
// 参数e是否包含在切片a中
func Contains(a []interface{}, e interface{}) bool {
	return IndexOf(a, e, func(i interface{}, i2 interface{}) bool {
		return i == i2
	}) != -1
}
// slice是否相等
func IsEqualSlice(a,b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	
	return true
}
