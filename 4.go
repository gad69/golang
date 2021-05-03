package main

import (
	"fmt"
	"math"
)

func get_bo(X [100]float64, Y [100]float64, n int) float64 {
	var bo float64 = 0
	bo = get_Average(Y, n) - (get_b1(X, Y, n) * get_Average(X, n))
	return bo
} // function for b0

func sum(array [100]float64) float64 {
	result := 0.0
	for _, v := range array {
		result += v
	}
	return result
} // function for sum

func correlationCoefficient(X [100]float64, Y [100]float64, n int) float64 {
	// corelation function
	var sum_X, sum_Y, sum_XY, squareSum_X, squareSum_Y, corr float64
	var i int
	for i = 0; i < n; i++ {
		sum_X = sum_X + X[i]
		sum_Y = sum_Y + Y[i]
		sum_XY = sum_XY + X[i]*Y[i]
		squareSum_X = squareSum_X + X[i]*X[i]
		squareSum_Y = squareSum_Y + Y[i]*Y[i]
	}

	// use formula for calculating correlation coefficient.
	corr = (float64(n)*sum_XY - sum_X*sum_Y) / (math.Sqrt((float64(n)*squareSum_X - sum_X*sum_X) * (float64(n)*squareSum_Y - sum_Y*sum_Y)))

	return corr
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
func get_standard_dev(arr [100]float64, n int) float64 {
	//function get_standard_dev()
	var sum1, average, variance float64
	var standard_deviation float64 = 0
	var j, m int
	var sum_st float64
	for m = 0; m < n; m++ {
		sum1 = sum1 + arr[m]
	}
	average = sum1 / float64(n)
	for j = 0; j < n; j++ {
		sum_st = sum_st + float64(math.Pow((arr[j]-average), 2))
	}
	variance = sum_st / float64(n-1)
	standard_deviation = float64(math.Sqrt((variance)))
	return standard_deviation
} // function for standard dev
func get_Average(arr [100]float64, n int) float64 {
	// passing parameter structure variable into  function get_Average()

	var i int = 0
	var avgr float64 = 0
	var sum float64 = 0
	for i = 0; i < n; i++ {
		sum += arr[i]
	}
	avgr = float64(sum / float64(n))
	return avgr
}
func main() {
	var Totalloc, calc, dat, IO, calca, da, I, avh [100]float64
	var Methods [100]float64
	var identifier_1 [100]string
	var type_1 [100]string // variables for using in proxy table

	var BASEADD_METHOD [100]float64
	var BASEADD [100]string // variables for using in base table
	var SIMILAR_BA [100]string

	var NEW_METHOD [100]float64
	var NEWADD [100]string // variable used in new table
	var SIMILAR_NEW [100]string

	var ESTLOC [100]float64
	var ACLOC [100]float64
	var ESTDURATION [100]float64 // variables used in project table
	var ACTUALDURATION [100]float64
	var identifier_2 [100]string

	var n1, n2, n3, i, j, n4 int
	var SIMILAR_BA_loc_calc, SIMILAR_BA_loc_data, SIMILAR_BA_loc_IO, BA, NO, E_LOC, N, H, productivity, duration, RSIZE, RDURATION float64
	var SIMILAR_NEW_loc_calc, SIMILAR_NEW_loc_data, SIMILAR_NEW_loc_IO float64
	var avg_lognormal_calc, avg_lognormal_data, avg_lognormal_IO, cr, ca float64
	var sdev_lognormal_calc, sdev_lognormal_data, sdev_lognormal_IO float64
	var calc_VS, calc_S, calc_M, calc_L, calc_VL, b_1, b_o float64
	var data_VS, data_S, data_M, data_L, data_VL, M, B, D float64
	var IO_VS, IO_S, IO_M, IO_L, IO_VL float64
	var calc_VS_LB, calc_S_LB, calc_M_LB, calc_L_LB, calc_VL_LB float64
	var calc_VS_UB, calc_S_UB, calc_M_UB, calc_L_UB, calc_VL_UB float64
	var data_VS_LB, data_S_LB, data_M_LB, data_L_LB, data_VL_LB float64
	var data_VS_UB, data_S_UB, data_M_UB, data_L_UB, data_VL_UB float64
	var IO_VS_LB, IO_S_LB, IO_M_LB, IO_L_LB, IO_VL_LB float64
	var IO_VS_UB, IO_S_UB, IO_M_UB, IO_L_UB, IO_VL_UB float64

	fmt.Println("enter the value of M (MODIFIED) HERE ITS IS 14")
	fmt.Scan(&M)
	fmt.Println("enter the value of D(DELETED) HERE IT IS 14")
	fmt.Scan(&D)
	fmt.Println("enter the value of B(BASE SIZE) HERE IT IS 369")
	fmt.Scan(&B)
	fmt.Println("enter number of elememts used in PROXYTABLE for  record of proxies used in past history HERE IT IS 33")
	fmt.Scan(&n1)
	if n1 <= 100 && n1 > 0 {
		for i = 0; i < n1; i++ {
			fmt.Scan(&identifier_1[i], &Totalloc[i], &Methods[i], &type_1[i])
		}

	}
	var k, l, m int
	for j = 0; j < n1; j++ {
		avh[j] = Totalloc[j] / Methods[j]
		if type_1[j] == "Calculation" {

			calca[k] = (Totalloc[j] / Methods[j])
			calc[k] = math.Log10(Totalloc[j] / Methods[j])
			k++
		}
		if type_1[j] == "Data" {
			da[l] = (Totalloc[j] / Methods[j])
			dat[l] = math.Log10(Totalloc[j] / Methods[j])

			l++
		}
		if type_1[j] == "I/O" {
			I[m] = (Totalloc[j] / Methods[j])
			IO[m] = math.Log10(Totalloc[j] / Methods[j])
			m++
		}
	}
	fmt.Print("\n")
	fmt.Println("the lognormal average of the calculation is ", get_Average(calc, k))
	fmt.Println("the lognormal average of the Data is ", (get_Average(dat, l)))
	fmt.Println("the lognormal average of the I/0 is ", (get_Average(IO, m)))
	fmt.Println("the lognormal standard deviation of the calculation is ", (get_standard_dev(calc, k)))
	fmt.Println("the lognormal standard deviation of the Data is ", (get_standard_dev(dat, l)))
	fmt.Println("the lognormal standard deviation of the I/O is ", (get_standard_dev(IO, m)))
	avg_lognormal_calc = (get_Average(calc, k))
	avg_lognormal_data = (get_Average(dat, l))
	avg_lognormal_IO = (get_Average(IO, m))
	sdev_lognormal_calc = (get_standard_dev(calc, k))
	sdev_lognormal_data = (get_standard_dev(dat, l))
	sdev_lognormal_IO = (get_standard_dev(IO, m))
	calc_VS = avg_lognormal_calc - 2*sdev_lognormal_calc
	calc_S = avg_lognormal_calc - sdev_lognormal_calc
	calc_M = avg_lognormal_calc
	calc_L = avg_lognormal_calc + sdev_lognormal_calc
	calc_VL = avg_lognormal_calc + 2*sdev_lognormal_calc
	data_VS = avg_lognormal_data - 2*sdev_lognormal_data
	data_S = avg_lognormal_data - sdev_lognormal_data
	data_M = avg_lognormal_data
	data_L = avg_lognormal_data + sdev_lognormal_data
	data_VL = avg_lognormal_data + 2*sdev_lognormal_data
	IO_VS = avg_lognormal_IO - 2*sdev_lognormal_IO
	IO_S = avg_lognormal_IO - sdev_lognormal_IO
	IO_M = avg_lognormal_IO
	IO_L = avg_lognormal_IO + sdev_lognormal_IO
	IO_VL = avg_lognormal_IO + 2*sdev_lognormal_IO
	fmt.Print("\n")
	fmt.Println("mid points for all buckets in lognormal form")
	fmt.Println("calculation")
	fmt.Println(calc_VS, calc_S, calc_M, calc_L, calc_VL)
	fmt.Println("data")
	fmt.Println(data_VS, data_S, data_M, data_L, data_VL)
	fmt.Println("I/O")
	fmt.Println(IO_VS, IO_S, IO_M, IO_L, IO_VL)

	fmt.Print("\n")

	fmt.Println("mid points for all buckets in antilog form (OBJECT CATEGORY SIZES)")
	fmt.Print("\n")
	fmt.Println("(OBJECT CATEGORY SIZES) for Calculation")
	fmt.Println(math.Pow(10, calc_VS), math.Pow(10, calc_S), math.Pow(10, calc_M), math.Pow(10, calc_L), math.Pow(10, calc_VL))
	fmt.Println("(OBJECT CATEGORY SIZES) for data")
	fmt.Println(math.Pow(10, data_VS), math.Pow(10, data_S), math.Pow(10, data_M), math.Pow(10, data_L), math.Pow(10, data_VL))
	fmt.Println("(OBJECT CATEGORY SIZES) for I/O")
	fmt.Println(math.Pow(10, IO_VS), math.Pow(10, IO_S), math.Pow(10, IO_M), math.Pow(10, IO_L), math.Pow(10, IO_VL))

	fmt.Print("\n")
	RSIZE = 319
	RDURATION = 145

	calc_M_LB = avg_lognormal_calc - 0.5*sdev_lognormal_calc
	calc_M_UB = avg_lognormal_calc + 0.5*sdev_lognormal_calc
	calc_VL_LB = avg_lognormal_calc + 1.5*sdev_lognormal_calc
	calc_VL_UB = math.Inf(1)
	calc_L_LB = avg_lognormal_calc + 0.5*sdev_lognormal_calc
	calc_L_UB = avg_lognormal_calc + 1.5*sdev_lognormal_calc
	calc_S_LB = avg_lognormal_calc - 1.5*sdev_lognormal_calc
	calc_S_UB = avg_lognormal_calc - 0.5*sdev_lognormal_calc
	calc_VS_LB = 0.0
	calc_VS_UB = avg_lognormal_calc - 1.5*sdev_lognormal_calc

	data_M_LB = avg_lognormal_data - 0.5*sdev_lognormal_data
	data_M_UB = avg_lognormal_data + 0.5*sdev_lognormal_data
	data_VL_LB = avg_lognormal_data + 1.5*sdev_lognormal_data
	data_VL_UB = math.Inf(1)
	data_L_LB = avg_lognormal_data + 0.5*sdev_lognormal_data
	data_L_UB = avg_lognormal_data + 1.5*sdev_lognormal_data
	data_S_LB = avg_lognormal_data - 1.5*sdev_lognormal_data
	data_S_UB = avg_lognormal_data - 0.5*sdev_lognormal_data
	data_VS_LB = 0.0
	data_VS_UB = avg_lognormal_data - 1.5*sdev_lognormal_data

	IO_M_LB = avg_lognormal_IO - 0.5*sdev_lognormal_IO
	IO_M_UB = avg_lognormal_IO + 0.5*sdev_lognormal_IO
	IO_VL_LB = avg_lognormal_IO + 1.5*sdev_lognormal_IO
	IO_VL_UB = math.Inf(1)
	IO_L_LB = avg_lognormal_IO + 0.5*sdev_lognormal_IO
	IO_L_UB = avg_lognormal_IO + 1.5*sdev_lognormal_IO
	IO_S_LB = avg_lognormal_IO - 1.5*sdev_lognormal_IO
	IO_S_UB = avg_lognormal_IO - 0.5*sdev_lognormal_IO
	IO_VS_LB = 0.0
	IO_VS_UB = avg_lognormal_IO - 1.5*sdev_lognormal_IO
	fmt.Println("lower bound an upper bound for all divisions in lognormal form")
	fmt.Print("\n")
	fmt.Println("lower bound an upper bound for calculation in lognormal form")
	fmt.Println(calc_VS_LB, calc_VS_UB)
	fmt.Println(calc_S_LB, calc_S_UB)
	fmt.Println(calc_M_LB, calc_M_UB)
	fmt.Println(calc_L_LB, calc_L_UB)
	fmt.Println(calc_VL_LB, calc_VL_UB)
	fmt.Print("\n")
	fmt.Println("lower bound an upper bound for data in lognormal form")
	fmt.Println(data_VS_LB, data_VS_UB)
	fmt.Println(data_S_LB, data_S_UB)
	fmt.Println(data_M_LB, data_M_UB)
	fmt.Println(data_L_LB, data_L_UB)
	fmt.Println(data_VL_LB, data_VL_UB)
	fmt.Print("\n")
	fmt.Println("lower bound an upper bound for I/O in lognormal form")
	fmt.Println(IO_VS_LB, IO_VS_UB)
	fmt.Println(IO_S_LB, IO_S_UB)
	fmt.Println(IO_M_LB, IO_M_UB)
	fmt.Println(IO_L_LB, IO_L_UB)
	fmt.Println(IO_VL_LB, IO_VL_UB)

	fmt.Println("lower bound an upper bound for all divisions in anti log form")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Println("lower bound an upper bound for calculation")
	fmt.Println(math.Pow(10, calc_VS_LB), math.Pow(10, calc_VS_UB))
	fmt.Println(math.Pow(10, calc_S_LB), math.Pow(10, calc_S_UB))
	fmt.Println(math.Pow(10, calc_M_LB), math.Pow(10, calc_M_UB))
	fmt.Println(math.Pow(10, calc_L_LB), math.Pow(10, calc_L_UB))
	fmt.Println(math.Pow(10, calc_VL_LB), math.Pow(10, calc_VL_UB))
	fmt.Print("\n")
	fmt.Println("lower bound an upper bound for data")
	fmt.Println(math.Pow(10, data_VS_LB), math.Pow(10, data_VS_UB))
	fmt.Println(math.Pow(10, data_S_LB), math.Pow(10, data_S_UB))
	fmt.Println(math.Pow(10, data_M_LB), math.Pow(10, data_M_UB))
	fmt.Println(math.Pow(10, data_L_LB), math.Pow(10, data_L_UB))
	fmt.Println(math.Pow(10, data_VL_LB), math.Pow(10, data_VL_UB))
	fmt.Print("\n")
	fmt.Println("lower bound an upper bound for IO")
	fmt.Println(math.Pow(10, IO_VS_LB), math.Pow(10, IO_VS_UB))
	fmt.Println(math.Pow(10, IO_S_LB), math.Pow(10, IO_S_UB))
	fmt.Println(math.Pow(10, IO_M_LB), math.Pow(10, IO_M_UB))
	fmt.Println(math.Pow(10, IO_L_LB), math.Pow(10, IO_L_UB))
	fmt.Println(math.Pow(10, IO_VL_LB), math.Pow(10, IO_VL_UB))

	fmt.Println("enter number of elememts for BASETABLE (base program loc table) HERE IT IS 2")
	fmt.Scan(&n2)
	if n2 <= 100 && n2 > 0 {
		for i = 0; i < n2; i++ {
			fmt.Scan(&BASEADD[i], &BASEADD_METHOD[i], &SIMILAR_BA[i])
		}

	}
	fmt.Print("\n")
	fmt.Println("PROXY SIZING FOR EACH OBJECT LOC(BA) ")
	fmt.Print("\n")
	for i = 0; i < n1; i++ {
		for k = 0; k < n2; k++ {
			if SIMILAR_BA[k] == identifier_1[i] {
				BA = BA + BASEADD_METHOD[k]*avh[i]
			}
			if SIMILAR_BA[k] == identifier_1[i] && type_1[i] == "Calculation" {
				SIMILAR_BA_loc_calc = BASEADD_METHOD[k] * avh[i]
				fmt.Println(" this base addition Belongs to Calculation")
				fmt.Println(SIMILAR_BA_loc_calc)
			}

			if SIMILAR_BA[k] == identifier_1[i] && type_1[i] == "Data" {
				SIMILAR_BA_loc_data = BASEADD_METHOD[k] * avh[i]
				fmt.Println(" this base addition Belongs to  Data")
				fmt.Println(SIMILAR_BA_loc_data)

			}

			if SIMILAR_BA[k] == identifier_1[i] && type_1[i] == "I/O" {
				SIMILAR_BA_loc_IO = BASEADD_METHOD[k] * avh[i]
				fmt.Println(" this base addition  Belongs to  I/O")
				fmt.Println(SIMILAR_BA_loc_IO)

			}
		}
	}

	fmt.Println("BASE ADDITION  LOC", BA)

	fmt.Println("enter number of elememts for NO_TABLE(new object table) HERE IT IS 6")
	fmt.Scan(&n3)
	if n3 <= 100 && n3 > 0 {
		for i = 0; i < n3; i++ {
			fmt.Scan(&NEWADD[i], &NEW_METHOD[i], &SIMILAR_NEW[i])
		}

	}
	fmt.Print("\n")
	fmt.Println("PROXY SIZING FOR EACH  NEW OBJECTS (NO)")
	fmt.Print("\n")
	for i = 0; i < n1; i++ {
		for k = 0; k < n3; k++ {
			if SIMILAR_NEW[k] == identifier_1[i] {
				NO = NO + NEW_METHOD[k]*avh[i]
			}
			if SIMILAR_NEW[k] == identifier_1[i] && type_1[i] == "Calculation" {
				SIMILAR_NEW_loc_calc = NEW_METHOD[k] * avh[i]
				fmt.Println("this new object belongs to  Calculation")
				fmt.Println(SIMILAR_NEW_loc_calc)
			}

			if SIMILAR_NEW[k] == identifier_1[i] && type_1[i] == "Data" {
				SIMILAR_NEW_loc_data = NEW_METHOD[k] * avh[i]
				fmt.Println("this new object belongs to Data")
				fmt.Println(SIMILAR_NEW_loc_data)

			}

			if SIMILAR_NEW[k] == identifier_1[i] && type_1[i] == "I/O" {
				SIMILAR_NEW_loc_IO = NEW_METHOD[k] * avh[i]
				fmt.Println(" this new object belongs to I/O")
				fmt.Println(SIMILAR_NEW_loc_IO)

			}
		}
	}

	fmt.Println(" NEW OBJECT  LOC  ", NO)
	E_LOC = NO + BA + M
	fmt.Println("total eastimated object LOC  ", E_LOC)

	fmt.Println("enter number of elememts of the table containing  actual and estimated of performance on the past project(HERE IT IS 20)")
	fmt.Scan(&n4)

	if n4 <= 100 && n4 > 0 {
		for i = 0; i < n4; i++ {
			fmt.Scan(&identifier_2[i], &ESTLOC[i], &ACLOC[i], &ESTDURATION[i], &ACTUALDURATION[i])
		}

	}
	cr = correlationCoefficient(ESTLOC, ACLOC, n4)

	fmt.Println("corelation coff for estimated loc and actualloc ", cr)
	if cr >= 0.7 {
		b_1 = get_b1(ESTLOC, ACLOC, n4)
		b_o = get_bo(ESTLOC, ACLOC, n4)
		fmt.Println("b1 for estimated loc and actual loc ", b_1)
		fmt.Println("b0 for estimated loc and actual loc ", b_o)
		N = ((E_LOC * b_1) + b_o)

	}

	fmt.Println("estimated new and changed loc ", N)

	H = N + B - D - M

	fmt.Println("estimated Total Loc", H)
	fmt.Print("\n")

	fmt.Println("DURATION ESTIMATE")
	fmt.Print("\n")

	productivity = (sum(ESTLOC) / sum(ESTDURATION)) * 60

	fmt.Println("planned LOC/HR productivity ", productivity)

	ca = correlationCoefficient(ESTLOC, ESTDURATION, n4)
	fmt.Println("corelation coff for estimated loc and estimated duration ", ca)
	if ca > 0.7 {
		b_1 = get_b1(ESTLOC, ESTDURATION, n4)
		b_o = get_bo(ESTLOC, ESTDURATION, n4)
		duration = ((E_LOC * b_1) + b_o)

	}

	if ca < 0.7 {
		b_1 = get_b1(ESTLOC, ACTUALDURATION, n4)
		b_o = get_bo(ESTLOC, ACTUALDURATION, n4)
		duration = ((E_LOC * b_1) + b_o)

	}

	fmt.Println("estimated project duration ", duration)

	fmt.Println("70% predected Range size is calculated is ", RSIZE)
	fmt.Println("70% predected Range Duration is calculated is ", RDURATION)

}
