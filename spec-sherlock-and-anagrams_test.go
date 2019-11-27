package main

import "testing"

func TestSherlockAndAnagrams1(t *testing.T) {
	s := "abba"
	// s := "abcdefgh"
	anagrams := sherlockAndAnagrams(s)

	if anagrams != 4 {
		t.Errorf("got %d instead of 4", anagrams)
	}
}

func TestSherlockAndAnagrams2(t *testing.T) {
	s := "abcd"
	anagrams := sherlockAndAnagrams(s)

	if anagrams != 0 {
		t.Errorf("got %d instead of 0", anagrams)
	}
}

// func TestSherlockAndAnagrams3(t *testing.T) {
// 	s := "ifailuhkqq"
// 	anagrams := sherlockAndAnagrams(s)

// 	if anagrams != 3 {
// 		t.Errorf("got %d instead of 3", anagrams)
// 	}
// }

// func TestSherlockAndAnagrams4(t *testing.T) {
// 	s := "kkkk"
// 	anagrams := sherlockAndAnagrams(s)

// 	if anagrams != 10 {
// 		t.Errorf("got %d instead of 10", anagrams)
// 	}
// }
