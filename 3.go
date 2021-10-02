package main

import (
	"fmt"
)

func get_Average(arr [100]float64, n int) float64 {

	var i int = 0
	var avgr float64 = 0
	var sum float64 = 0
	for i = 0; i < n; i++ {
		sum += arr[i]
	}
	avgr = float64(sum / float64(n))
	return avgr
}
func get_b1(X [100]float64, Y [100]float64, n int) float64 {
	var B1 float64 = 0
	var sop float64 = 0
	var square float64 = 0
	for i := 0; i < n; i++ {
		sop = sop + X[i]*Y[i] // caliculating sigma(xi*yi)

		square = square + X[i]*X[i] // caliculating sigma(square(Xi))
	}
	B1 = (sop - (float64(n) * float64(get_Average(X, n)) * (get_Average(Y, n)))) / (square - (float64(n) * (get_Average(X, n)) * (get_Average(X, n))))
	return B1
}
func get_bo(X [100]float64, Y [100]float64, n int) float64 {
	var bo float64 = 0
	bo = get_Average(Y, n) - (get_b1(X, Y, n) * get_Average(X, n))
	return bo
}
func main() {
	var X [100]float64
	var Y [100]float64
	var n, i int
	fmt.Printf("enter number of elememts for the test case\t")
	fmt.Scan(&n)
	fmt.Printf("ENTER X AND   Y VALUES ONE AFTER OTHER \n")

	if n > 100 || n <= 0 { // handling error user  can enter between 0 -101 excluding them
		fmt.Printf("enter the number between 0 to 101 exited succesfully \n")
		return // handling error for n < 0 and n> 100
	}
	if n == 1 {
		fmt.Printf("please enter more than 1 value\n")
		return // handling error for n = 1
	}
	if n <= 100 && n > 0 {
		for i = 0; i < n; i++ {
			fmt.Printf("enter X_%d \t", i+1)
			fmt.Scan(&X[i])
			fmt.Printf("enter Y_%d \t", i+1)
			fmt.Scan(&Y[i])
		}

	}
	fmt.Printf("the vaue of bo is %f \n", get_bo(X, Y, n))
	fmt.Printf("the vaue of b1 is %f \n", get_b1(X, Y, n))
	fmt.Printf("the vaue of b1 is %f \n", get_b1(X, Y, n))

}
