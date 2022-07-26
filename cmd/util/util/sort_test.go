package util

import (
	"fmt"
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println("\n##### SearchINts not works in descending order #####")
	intSlice := []int{55, 54, 53, 52, 51, 50, 48, 36, 15, 5}
	x := 36
	pos := sort.SearchInts(intSlice, x)
	fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)

	fmt.Println("\n###### Search work in descending order $$$$$")
	i := sort.Search(len(intSlice), func(i int) bool {
		return intSlice[i] <= x
	})
	fmt.Printf("Found %d at index %d in %v\n", x, i, intSlice)

	strSlice := sort.StringSlice{"Washington", "Texas", "Ohio", "Nevada", "Montana", "Indiana", "Alaska"}
	strSlice2 := []string{"Washington", "Texas", "Ohio", "Nevada", "Montana", "Indiana", "Alaska"}
	y := "Montana"
	posStr := sort.SearchStrings(strSlice, y)
	fmt.Println(posStr, strSlice.Search(y))

	posStrWorks := sort.Search(len(strSlice2), func(i int) bool {
		return strSlice[i] == y
	})
	fmt.Println(posStrWorks)

	fltSlice := []float64{10.10, 20.10, 30.15, 40.15, 58.95}
	z := 40.15
	k := sort.Search(len(fltSlice), func(i int) bool {
		return fltSlice[i] >= z
	})
	fmt.Printf("Found %f at index %d in %v\n", z, k, fltSlice)
}

type Mobile struct {
	Brand string
	Price int
}

type ByPrice []Mobile

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }

type ByBrandDesc []Mobile

func (a ByBrandDesc) Len() int           { return len(a) }
func (a ByBrandDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBrandDesc) Less(i, j int) bool { return a[i].Brand > a[j].Brand }

func TestFields(t *testing.T) {
	mobile := []Mobile{
		{"Sony", 952},
		{"Nokia", 468},
		{"Apple", 1219},
		{"Samsung", 1045},
	}

	fmt.Println("\n##### Before Sort #######")
	PrintList(mobile)

	fmt.Println("#\n\n###### Sort By Price [ascending] #######")
	fmt.Println("\nFound mobile1 price is sorted :", sort.IsSorted(ByPrice(mobile))) // false
	sort.Sort(ByPrice(mobile))
	PrintList(mobile)
	fmt.Println("\nFound mobile1 price is sorted :", sort.IsSorted(ByPrice(mobile))) // false
	sort.Sort(sort.Reverse(ByPrice(mobile)))
	PrintList(mobile)
	fmt.Println("#\n\n###### Sort By Brand ")
	fmt.Println("\nFound mobile1 price is sorted :", sort.IsSorted(ByBrandDesc(mobile))) // false
	sort.Sort(ByBrandDesc(mobile))
	fmt.Println("\nFound mobile1 price is sorted :", sort.IsSorted(ByBrandDesc(mobile))) // false
	PrintList(mobile)

	sort.Slice(mobile, func(i, j int) bool {
		return mobile[i].Brand < mobile[j].Brand
	})
	PrintList(mobile)
	sort.Slice(mobile, func(i, j int) bool {
		return mobile[i].Brand > mobile[j].Brand
	})
	PrintList(mobile)
	sort.Slice(mobile, func(i, j int) bool {
		return mobile[i].Brand > mobile[j].Brand
	})
	PrintList(mobile)
}

func PrintList(mobile []Mobile) {
	for _, mob := range mobile {
		fmt.Println(mob.Brand, mob.Price)
	}
	println()
}
