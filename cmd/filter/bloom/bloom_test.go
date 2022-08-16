package bloom

import (
	"fmt"
	"testing"
)

var words = []string{
	"abound", "abounds", "abundance",
	"abundant", "accessible", "bloom",
	"blossom", "bolster", "bonny",
	"bonus", "bonuses", "coherent",
	"cohesive", "colorful", "comely",
	"comfort", "gems", "generosity",
	"generous", "generously", "genial",
	"bluff", "cheater", "hate",
	"war", "humanity", "racism",
	"hurt", "nuke", "gloomy",
	"facebook", "geeksforgeeks", "twitter",
}

func TestName(t *testing.T) {
	bloomFilter := NewFilter(100)

	for _, word := range words {
		if ok := bloomFilter.Insert(word); !ok {
			fmt.Println(word, "is Probably already present")
		} else {
			fmt.Println(word, "inserted")
		}
	}

	for _, word := range words {
		fmt.Println(word, bloomFilter.Lookup(word))
	}
	word := "facebook"
	fmt.Println(word, bloomFilter.Lookup(word))
}
