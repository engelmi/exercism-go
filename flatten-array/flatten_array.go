package flatten

func Flatten(v interface{}) interface{} {
	res := []interface{}{}
	flatten(v, &res)
	return res
}

/*
goos: linux
goarch: amd64
pkg: flatten
BenchmarkFlatten-8   	  788156	      2183 ns/op	    1328 B/op	      20 allocs/op
PASS
ok  	flatten	1.740s
*/
func flatten(v interface{}, res *[]interface{}) {
	if v == nil {
		return
	}
	if slice, ok := isSlice(v); ok {
		for _, elem := range slice {
			flatten(elem, res)
		}
	} else {
		*res = append(*res, v)
	}
}

/*
goos: linux
goarch: amd64
pkg: flatten
BenchmarkFlatten-8   	  509844	      2947 ns/op	    2064 B/op	      43 allocs/op
PASS
ok  	flatten	1.530s
*/
func flatten2(v interface{}, res *[]interface{}) {

	if slice, ok := v.([]interface{}); ok {
		for _, elem := range slice {
			if elem != nil {
				if slice2, ok2 := elem.([]interface{}); ok2 {
					flatten2(slice2, res)
				} else {
					*res = append(*res, elem)
				}
			}
		}
	} else if v != nil {
		*res = append(*res, v)
	}

}

func isSlice(v interface{}) ([]interface{}, bool) {
	slice, ok := v.([]interface{})
	return slice, ok
}
