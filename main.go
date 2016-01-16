// go vs python vs C++ comparison?
package main

import "fmt"
import "math"

func recursive (string1, string2 string) float64{
	if (len(string1) == 0) {
		return float64(len(string2))
	}

	if (len(string2) == 0) {
		return float64(len(string1))
	}

	if (len(string2) == 0) &&  (len(string1) == 0) {
		return 0.0
	}
	
	mismatch_1 := recursive(string1, string2[0:len(string2) -1]) + 1.0
	mismatch_2 := recursive(string1[0:len(string1) -1], string2 ) + 1.0

	var cost float64
	if string1[len(string1)-1] == string2[len(string2)-1] {
		cost = 0.0
	} else {
		cost = 1.0
	}
	mismatch_3 := recursive(string1[0:len(string1) -1], string2[0:len(string2) -1]) + cost

	return math.Min(math.Min(mismatch_1, mismatch_2), mismatch_3)
}

func main(){

	str1 := "shakespeare"
	str2 := "shakespare"
	fmt.Println(recursive(str1, str2))
}
