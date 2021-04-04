package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	fmt.Printf("enter the file name\n")
	fmt.Scanln(&filename)
	file, ferr := os.Open(filename)

	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)
	ctr := 0
	ctr1 := 0
	fmt.Printf(" the function prototypes are\n")
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Text() != "" {
			ctr++
		}
		if substr(line, "//PROXY START") == line {

			fmt.Println(strings.Trim(line, "//PROXY START"))
			ctr1++

		}

	}
	fmt.Printf(" total no of Lines excluding blank lines in the file %d\n", ctr)
	fmt.Printf(" total no of functions %d\n", ctr1)

}

func substr(str1 string, str2 string) string {
	var m, n, i int
	var j, flag int
	m = len(str1)
	n = len(str2)
	if m < 1 || n < 1 {
		return "null"
	}
	if m < n {
		return "null"

	}

	for i = 0; i <= (m - n); i++ {

		for j = i; j < (i + n); j++ {
			flag = 1
			if str1[j] != str2[j-i] {
				flag = 0
				break

			}

		}
		if flag == 1 {
			break
		}

	}
	if flag == 1 {
		return str1

	}
	return "null"
}
