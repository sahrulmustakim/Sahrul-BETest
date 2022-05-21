package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Write your sentences : ")
	fmt.Scanln(&str)

	var out = longestPalindrome(str)
	fmt.Println("output : ", out)
}

func longestPalindrome(s string) string {
	start, end := 0, 0
	part := 1
	maxPal, maxStart, maxEnd := 0, 0, 0
	lenS := len(s)
	nowBreak := false
	for i := 0; i < lenS; i++ {
		// looping for first try, the even pal
		for {
			if i-part < 0 || i+part >= lenS {
				nowBreak = true
			} else {

				if s[i-part] == s[i+part] {
					start = i - part
					end = i + part
				} else {
					nowBreak = true
				}
			}
			if nowBreak {
				if end-start+1 > maxPal {
					maxPal = end - start + 1
					maxStart = start
					maxEnd = end
				}
				nowBreak = false // reset it now!
				part = 1
				break
			}
			part = part + 1
		}

		// looping for second
		if i+1 < lenS && s[i] == s[i+1] {
			// it found
			start, end = i, i+1
			for {

				if i-part < 0 || i+1+part >= lenS {
					nowBreak = true
				} else {
					if s[i-part] == s[i+1+part] {
						start = i - part
						end = i + 1 + part
					} else {
						nowBreak = true
					}
				}
				if nowBreak {
					if end-start+1 > maxPal {
						maxPal = end - start + 1
						maxStart = start
						maxEnd = end
					}
					nowBreak = false // reset it now!
					part = 1
					break
				}
				part = part + 1
			}
		}

	}
	return s[maxStart : maxEnd+1]
}
