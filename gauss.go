package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n int
	fmt.Scanln(&n)
	num := make([][]int, n)
	denum := make([][]int, n)
	ansNum := make([]int, n)
	ansDenum := make([]int, n)
	for i := range num {
		num[i] = make([]int, n+1)
		denum[i] = make([]int, n+1)
		for j := range num[i] {
			fmt.Scan(&num[i][j])
			denum[i][j] = 1
		}
	}
	isSolution := 1
	var row_change = 0
	for i := 0; i < n; i++ {
		var notNullnum, notNulldenum int
		notNullnum = 99999.0
		notNulldenum = 99999.0
		for j := row_change; j < n; j++ {
			if num[j][i] != 0 {
				notNullnum = num[j][i]
				notNulldenum = denum[j][i]
				tempRownum := make([]int, len(num[row_change]))
				copy(tempRownum, num[j])
				copy(num[j], num[row_change])
				copy(num[row_change], tempRownum)

				tempRowdenum := make([]int, len(denum[row_change]))
				copy(tempRowdenum, denum[j])
				copy(denum[j], denum[row_change])
				copy(denum[row_change], tempRowdenum)

				row_change += 1
				break
			}
		}
		//fmt.Println(notNullnum)
		//fmt.Println(notNulldenum)
		if notNullnum == 99999.0 {
			isSolution = 0
		} else {
			for k := 0; k < n+1; k++ {
				num[row_change-1][k] = num[row_change-1][k] * notNulldenum
				denum[row_change-1][k] = denum[row_change-1][k] * notNullnum
				var GCD = gcd(num[row_change-1][k], denum[row_change-1][k])
				num[row_change-1][k] = num[row_change-1][k] / GCD
				denum[row_change-1][k] = denum[row_change-1][k] / GCD
			}

			for k := row_change; k < n; k++ {
				var notNulldenum, notNullnum int
				notNullnum = 0
				notNullnum = num[k][row_change-1]
				notNulldenum = denum[k][row_change-1]
				if notNullnum != 0 {
					for p := 0; p < n+1; p++ {
						num[k][p] = num[k][p]*denum[row_change-1][p]*notNulldenum - denum[k][p]*num[row_change-1][p]*notNullnum
						denum[k][p] *= (denum[row_change-1][p] * notNulldenum)
						var GCD = gcd(num[k][p], denum[k][p])

						num[k][p] = num[k][p] / GCD
						denum[k][p] = denum[k][p] / GCD
					}
				}

			}
		}
	}
	if num[n-1][n-1] == 0 {
		isSolution = 0
	}
	if isSolution == 0 {
		fmt.Println("No solution")
	} else {
		for i := n - 1; i >= 0; i-- {
			ansNum[i] = num[i][n]
			ansDenum[i] = denum[i][n]
			for j := n - 1; j > i; j-- {
				ansNum[i] = ansNum[i]*ansDenum[j]*denum[i][j] - ansDenum[i]*ansNum[j]*num[i][j]
				ansDenum[i] = ansDenum[i] * ansDenum[j] * denum[i][j]
				var GCD = gcd(ansNum[i], ansDenum[i])
				ansNum[i] = ansNum[i] / GCD
				ansDenum[i] = ansDenum[i] / GCD
			}
			ansDenum[i] = ansDenum[i] * denum[i][i]
			var GCD = gcd(ansNum[i], ansDenum[i])
			ansNum[i] = ansNum[i] / GCD
			ansDenum[i] = ansDenum[i] / GCD
			if ansNum[i] > 0 && ansDenum[i] < 0 {
				ansNum[i] = -ansNum[i]
				ansDenum[i] = -ansDenum[i]
			}
		}
		for i := 0; i < n; i++ {
			fmt.Print(ansNum[i])
			fmt.Print("/")
			fmt.Print(ansDenum[i])
			fmt.Print("\n")
		}
	}
}

//go build gauss.go
//go run gauss.go
