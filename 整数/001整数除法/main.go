package main

import "math"

// 第一种能想到的解法，最后超时了
func divideNormal(a int, b int) int {
	newa, newb := a, b
	flaga, flagb := 1, 1
	if a < 0 {
		newa = -a
		flaga = -1
	}

	if b < 0 {
		newb = -b
		flagb = -1
	}

	cnt := 0
	for newa-newb >= 0 {
		newa = newa - newb
		cnt++
	}

	if cnt > int(math.Pow(2, 31)-1) {
		cnt = int(math.Pow(2, 31) - 1)
	}

	return cnt * flaga * flagb
}

// 第二种方法
// 快速加
// 如果b是偶数，a * b = a * b / 2 + a * b / 2
// 如果b是奇数，a * b = a * b / 2 + a * b / 2 + a
func qadRecursion(a, b int) int {
	if b == 0 {
		return 0
	}
	t := qadRecursion(a, b/2)
	if b&1 > 0 {
		// 说明是奇数
		return t + t + a
	}
	return t + t
}

func qad(a, b int) int {
	res := 0
	for b > 0 {
		if b&1 > 0 {
			res += a
		}
		a += a
		b >>= 1
	}
	return res
}
