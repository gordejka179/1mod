package main

import (
	"bufio"
	"fmt"
	"os"
)

func lex(s string) int {
	var str_len, ind, lc, rc int
	lc = 0
	rc = 0
	var sign rune
	var s2, s3 string
	str_len = len(s)
	ind = 0
	if str_len == 3 {
		digit := rune(s[1]) - '0'
		return int(digit)
	}
	for ind < str_len {
		if ind == 1 {
			sign = rune(s[ind])
		}
		if string(s[ind]) == "(" {
			lc++
		}
		if string(s[ind]) == ")" {
			rc++
			if rc+1 == lc {
				s2 = string(s[2 : ind+1])
				s3 = string(s[ind+1 : str_len-1])
				break
			}
		}
		ind += 1
	}
	if sign == '*' {
		return lex(s2) * lex(s3)
	}
	if sign == '+' {
		return lex(s2) + lex(s3)
	}
	if sign == '-' {
		return lex(s2) - lex(s3)
	}
	return 0
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var copy string
	var ind int
	ind = 0
	copy = ""
	for ind < len(s) {
		if rune(s[ind]) != ' ' && rune(s[ind]) != '\n' {
			if (0 <= int(rune(s[ind])-'0')) && (9 >= int(rune(s[ind])-'0')) {
				copy += "(" + string(s[ind]) + ")"
			} else {
				copy += string(s[ind])
			}
		}
		ind++
	}
	fmt.Println(lex(copy))
}
