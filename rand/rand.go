package rand

import (
	"math/rand"
	"time"
)

const (
	NUM_CHARS   = "0123456789"
	LOWER_CHARS = "abcdefghikjlmnopqrstuvwxyz"
	UPPER_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func String(n uint) string {
	chars := NUM_CHARS + LOWER_CHARS + UPPER_CHARS
	return strFromChars(n, chars)
}

func LowerString(n uint) string {
	chars := NUM_CHARS + LOWER_CHARS
	return strFromChars(n, chars)
}

// TimeString() 生成n位的由时间和随机数连成的字符串
func TimeString(n uint) string {
	timeStr := time.Now().Format("20060102150405")
	return timeStr + NumString(n-14)
}

func NumString(n uint) string {
	return strFromChars(n, NUM_CHARS)
}

func strFromChars(n uint, chars string) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = chars[rand.Int63()%int64(len(chars))]
	}
	return string(s)
}

func IntRange(min, max int) (num int) {
	return rand.Intn(max-min+1) + min
}

func ShuffleStringsN(s []string, n int) (res []string) {
	res = make([]string, 0)
	randIdx := rand.Perm(len(s))

	max := 0
	if len(s) < n {
		max = len(s)
	} else {
		max = n
	}

	for i := 0; i < max; i++ {
		res = append(res, s[randIdx[i]])
	}
	return
}

func Jackpot(prob float64) bool {
	return rand.Float64() < prob
}

func IntInSlice(s []int) int {
	randIndex := rand.Intn(len(s))
	return s[randIndex]
}

func Uint32InSlice(s []uint32) uint32 {
	randIndex := rand.Intn(len(s))
	return s[randIndex]
}

// s 包含从大到小的概率, 根据概率随机返回 s 的索引
func IndexByProb(s []int) int {
	l := len(s)
	r := rand.Intn(100)
	down := 0
	up := 0
	for i := l - 1; i >= 0; i-- {
		up = down + s[i]
		if r >= down && r < up {
			return i
		}
		down += s[i]
	}
	return 0
}
