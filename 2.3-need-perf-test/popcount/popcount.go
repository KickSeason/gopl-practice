package popcount

var pc [256] byte

func init() {
	for i := range pc {
		pc[i] = pc[i / 2] + byte(i & 1)
	}
}

// func PopCount(x uint64) int {
// 	return int(pc[byte(x >> (0 * 8))] + 
// 			pc[byte(x >> (1 * 8))] + 
// 			pc[byte(x >> (2 * 8))] + 
// 			pc[byte(x >> (3 * 8))] + 
// 			pc[byte(x >> (4 * 8))] + 
// 			pc[byte(x >> (5 * 8))] + 
// 			pc[byte(x >> (6 * 8))] + 
// 			pc[byte(x >> (7 * 8))])
// }

func PopCount(x uint64) int {
	var popcount int
	for i := uint(0); i < 8; i++ {
		popcount += int(pc[x >> (i * 8)])
	}
	return popcount
 }

func PopCountBit(x uint64) int {
	var popcount int
	for  0 < x {
		if 0 < (x & 1) {
			popcount++
		}
		x >>= 1
	}
	return popcount
}

func PopCountAnd(x uint64) int {
	var popcount int
	for  0 < x {
		x = x & (x - 1)
		popcount++
	}
	return popcount
}