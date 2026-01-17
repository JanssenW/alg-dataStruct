package main

import "fmt"

func main() {

	//s := "lleetcode"
	//idx := firstUniqCharSimple(s)
	//fmt.Println("First unique char index is: ", idx)
	//
	//nums := []int{2, 7, 11, 15}
	//target := 9
	//pair := twoSum(nums, target)
	//fmt.Println("Two sum indices are: ", pair[0], " , ", pair[1])
	roman := "XXIV"
	integer := romanToInt(roman)
	fmt.Println("Roman numeral ", roman, " is integer ", integer)
}

func romanToInt(s string) int {

	sRune := []rune(s)
	integer, i := 0, 0

	for i < len(sRune) {
		roman := string(sRune[i])
		romanNext := ""
		if i+1 < len(sRune) {
			romanNext = string(sRune[i+1])
		}

		// Populate the map with roman numerals and their integer values
		romanMap := map[byte]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}

		// Check if next roman numeral is greater (subtractive case)
		if romanNext != "" && romanMap[romanNext[0]] > romanMap[roman[0]] {
			integer = romanMap[romanNext[0]] - romanMap[roman[0]] + integer
			i += 2
		} else {
			integer = romanMap[roman[0]] + integer
			i += 1
		}
	}

	return integer
}

func twoSum(nums []int, target int) []int {

	// The first approach were brute force, with two nested loops, O(n^2) time complexity.
	// A better approach is to use a map to store the numbers and their indices while iterating the array only once, O(n) time complexity.

	//Create map to store number and its index
	m := make(map[int]int)
	diff := 0

	//Iterate over the numbers
	for i, num := range nums {
		//Calculate the difference needed to reach the target
		diff = target - num

		//Check if the difference is already in the map - we are using the value as key
		// val_idx gives the value stored in the map for key diff, ok is true if key exists
		if val_idx, ok := m[diff]; ok {
			// which means that if we found it, return the indices (map value)
			return []int{val_idx, i}
		} else {
			// If not found, store the number (key) and its index (value) in the map
			m[num] = i
		}
	}

	return []int{}
}

func firstUniqCharSimple(s string) int {

	freq := make(map[rune]int)

	sRune := []rune(s)

	for _, ch := range sRune {
		freq[ch]++
	}

	for i, ch := range sRune {
		if freq[ch] == 1 {
			return i
		}
	}

	return -1

}

func firstUniqChar(s string) int {

	//Create structure to hold count and index
	type info struct {
		count int
		idx   int
	}

	//Transform string in a rune slice to handle unicode characters
	sRune := []rune(s)

	//Create map [key: rune] value: info - empty
	m := make(map[rune]info)

	//iterate over rune slice. This for gives me indices and characters
	for i, ch := range sRune {

		//Entry is a copy of the value stored in map for key ch. OK is true if key exists.
		if entry, ok := m[ch]; ok {
			//Count++ on the copied structure
			entry.count++
			//Store back the updated structure in the map
			m[ch] = entry

		} else {
			//If key does not exist directly create a new structure and store it in the map
			m[ch] = info{count: 1, idx: i}

		}

	}

	//Iterate again over the rune slice to return the index of the first unique character
	for j, ch := range sRune {
		if m[ch].count == 1 {
			return j
		}
	}

	return -1

}
