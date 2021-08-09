package util

// 对照着小灰的漫画算法做的（原文用的是java实现）
var words []int64

var size int

// 获取bitmap某一位所对于的word
func getWordIndex(bitIndex int) int {
	return bitIndex >> 6
}

func setBit(bitIndex int) {
	if bitIndex < 0 || bitIndex > size-1 {
		panic("123")
	}
	wordIndex := getWordIndex(bitIndex)
	words[wordIndex] |= 1 << bitIndex
}

func getBit(bitIndex int) bool {
	if bitIndex < 0 || bitIndex > size-1 {
		panic("123")
	}
	wordIndex := getWordIndex(bitIndex)
	return (words[wordIndex] & (1 << bitIndex)) != 0
}

func NewBitMap(s int) {
	size = s
	words = make([]int64, getWordIndex(size-1)+1)
}
