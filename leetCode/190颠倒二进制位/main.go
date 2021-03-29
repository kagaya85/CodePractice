package main

// 颠倒给定的 32 位无符号整数的二进制位。

const (
	m1 = 0x55555555 // 01010101010101010101010101010101
	m2 = 0x33333333 // 00110011001100110011001100110011
	m4 = 0x0f0f0f0f // 00001111000011110000111100001111
	m8 = 0x00ff00ff // 00000000111111110000000011111111
)

func reverseBits(nums uint32) uint32 {
	nums = nums>>1&m1 | nums&m1<<1
	nums = nums>>2&m2 | nums&m2<<2
	nums = nums>>4&m4 | nums&m4<<4
	nums = nums>>8&m8 | nums&m8<<8
	return nums>>16 | nums<<16
}
