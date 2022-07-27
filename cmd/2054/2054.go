package _054

import (
	"fmt"
	"sort"
)

type event struct {
	startTime int
	endTime   int
	value     int
}

type EventsSortable []event

func (a EventsSortable) Len() int           { return len(a) }
func (a EventsSortable) Less(i, j int) bool { return a[i].startTime < a[j].startTime }
func (a EventsSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type EventsSortableEnd []event

func (a EventsSortableEnd) Len() int           { return len(a) }
func (a EventsSortableEnd) Less(i, j int) bool { return a[i].endTime < a[j].endTime }
func (a EventsSortableEnd) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func maxTwoEvents(events_array [][]int) int {
	events := make([]event, len(events_array))
	for i := 0; i < len(events_array); i++ {
		tmp_event := event{}
		tmp_event.startTime = events_array[i][0]
		tmp_event.endTime = events_array[i][1]
		tmp_event.value = events_array[i][2]
		events[i] = tmp_event
	}

	return getMaxVersionFour(events)
	// return getMaxVersionThree(events)
}

func getMaxVersionFour(right []event) int {
	n := len(right)
	left := make([]event, n)
	for i := 0; i < n; i++ {
		left[i] = right[i]
	}

	sort.Sort(EventsSortable(right))
	sort.Sort(EventsSortableEnd(left))

	ans := 0
	l := 0
	maxLeft := 0
	fmt.Println(right, "\n", left)
	for _, ev := range right {
		threshold := ev.startTime
		for l < n && left[l].endTime < threshold {
			if left[l].value > maxLeft {
				maxLeft = left[l].value
			}
		}

		if maxLeft+ev.value > ans {
			ans = maxLeft + ev.value
		}
	}
	return ans
}
