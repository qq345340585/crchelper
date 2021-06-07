package crchelper

import (
	"strconv"

	"github.com/astaxie/beego/logs"
)

func CheckSum(data []byte, polystr string, initstr string, xoroutstr string, refin bool, refout bool, width uint64) (uint64, int) {
	var (
		crc_reg uint64
		poly    uint64
		xorout  uint64
		err     error
	)
	if crc_reg, err = strconv.ParseUint(initstr, 16, int(width)); err != nil {
		return 0, 1
	}
	if xorout, err = strconv.ParseUint(xoroutstr, 16, int(width)); err != nil {
		return 0, 1
	}
	if poly, err = strconv.ParseUint(polystr, 16, int(width)); err != nil {
		return 0, 1
	}

	if refin {
		poly = reverse_poly(poly, width)
	}
	var (
		j int
	)
	pw := uint64(pow(2, width-1))
	for _, tt := range data {
		if !refin {
			crc_reg ^= (uint64(tt) << (width - 8)) & uint64(pow(2, width)-1)
		} else {
			crc_reg ^= uint64(tt)
		}
		j++
		for a := 0; a < 8; a++ {
			if !refin {
				if crc_reg&pw > 0 {
					crc_reg = (crc_reg << 1) ^ poly
				} else {
					crc_reg = (crc_reg << 1)
				}
			} else {
				if crc_reg&1 == 1 {
					crc_reg = (crc_reg >> 1) ^ poly
				} else {
					crc_reg = (crc_reg >> 1)
				}
			}
		}
	}
	crc_reg ^= xorout
	crc_reg &= uint64(pow(2, width) - 1)
	if refin && !refout {
		crc_reg = reverse_poly(crc_reg, width)
	} else if !refin && refout {
		crc_reg = reverse_poly(crc_reg, width)
	}
	return crc_reg, 0
}

// 反转poly
func reverse_poly(poly uint64, width uint64) uint64 {
	num := Str2DEC(reverseString(convertToBin(poly, width)))
	logs.Debug("polyIn:0x%0x,numOut:0x%0x", poly, num)
	return num
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//str转十进制
func Str2DEC(s string) (num uint64) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (uint64(s[l-i-1]) - 48) << uint8(i)
	}
	return
}

// 将十进制数字转化为二进制字符串
func convertToBin(num uint64, width uint64) string {
	s := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.FormatUint(lsb, 10) + s
	}
	length := uint64(len(s))
	for {
		if length < width {
			s = "0" + s
			length++
		} else {
			break
		}
	}
	return s
}

//算次方
func pow(x float64, n uint64) float64 {
	if x == 0 {
		return 0
	}
	result := calPow(x, n)
	if n < 0 {
		result = 1 / result
	}
	return result
}

//算次方
func calPow(x float64, n uint64) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	// 向右移动一位
	result := calPow(x, n>>1)
	result *= result

	// 如果n是奇数
	if n&1 == 1 {
		result *= x
	}

	return result
}
