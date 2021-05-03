package main

import (
	"fmt"
	"math"
)
// t dis function
func f(x float64, n int) float64 {
	return (gamma((float64(n)+1)/2)/((gamma(float64(n)/2))*(math.Sqrt(float64(n)*3.14))))*(1 /math.Pow((1+((x*x)/float64(n))), ((float64(n)+1)/2)))
}
// gamma func
func gamma(n float64) float64 {
	if n == 1 {
		return 1
	}
	if n == 0.5 {
		return 1.772453
	}
	return gamma(n-1) * (n - 1)
}
// simpsons for tail = 1
func Simpson_t1(t float64 ,d int)  {
	var a,b,h,x,integral,half,sum float64
	a = -t
	b = t
	h = math.Abs(b - a) /float64(128)

	for i := 1; i <128; i++ {
		x = a + float64(i)*h
		if i%2 == 0 {
			sum = sum + 2*f(x,d)
		}else{
		sum = sum + 4*f(x,d)
		}
	}
	integral = (h / 3) * (f(a,d) + f(b,d) + sum)
	half = ((1-integral)/2) + integral
	fmt.Println(half)
}
// simpsons for t = 2
func Simpson_t2(t float64 ,d int)  {
	var a,b,h,x,integral,sum float64
	a = -t
	b = t
	h = (b - a) /float64(128)

	for i := 1; i <128; i++ {
		x = a + float64(i)*h
		if i%2 == 0 {
			sum = sum + 2*f(x,d)
		}else{
		sum = sum + 4*f(x,d)
		}
	}
	integral = (h / 3) * (f(a,d) + f(b,d) + sum)
	fmt.Println(integral)

}
func main() {
	var n int
	var t [100]float64
	var degre_of_freedm [100] int 
	var tail [100] int
	fmt.Println("enter number of elememts in the table (in given excel sheet it is 6)")
	fmt.Scan(&n)
	if n <= 100 && n > 0 {
		for i:= 0; i < n; i++ {
			fmt.Scan(&t[i], &degre_of_freedm[i], &tail[i])

		}

	}
	for i:=0; i<n ; i++ {
		if tail[i] == 1 {
		
			Simpson_t1(t[i],degre_of_freedm[i]) 
			

		}
		if tail[i] == 2{
			Simpson_t2(t[i],degre_of_freedm[i]) 


		}
	}
}
