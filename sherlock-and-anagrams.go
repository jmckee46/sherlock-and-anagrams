package main

import (
	"sort"
)

func sherlockAndAnagrams(s string) int32 {
	sLength := int32(len(s))
	var count int32
	var sortedBaseString string
	var sortedComparedToString string

	for w := int32(0); w < sLength-1; w++ {
		for i := int32(0); i < sLength-w-1; i++ {
			sortedBaseString = SortString(s[i : i+w+1])
			for j := i + 1; j < sLength-w; j++ {
				sortedComparedToString = SortString(s[j : j+w+1])
				if sortedBaseString == sortedComparedToString {
					count++
				}
			}
		}
	}

	return count
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortString(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {}
