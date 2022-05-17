package main

import (
	"fmt"
	"strconv"
)

type RabinKarp struct {
	Pattern string
	PatHash int64
	M       int
	Q       int64
	R       int64
	RM      int64
}

func RK(pat string, q int64) *RabinKarp {
	rk := RabinKarp{}
	rk.Pattern = pat
	rk.M = len(pat)
	rk.Q = q
	rk.R = 26
	rk.RM = 1
	for i := 1; i <= rk.M-1; i++ {
		rk.RM = (rk.R * rk.RM) % rk.Q
	}
	rk.PatHash = rk.hash([]byte(pat), rk.M)
	return &rk
}

func main() {
	var n int64
	var str, pattern string
	fmt.Scan(&n)
	fmt.Scan(&pattern, &str)

	rk := RK(pattern, n)
	matches := "Matches: "
	spurious := "Spurious hits: "
	rk.search(str, &matches, &spurious)
	fmt.Println(matches)
	fmt.Println(spurious)
}

func (c RabinKarp) search(txt string, matches *string, spurious *string) {
	n := len(txt)
	txtRunes := []byte(txt)
	txtHash := c.hash(txtRunes, c.M)
	if c.PatHash == txtHash {
		if c.check(txtRunes, 0) {
			*matches = *matches + "0 "
		} else {
			*spurious = *spurious + "0 "
		}
	}
	for i := c.M; i < n; i++ {
		txtHash = (txtHash + c.Q - c.RM*int64(txtRunes[i-c.M])%c.Q) % c.Q
		txtHash = (txtHash*c.R + int64(txtRunes[i])) % c.Q
		if c.PatHash == txtHash {
			if c.check(txtRunes, i-c.M+1) {
				*matches = *matches + strconv.Itoa(i-c.M+1) + " "
			} else {
				*spurious = *spurious + strconv.Itoa(i-c.M+1) + " "
			}
		}
	}
}

func (c RabinKarp) check(txt []byte, start int) bool {
	for i := 0; i < c.M; i++ {
		if txt[i+start] != c.Pattern[i] {
			return false
		}
	}
	return true
}

func (c RabinKarp) hash(key []byte, m int) int64 {
	var h int64 = 0
	for j := 0; j < m; j++ {
		h = (c.R*h + int64(key[j])) % c.Q
	}
	return h
}
