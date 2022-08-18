package hyperloglog

import (
	"hash/fnv"
	"math"
	"math/bits"
)

type HyperLogLog struct {
	registers []int
	m         uint // number of registers
	b         uint // bits to calculate [4..16]
}

func NewHyperLogLog(m uint) HyperLogLog {
	return HyperLogLog{
		registers: make([]int, m),
		m:         m,
		b:         uint(math.Ceil(math.Log2(float64(m)))),
	}
}

func (h HyperLogLog) Add(data []byte) HyperLogLog {
	x := createHash(data)
	k := 32 - h.b // first b bits
	r := leftmostActiveBit(x << h.b)
	j := x >> k

	if r > h.registers[j] {
		h.registers[j] = r
	}
	return h
}

func (h HyperLogLog) Count() uint64 {
	sum := 0.
	m := float64(h.m)
	for _, v := range h.registers {
		sum += math.Pow(math.Pow(2, float64(v)), -1)
	}
	estimate := .79402 * m * m / sum
	return uint64(estimate)
}

func leftmostActiveBit(x uint32) int {
	return 1 + bits.LeadingZeros32(x)
}

// create a 32-bit hash
func createHash(stream []byte) uint32 {
	h := fnv.New32()
	_, _ = h.Write(stream)
	sum := h.Sum32()
	h.Reset()
	return sum
}
