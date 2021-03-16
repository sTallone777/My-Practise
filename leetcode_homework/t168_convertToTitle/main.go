package main

import "fmt"

func convertToTitle(columnNumber int) string {
	str := ""
	base := 26
	for columnNumber > 0 {
		tmp := columnNumber % base
		if tmp == 0 {
			tmp = base
			columnNumber = columnNumber/base - 1
		} else {
			columnNumber /= base
		}

		str = string(rune(64+tmp)) + str

	}

	return str
}

func main() {
	fmt.Printf("convert to title: %s\r\n", convertToTitle(1))
}
