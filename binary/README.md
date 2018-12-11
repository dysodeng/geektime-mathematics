## 第一课：二进制，不了解计算机的源头，你学什么编程？

##### 什么是二进制
- 10进制表示以10为基数进位的计数法，10进制以10^n进位
- 2进制则表示以2为基数的进位计数法，2进制以2^n进位

##### 二进制与十进制相互转换
- 10进制转2进制：除2取余法
```
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
```
- 2进制转10进制：按权展开法
```
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
```

##### 计算机为什么使用二进制
- 计算机逻辑电路"开"与"关"可用二进制 1和0表示
- 二进制非常适合逻辑运算，"真"与"假"对于 1和0

##### 二进制的位操作
- 向左移位：二进制左移一位，就是将数字翻倍，即 x*2^n，x为被操作数字，n代表左移位数
- 向右移位：二进制右移一位，就是将数字除以2并求整数商，即 x/2^n，x为被操作数字，n代表右移位数
- 右移分为逻辑右移与算术右移：由于数字存在符号位，所以逻辑右移就是把符号位也向右移一位，然后补0，所以负数会变正数；算术右移则符号位不变，符号位后面一位开始右移并补0
- 或：参与操作的位中只要有一个位是1，最终结果就是1，即 两个为0，则为0，否则为1
- 与：参与操作的位中全部为1，最终结果才为1，否则为0，即 两个为1，则为1，否则为0
- 异或：参与操作的位相同，最终结果为0，否则为1，即 相同为0，不同为1
