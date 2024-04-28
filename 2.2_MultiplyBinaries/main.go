package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Introduction to Software Engineering, Labwork #2.2\nVariant 3")
	fmt.Println("Task: Simulate the algorithm for multiplying numbers in fixed-point format after the lowest bit according to scheme #3 with a bit depth of 8 bits represented in direct code")
	var bin1, bin2 string

tryAgain:
	fmt.Printf("Enter the first 8-bit binary: ")
	fmt.Scan(&bin1)
	fmt.Printf("Enter the second 8-bit binary: ")
	fmt.Scan(&bin2)

	if len(bin1) != 8 || len(bin2) != 8 {
		fmt.Println("Please enter the proper length number")
		goto tryAgain
	} else {
		fmt.Printf("\n%s x %s = %s\n\n", bin1, bin2, multiplyTwoBinaries(bin1, bin2))
		fmt.Printf("%d x %d = %d\n", convertBinaryToDecimal(bin1), convertBinaryToDecimal(bin2), convertBinaryToDecimal(multiplyTwoBinaries(bin1, bin2)))
	}

}

func multiplyTwoBinaries(bin1, bin2 string) string {
	result := "0"
	for i := len(bin2) - 1; i >= 0; i-- {
		if bin2[i] == '1' {
			result = getSumOfTwoBinaries(result, bin1)
		}
		bin1 += "0"
	}
	return result
}

func getSumOfTwoBinaries(bin1, bin2 string) string {
	result := ""
	sum := 0
	i, j := len(bin1)-1, len(bin2)-1

	for i >= 0 || j >= 0 || sum == 1 {
		if i >= 0 {
			sum += int(bin1[i] - '0')
		}
		if j >= 0 {
			sum += int(bin2[j] - '0')
		}
		result = string(sum%2+'0') + result
		sum /= 2
		i--
		j--
	}

	return result
}

func convertBinaryToDecimal(bin string) int {
	decimal := 0
	power := len(bin) - 1
	for _, digit := range bin {
		if digit == '1' {
			decimal += int(math.Pow(2, float64(power)))
		}
		power--
	}
	return decimal
}
