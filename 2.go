package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//PROXY START main
func main() {

	var filename string
	var fc []string
	var funcnames [20]string // string array for holding funcnames
	var start [20]string     // string array to hold start with function
	var end [20]string       // string array to hold end proxy with function
	j := 0
	var Lines_of_file []string // string  array for storing lines of code
	var temp string
	fmt.Println("Enter  file with go extension ")
	fmt.Scanln(&filename)
	file, ferr := os.Open(filename)
	if ferr != nil { // file error handling
		fmt.Println("PLEASE ENTER A VALID FILE NAME")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fc = append(fc, scanner.Text())
	}
	for i := 0; i < len(fc); i++ { //parsing  the file to remove space or tab character in the lines
		temp = remove_Spaces(fc[i]) // file  parsing  to remove space
		fc[i] = temp
	}
	for i := 0; i < len(fc); i++ {
		if fc[i] != "" { // remove empty lines with appending non empty lines
			Lines_of_file = append(Lines_of_file, fc[i])
		}
	}

	fmt.Println("\nLOC  for the whole program without blanks ", len(Lines_of_file))

	for i := 0; i < len(Lines_of_file); i++ {
		if substr(Lines_of_file[i], "//PROXYSTART") == Lines_of_file[i] {
			funcnames[j] = strings.Trim(Lines_of_file[i], "//PROXYSTART") // triming the proxystart from file
			j++

		}
	}
	fmt.Println("FUNCTION PROTOTYPES ARE AND LOC OF EACH \n")
	for i := 0; i <= j; i++ {

		start[i] = "//PROXYSTART" + funcnames[i]
		end[i] = "//PROXYEND" + funcnames[i]
	}
	ctr := 0
	var t = false
	for i := 0; i <= j; i++ { // counting lines of code for each function
		for k := 0; k < len(Lines_of_file); k++ {
			if !t && Lines_of_file[k] == start[i] {
				t = true
			}
			if t {
				ctr++
			}
			if t && Lines_of_file[k] == end[i] {
				t = false
				fmt.Printf("NO_OF_LINES_IN of %s - %d\n", funcnames[i], (ctr - 2))
				ctr = 0
			}

		}
	}
}

//PROXY END main

//PROXY START remove_Spaces
func remove_Spaces(str string) string { //function for removing spaces in the file container
	temp := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " && string(str[i]) != "	" {
			temp = temp + string(str[i])
		}
	}
	return temp
}

//PROXY END remove_Spaces

//PROXY START substr
func substr(str1 string, str2 string) string { // function for substring matching
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

//PROXY END substr
