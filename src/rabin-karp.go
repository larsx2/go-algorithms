package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

func RabinKarpSearch(text string, pattern string) int {
	patternHash := getMD5(pattern)

	for i, _ := range text {
		if i+len(pattern) > len(text) {
			break
		}

		sub := text[i : i+len(pattern)]
		subHash := getMD5(sub)

		fmt.Println("Checking:", sub, "==", pattern)
		if patternHash == subHash && pattern == sub {
			return i
		}
	}

	return -1
}

func main() {
	text := "foundmedudefoundmenow"
	pattern := "dude"

	index := RabinKarpSearch(text, pattern)

	if index != -1 {
		fmt.Println("Found at index:", index)
	} else {
		fmt.Println("Not found")
	}

}
