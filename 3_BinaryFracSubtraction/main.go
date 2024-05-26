package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var separator = []rune{' ', '*', '^'}

func main() {
	fmt.Print("\n1. Введіть дріб з плаваючою крапкою (наприклад, -0.101 * 2^101): ")
	// var input1 string
	// fmt.Scanln(&input1)
	input1 := getInput()
	fmt.Print("\n2. Введіть дріб з плаваючою крапкою (наприклад, -0.101 * 2^101): ")
	// var input2 string
	// fmt.Scanln(&input2)
	input2 := getInput()

	fmt.Println(CalculateFloatingPointDifference(input1, input2))
}

func CalculateFloatingPointDifference(input1, input2 string) string {
	a_mnt_sign, a_ord_sign := ParseInput(input1)
	b_mnt_sign, b_ord_sign := ParseInput(input2)

	fmt.Printf("\nA мантиса (двійкова): %v = %v%v\n", a_mnt_sign, bitToSign(a_mnt_sign), Float64DecimalSigned(a_mnt_sign))
	fmt.Printf("A період (двійкова): %v = %v%v\n", a_ord_sign, bitToSign(a_ord_sign), IntDecimalSigned(a_ord_sign))
	fmt.Printf("B мантиса (двійкова): %v = %v%v\n", b_mnt_sign, bitToSign(b_mnt_sign), Float64DecimalSigned(b_mnt_sign))
	fmt.Printf("B період (двійкова): %v = %v%v\n", b_ord_sign, bitToSign(b_ord_sign), IntDecimalSigned(b_ord_sign))

	fmt.Printf("\nA = %s%s\nB = %s%s\n", a_mnt_sign, a_ord_sign, b_mnt_sign, b_ord_sign)

	ngtv_b_ord := InvertSign(b_ord_sign)
	a_ord_supp := IntToSupp(a_ord_sign)
	ngtv_b_ord_supp := IntToSupp(ngtv_b_ord)

	fmt.Printf("\nA період (додатковий код): %s = %d\n", a_ord_supp, IntDecimalSupp(a_ord_supp))
	fmt.Printf("від'ємний B період (прямий код): %s\n", ngtv_b_ord)
	fmt.Printf("від'ємний B період (додатковий код): %s = %d\n", ngtv_b_ord_supp, IntDecimalSupp(ngtv_b_ord_supp))

	ord_diff_supp := GetSumOfTwoBinaries(a_ord_supp, ngtv_b_ord_supp)
	ord_diff_dec := int(math.Abs(float64(IntDecimalSupp(ord_diff_supp))))

	fmt.Printf("\nРізниця періодів: |%s - %s| = |%s + %s| = %s (%d)\n", a_ord_sign, b_ord_sign, a_ord_supp, ngtv_b_ord_supp, ord_diff_supp, ord_diff_dec)

	ord_sign := a_ord_sign
	if ord_diff_dec != 0 {
		if ord_diff_supp[0] == '0' {
			b_mnt_sign = fmt.Sprintf("%c", b_mnt_sign[0]) + ShiftRight(b_mnt_sign[1:], ord_diff_dec)
			ord_sign = b_ord_sign
		} else {
			a_mnt_sign = fmt.Sprintf("%c", a_mnt_sign[0]) + ShiftRight(a_mnt_sign[1:], ord_diff_dec)
			ord_sign = a_ord_sign
		}
	}

	fmt.Printf("\nЗсунуті значення:\nA мантиса (прямий код): %s = %s%v\n", a_mnt_sign, bitToSign(a_mnt_sign), Float64DecimalSigned(a_mnt_sign))
	fmt.Printf("B мантиса (прямий код): %s = %s%v\n", b_mnt_sign, bitToSign(b_mnt_sign), Float64DecimalSigned(b_mnt_sign))

	ngtv_b_mnt_sign := InvertSign(b_mnt_sign)
	a_mnt_supp := FloatToSupp(a_mnt_sign)
	ngtv_b_mnt_supp := FloatToSupp(ngtv_b_mnt_sign)

	fmt.Printf("\nA - B = A_дод + (-B)_дод:\nA мантиса (додатковий код): %s = %v\n", a_mnt_supp, Float64DecimalSupp(a_mnt_supp))
	fmt.Printf("від'ємна B мантиса (прямий код): %s = %s%v\n", ngtv_b_mnt_sign, func() string {
		if ngtv_b_mnt_sign[0] == '1' {
			return "-"
		}
		return ""
	}(), Float64DecimalSupp(ngtv_b_mnt_sign))
	fmt.Printf("від'ємна B мантиса (додатковий код): %s  = %v\n", ngtv_b_mnt_supp, Float64DecimalSupp(ngtv_b_mnt_supp))

	res := GetSumOfTwoBinaries(a_mnt_supp, ngtv_b_mnt_supp)

	fmt.Printf("\nA_n + (-B_n) = %s (додатковий код)\n", res)

	res = toSigned(res)
	// fmt.Sprintf()

	fmt.Printf("\nA - B = %s|%s %s|%s (прямий код) = %v\n", fmt.Sprintf("%c", res[0]), res[1:], string(ord_sign[0]), ord_sign[1:], Float64DecimalSupp(res))
	fmt.Printf("\nC = %s%s\n", res, a_ord_sign)

	leadingZeros := 0
	for i := 1; i < len(res)-1 && res[i] == '0'; i++ {
		leadingZeros++
	}

	exponent := IntDecimalSigned(ord_sign) - leadingZeros
	adjustedExponent := IntToSupp(intToBinary(exponent))

	trimmedMantissa := res[:1] + res[1+leadingZeros:]

	fmt.Printf("\nC = %s|%s %s|%s (прямий код) = %v\n", string(trimmedMantissa[0]), trimmedMantissa[1:], string(adjustedExponent[0]), adjustedExponent[1:], Float64DecimalSupp(trimmedMantissa))

	return fmt.Sprintf("\nC = %s%s", trimmedMantissa, adjustedExponent)
}

func ParseInput(i string) (mStr, oStr string) {
	parts := strings.FieldsFunc(i, func(r rune) bool {
		for _, sep := range separator {
			if r == sep {
				return true
			}
		}
		return false
	})

	part1 := parts[0]
	part2 := parts[2]

	if part1[0] == '-' {
		mStr = "1" + part1[3:]
	} else {
		mStr = "0" + part1[2:]
	}

	if part2[0] == '-' {
		oStr = "1" + part2[1:]
	} else {
		oStr = "0" + part2
	}

	return mStr, oStr
}

func ShiftRight(n string, order int) string {
	shift := strings.Repeat("0", order)
	return shift + n[:len(n)-order]
}

func GetSumOfTwoBinaries(bin1, bin2 string) string {
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
		result = fmt.Sprintf("%d", sum%2+'0') + result
		sum /= 2
		i--
		j--
	}

	return result
}

func IntDecimalSigned(bin string) int {
	dec := 0
	length := len(bin) - 1

	for i := 1; i <= length; i++ {
		if bin[i] == '1' {
			dec += int(math.Pow(2, float64(length-i)))
		}
	}
	return dec
}

func Float64DecimalSigned(bin string) float64 {
	dec := 0.0
	length := len(bin) - 1

	for i := 1; i <= length; i++ {
		if bin[i] == '1' {
			dec += math.Pow(2, float64(-i))
		}
	}
	return dec
}

func IntDecimalSupp(bin string) int {
	len := len(bin) - 1
	dec := 0

	for i := 1; i <= len; i++ {
		if bin[i] == '1' {
			dec += int(math.Pow(2, float64(len-i)))
		}
	}

	if bin[0] == '1' {
		dec -= int(math.Pow(2, float64(len)))
	}

	return dec
}

func Float64DecimalSupp(bin string) float64 {
	len := len(bin) - 1
	dec := 0.0

	for i := 1; i <= len; i++ {
		if bin[i] == '1' {
			dec += math.Pow(2, float64(-(i + 1)))
		}
	}

	if bin[0] == '1' {
		dec -= math.Pow(2, -1)
	}

	return dec
}

func IntToSupp(num_bin string) string {
	if num_bin[0] == '0' {
		return num_bin
	}

	bin := []rune(num_bin)
	i := len(bin) - 1

	for i > 0 {
		if bin[i] == '1' {
			break
		}
		i--
	}

	for i > 0 {
		if bin[i] == '1' {
			bin[i] = '0'
		} else {
			bin[i] = '1'
		}
		i--
	}

	return string(bin)
}

func FloatToSupp(num_bin string) string {
	bin := []rune(num_bin)

	if bin[0] == '0' {
		for j := 0; j < len(bin)-1; j++ {
			bin[j] = bin[j+1]
		}
		bin[len(bin)-1] = '0'
		return string(bin)
	}

	i := len(bin) - 1
	for i >= 0 {
		if bin[i] == '1' {
			break
		}
		i--
	}

	for i >= 0 {
		if bin[i] == '1' {
			bin[i] = '0'
		} else {
			bin[i] = '1'
		}
		i--
	}

	for j := 0; j < len(bin)-1; j++ {
		bin[j] = bin[j+1]
	}
	bin[len(bin)-1] = '0'

	return fmt.Sprintf("%d", bin)
}

func toSigned(num_bin string) string {
	bin := []rune(num_bin)

	if bin[0] == '1' {
		i := len(bin) - 1
		for i >= 0 && bin[i] != '1' {
			i--
		}
		i--
		for i >= 0 {
			if bin[i] == '1' {
				bin[i] = '0'
			} else {
				bin[i] = '1'
			}
			i--
		}
	}

	return fmt.Sprintf("%d", bin)
}

func InvertSign(bin string) string {
	if bin[0] == '0' {
		return "1" + bin[1:]
	} else {
		return "0" + bin[1:]
	}
}

func bitToSign(bin string) string {
	if bin[0] == '1' {
		return "-"
	}
	return ""
}

func intToBinary(value int) string {
	if value == 0 {
		return "0"
	}

	var binary string

	isNegative := value < 0
	value = int(math.Abs(float64(value)))

	for value > 0 {
		binary = fmt.Sprintf("%d", value%2) + binary
		value /= 2
	}

	if isNegative {
		binary = "1" + binary
		return FloatToSupp(binary)
	}
	return binary
}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
