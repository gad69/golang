package main

import (
	"fmt"
	"math"
)

type parameters struct {
	arr [100]float64
	n   int
} // structure to pass into 3 functions

func (p parameters) get_Average() float64 {
	// passing parameter structure variable into function  getAverage()
	// function of mean

	var i int = 0
	var avgr float64 = 0
	var sum float64 = 0
	for i = 0; i < p.n; i++ {
		sum += p.arr[i]
	}
	avgr = float64(sum / float64(p.n))
	return avgr
}
func (p parameters) get_standard_dev() float64 {
	// function for getstandard_dev()
	var sum1, average, variance float64
	var standard_deviation float64 = 0
	var j, m int
	var sum_st float64
	for m = 0; m < p.n; m++ {
		sum1 = sum1 + p.arr[m]
	}
	average = sum1 / float64(p.n)
	for j = 0; j < p.n; j++ {
		sum_st = sum_st + float64(math.Pow((p.arr[j]-average), 2))
	}
	variance = sum_st / float64(p.n-1)
	standard_deviation = float64(math.Sqrt((variance)))
	return standard_deviation
}
func (p parameters) get_median() float64 {
	// function for getmedian()
	var k, l int
	var temp float64
	for k = 0; k < p.n-1; k++ { // median using bubble sort
		for l = 0; l < p.n-k-1; l++ {
			if p.arr[l] > p.arr[l+1] {
				temp = p.arr[l+1]
				p.arr[l+1] = p.arr[l]
				p.arr[l] = temp
			}
		}
	}
	//sort(arr,arr+n);   can even use this
	if p.n%2 != 0 {
		return float64(p.arr[p.n/2])
	}
	return float64((p.arr[(p.n-1)/2] + p.arr[p.n/2]) / 2.0) // formula for finding median for odd case
}
func main() {
	// infinite loop
	var array [100]float64
	var n, i int
	fmt.Printf("enter number of elememts for the test case \t")
	fmt.Scan(&n)

	if n > 100 || n <= 0 {
		fmt.Printf("enter the number between 0 to 101 \n")
	}
	if n <= 100 && n > 0 {
		for i = 0; i < n; i++ {
			fmt.Printf("enter number  %d \t", i+1)
			fmt.Scan(&array[i])
		}
		p := parameters{array, n}                                                                // structure variable used in getAverage()	                                                               // structure variable used in getstandard_dev()                                                           // structure variable used in getmedian()
		fmt.Println("\nthe average of the entered elememnts is ", p.get_Average())               // accesing the getAverage() method using p
		fmt.Println("the standard deviation of the entered elememnts is ", p.get_standard_dev()) // accesing the getstandard_dev() method using p
		fmt.Println("the median of the entered elememnts is", p.get_median())                    // accesing the getmedian() method using p

	}
}
