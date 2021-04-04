package main
import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"strconv"
	
)
func readOrders(name string) [][]string {

	f, err := os.Open(name)

	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','	
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}
func main() {
	rows := readOrders("record.csv")
	for _, row := range rows {
        for _, val := range row {
            fmt.Print(val)
        }
		fmt.Println("\n")
    }
	for i := range rows {
	strconv.ParseFloat(rows[i][4], 8)
	}
	
	
}