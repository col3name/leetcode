package hyperloglog

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func getRandomData() (out [][]byte, intout []uint32) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < math.MaxInt16*1000; i++ {
		i := rand.Uint32()
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, i)
		out = append(out, b)
		intout = append(intout, i)
	}
	return
}

func classicCountDistinct(input []uint32) int {
	m := map[uint32]struct{}{}
	for _, i := range input {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
		}
	}
	return len(m)
}

func TestName(t *testing.T) {
	data, intout := getRandomData()
	start := time.Now()
	dt := classicCountDistinct(intout)
	elapsedClassic := time.Since(start)
	start = time.Now()
	h := NewHyperLogLog(64)
	for _, b := range data {
		h.Add(b)
	}
	hd := h.Count()
	elapsedHyperLogLog := time.Since(start)
	nanoseconds := elapsedClassic.Nanoseconds()
	i := elapsedHyperLogLog.Nanoseconds()
	fmt.Println(float64(nanoseconds) / float64(i))
	fmt.Println("class time", nanoseconds, "ns", elapsedClassic.String())
	fmt.Println("hyperloglog time", i, "ns", elapsedHyperLogLog.String())
	fmt.Printf("classic estimate: %v\n", dt)
	fmt.Printf("hyperloglog estimate: %v\n", hd)
	fmt.Printf("percentage missed: %.2f\n", 100.-(float64(hd)/float64(dt))*100)
}
