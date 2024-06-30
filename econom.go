package main

import (
	"bufio"
	"fmt"
	"os"
)

func lex(s string, map1 map[string]int) {
	var str_len, ind, lc, rc int
	lc = 0
	rc = 0
	var s2, s3 string
	str_len = len(s)
	ind = 0
	if str_len != 3 {
		for ind < str_len {
			if string(s[ind]) == "(" {
				lc++
			}
			if string(s[ind]) == ")" {
				rc++
				if rc+1 == lc {
					s2 = string(s[2 : ind+1])
					s3 = string(s[ind+1 : str_len-1])
					if len(s) != 3 {
						map1[s] = 1
					}
					if len(s2) != 3 {
						map1[s2] = 1
					}
					if len(s3) != 3 {
						map1[s3] = 1
					}
					break
				}
			}
			ind += 1
		}
		lex(s2, map1)
		lex(s3, map1)
	}
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	map1 := make(map[string]int)
	var copy string
	var ind int
	ind = 0
	copy = ""
	for ind < len(s) {
		if rune(s[ind]) != ' ' && rune(s[ind]) != '\n' {
			if ('a' <= rune(s[ind])) && ('z' >= rune(s[ind])) {
				copy += "(" + string(s[ind]) + ")"
			} else {
				copy += string(s[ind])
			}
		}
		ind++
	}
	lex(copy, map1)
	fmt.Println(len(map1))
}
