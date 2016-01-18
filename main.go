package main

import "fmt"
import "math"
import "time"

func print_table(t [][]float64) {
	for _, j := range t {
		fmt.Println(j)
	}
}

func dynamic(string1, string2 string) float64 {
	x := len(string1) + 1
	y := len(string2) + 1
	table := make([][]float64, x)
	for i := 0; i < x; i++ {
		table[i] = make([]float64, y)
	}

	// Step 1: set up the first row
	for i, row := range table {
		row[0] = float64(i)
	}
	// Step 2: set up the first column
	for i, _ := range table[0] {
		table[0][i] = float64(i)
	}

	// Outer loop loops over rows
	for j := 1; j < len(table); j++ { // row index
		//Inner loop loops over columns
		for i := 1; i < len(table[j]); i++ {
			char1 := string1[j-1]
			char2 := string2[i-1]

			mismatch_score := 1.0
			if char1 == char2 {
				mismatch_score = 0.0
			}

			diag := table[j-1][i-1]
			left := table[j][i-1]
			up := table[j-1][i]

			table[j][i] = math.Min(math.Min(mismatch_score+diag, up+1.0), left+1.0)

		}
	}
	fmt.Println("Edit distance table:")
	print_table(table)
	return table[x-1][y-1]
}

func recursive(string1, string2 string) float64 {
	if len(string1) == 0 {
		return float64(len(string2))
	}

	if len(string2) == 0 {
		return float64(len(string1))
	}

	if (len(string2) == 0) && (len(string1) == 0) {
		return 0.0
	}

	mismatch_1 := recursive(string1, string2[0:len(string2)-1]) + 1.0
	mismatch_2 := recursive(string1[0:len(string1)-1], string2) + 1.0

	var cost float64
	if string1[len(string1)-1] == string2[len(string2)-1] {
		cost = 0.0
	} else {
		cost = 1.0
	}
	mismatch_3 := recursive(string1[0:len(string1)-1], string2[0:len(string2)-1]) + cost

	return math.Min(math.Min(mismatch_1, mismatch_2), mismatch_3)
}

func main() {
	// create two example strings, calculate the edit distance between them
	str1 := "exciting"
	str2 := "executed"
	t0 := time.Now()
	fmt.Println("Edit distance, recursive algorithm:", recursive(str1, str2))
	fmt.Println("Recursive algorithm execution time:", time.Since(t0))
	t1 := time.Now()
	fmt.Println("Edit distance, dynamic programming algorithm:", dynamic(str1, str2))
	fmt.Println("Dynamic programming execution time: ", time.Since(t1))
}
