package main

import "fmt"

func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }
	retStr := ""
	retNum := make([]string, numRows)
	k := -1
	j := 0
	for _, v := range s {
		if j == 0 || j == numRows-1 {
			k = -k
		}
		retNum[j] = retNum[j] + string(v)
		j = j + k
	}

	for _, v := range retNum {
		retStr += v
	}
	return retStr
}

func main() {
	// fmt.Println(convert("leetcodeishiring", 4))
	fmt.Println(convert("A", 100))
}
