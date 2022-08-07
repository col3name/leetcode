package main

import (
	"flag"
	"fmt"
	"github.com/col3name/algo/cmd/performacnce/pointer-vs-stack/model"
	"github.com/col3name/algo/cmd/performacnce/pointer-vs-stack/stat"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func main() {
	count := flag.Int64("count", 100_000, "-count=23")
	flag.Parse()

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	start := time.Now()
	var i int64
	i = 0
	result := make([]*model.Gopher, *count)
	for i < *count {
		result[i] = &model.Gopher{
			Id:        i,
			Name:      "Gopher" + strconv.FormatInt(i, 10),
			Path:      "usr/bin/gopher" + strconv.FormatInt(i, 10),
			StackSize: rand.Intn(10000)}
		i++
		if i%10000 == 0 {
			fmt.Println(i)
		}
		//PrintMemUsage()
	}

	stat.PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	stat.PrintMemUsage()
	fmt.Println(time.Since(start).String())
}

