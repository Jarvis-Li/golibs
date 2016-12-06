package math

import (
	"math"
)

const (
	UINT_MAX = ^uint(0)
	UINT_MIN = 0
	INT_MAX  = int(UINT_MAX >> 1)
	INT_MIN  = -INT_MAX - 1
)

// ExponentIncr() 返回按指数增长的第n次的数值，n从1开始，n为1时返回初始值
func ExponentIncr(base float64, incr float64, n uint32) (res float64) {
	if n == 0 {
		return
	}
	res = base * math.Pow(1+incr, float64(n-1))
	return
}

// 计算等比数列第n项
func ExponentN(a1, q, n float64) float64 {
	return a1 * math.Pow(q, n-1)
}

// 计算等比数列前n项和
func ExponentSum(a1, q, n float64) float64 {
	return a1 * (1 - math.Pow(q, n)) / (1 - q)
}

func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func PointYAtLine(x1, y1, x2, y2, x3 float64) (y3 float64) {
	return y1 + (y2-y1)/(x2-x1)*(x3-x1)
}

func RoundMax(a float64) float64 {
	factor := 10.0
	allFactor := 1.0

	for {
		if a < 10 {
			break
		}
		a /= factor
		allFactor *= factor
	}

	return math.Ceil(a) * allFactor
}
