package binary

import (
	"strconv"
	"fmt"
	"math"
)

// 10进制转2进制
func DecimalToBinary(decimal int) string {
	binary := ""
	var remainder []int

	for  {
		if decimal <= 0 {
			break
		}
		mod := decimal % 2
		decimal = decimal / 2
		remainder = append(remainder, mod)
	}

	l := len(remainder) - 1
	for i:=l; i >= 0; i-- {
		binary += strconv.Itoa(remainder[i])
	}

	return binary
}

// 2进制转10进制
func BinaryToDecimal(binary string) int {

	binaryLen := len(binary)
	decimal := 0

	// 反转二进制字符串
	runes := []rune(binary)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	binary = string(runes)

	for i := 0; i < binaryLen; i++ {
		numString := string(binary[i])
		num, err := strconv.Atoi(numString)
		if err != nil {
			num = 0
		}

		decimal += int(float64(num) * math.Pow(2, float64(i)))

	}

	return decimal
}

func Run() {
	decimal := 1234
	binary := DecimalToBinary(decimal)
	fmt.Println("10进制转2进制：", decimal, "<--->", binary) // 1100100
	fmt.Println("2进制转10进制：", binary, "<--->", BinaryToDecimal(binary)) // 1100100
}
