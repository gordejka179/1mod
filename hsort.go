package main

func hsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	for i := (n - 1) / 2; i >= 0; i-- {
		down(n, i, less, swap)
	}
	for i := n - 1; i > 0; i-- {
		swap(0, i)
		down(i, 0, less, swap)
	}
}

func down(n, i int, less func(i, j int) bool, swap func(i, j int)) {
	for 2*i+1 < n {
		maxChild := 2*i + 1
		if maxChild+1 < n && less(maxChild, maxChild+1) {
			maxChild++
		}
		if !less(i, maxChild) {
			break
		}
		swap(i, maxChild)
		i = maxChild
	}
}

func main() {

}
