package main

import (
	"github.com/col3name/algo/cmd/performacnce/pointer-vs-stack/stat"
	"runtime"
	"time"
)

func main() {
	// Below is an example of using our PrintMemUsage() function
	// Print our starting memory usage (should be around 0mb)
	stat.PrintMemUsage()

	var overall [][]int
	for i := 0; true; i++ {

		// Allocate memory using make() and append to overall (so it doesn't get
		// garbage collected). This is to create an ever increasing memory usage
		// which we can track. We're just using []int as an example.
		a := make([]int, 0, 99_999_999)
		overall = append(overall, a)

		// Print our memory usage at each interval
		stat.PrintMemUsage()
		time.Sleep(time.Second)
	}

	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	overall = nil
	stat.PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	stat.PrintMemUsage()
}
