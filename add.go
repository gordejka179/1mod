package main

import "math"

func add(a, b []int32, p int) []int32 {
	var res []int32
	var rem int32 = 0
	for i := 0; i < int(math.Max(float64(len(a)), float64(len(b)))); i++ {
		if i < len(a) && i < len(b) {
			rem = rem + a[i] + b[i]
		} else if i < len(b) {
			rem += b[i]
		} else if i < len(a) {
			rem += a[i]
		}
		var last int32 = rem % int32(p)
		rem = rem / int32(p)
		res = append(res, last)
	}
	if rem > 0 {
		res = append(res, rem)
	}
	return res
}
func main() {

}
