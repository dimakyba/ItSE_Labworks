package main

import (
	"fmt"
	"math"
)

func main() {
	// 	  Варіант 7.

	// a) Переведення дробових чисел 10 → 2
	// b) Перетворення цілих чисел з прямого коду в додатковий при довжині розрядної сітки 8.

	fmt.Println("Introduction to Software Engineering, Labwork #2\nVariant 7")

tryAgain:
	fmt.Println("Choose the case:\n1) convert a decimal fraction to binary;\n2) convert an integer from a direct code to an additional code with a bit grid length of 8.")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Enter the decimal fraction: ")
		var frac float64
		fmt.Scan(&frac)
		fmt.Println(convertDecimalFracToBinary(frac))
		goto tryAgain
	case 2:
		fmt.Printf("Enter a direct code integer in binary with the length of 8. (for example: 1|0100110): ")
		var direct string
		fmt.Scan(&direct)
		fmt.Println(convertDirectIntoAdditional(direct))
		goto tryAgain
	case 0:
		break
	}

}

func convertDecimalFracToBinary(frac float64) string {
	var temp int
	precision := 8
	result := "0."
	for i := 0; i < precision; i++ {
		frac *= 2
		temp = int(math.Floor(float64(frac)))
		if temp > 1 {
			temp--
		}
		if frac >= 1 {
			frac--
		}
		result += fmt.Sprint(temp)
	}
	result += " (2)\n"

	return result
}

func convertDirectIntoAdditional(direct string) string {
	length := len(direct)
	additional := []byte(direct[2:])

	if direct[0] == '0' {
		return direct
	} else if direct[0] == '1' {
		for i := 0; i < length-2; i++ {
			if additional[i] == '0' {
				additional[i] = '1'
			} else if additional[i] == '1' {
				additional[i] = '0'
			}
		}
		additional = addOneToBinary(additional)
	}

	return "1|" + string(additional)
}

func addOneToBinary(binary []byte) []byte {
	carry := 1
	for i := len(binary) - 1; i >= 0; i-- {
		if carry == 0 {
			break
		}
		sum := int(binary[i]-'0') + carry
		binary[i] = byte(sum%2) + '0'
		carry = sum / 2
	}
	if carry > 0 {
		binary = append([]byte{'1'}, binary...)
	}
	return binary
}
