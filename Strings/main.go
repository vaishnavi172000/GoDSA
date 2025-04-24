package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/isomorphic-strings/description/
// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters.
//
//	No two characters may map to the same character, but a character may map to itself.
//
// Example 1:
// Input: s = "egg", t = "add"
// Output: true
// Explanation:
// The strings s and t can be made identical by:
// Mapping 'e' to 'a'.
// Mapping 'g' to 'd'.
// Example 2:
// Input: s = "foo", t = "bar"
// Output: false
// Explanation:
// The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.
// Example 3:
// Input: s = "paper", t = "title"
// Output: true
func isIsomorphic(s string, t string) bool {
	// s:t
	m1 := make(map[rune]rune)
	// t:s
	m2 := make(map[rune]rune)
	l1 := len(s)
	l2 := len(t)
	if l1 != l2 {
		return false
	}
	for i := range l1 {
		s1 := rune(s[i])
		t1 := rune(t[i])
		if val, ok := m1[s1]; !ok {
			m1[s1] = t1
		} else {
			if val != t1 {
				return false
			}
		}

		if val, ok := m2[t1]; !ok {
			m2[t1] = s1
		} else {
			if val != s1 {
				return false
			}
		}

	}
	return true

}

// Time comeplexity NlogN
func SortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}

// https://leetcode.com/problems/group-anagrams/description/
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Explanation:
// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
// Example 2:
// Input: strs = [""]
// Output: [[""]]
// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	mp := make(map[string][]string)
	for _, str := range strs {
		key := SortString(str)
		if _, ok := mp[key]; !ok {
			mp[key] = []string{str}
		} else {
			mp[key] = append(mp[key], str)
		}
	}
	for _, v := range mp {
		result = append(result, v)
	}
	return result

}

// https://leetcode.com/problems/remove-outermost-parentheses/description/
// A valid parentheses string is either empty "", "(" + A + ")", or A + B,
// where A and B are valid parentheses strings, and + represents string concatenation.
// For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.
// A valid parentheses string s is primitive if it is nonempty, and there does not exist a way to split it into s = A + B,
// with A and B nonempty valid parentheses strings.
// Given a valid parentheses string s, consider its primitive decomposition: s = P1 + P2 + ... + Pk,
// where Pi are primitive valid parentheses strings.
// Return s after removing the outermost parentheses of every primitive string in the primitive decomposition of s.
// Example 1:

// Input: s = "(()())(())"
// Output: "()()()"
// Explanation:
// The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
// After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
// Example 2:

// Input: s = "(()())(())(()(()))"
// Output: "()()()()(())"
// Explanation:
// The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
// After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
// Example 3:

// Input: s = "()()"
// Output: ""
// Explanation:
// The input string is "()()", with primitive decomposition "()" + "()".
// After removing outer parentheses of each part, this is "" + "" = "".

func removeOuterParentheses(s string) string {
	depth := 0
	var chars []rune
	for _, v := range s {

		if string(v) == "(" {
			if depth != 0 {
				chars = append(chars, v)
			}
			depth++
		} else {
			if depth != 1 {
				chars = append(chars, v)
			}
			depth--
		}
	}
	return string(chars)

}

func main() {
	fmt.Println("Enter 2 strings")
	var s1, s2 string
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	fmt.Println("Enter your Choice \n1. Sort \n2. Isomorphic")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("S1 after sorting", SortString(s1))
	case 2:
		if isIsomorphic(s1, s2) {
			fmt.Println("Strings are isomorphic")
		} else {
			fmt.Println("Strings are not isomorphic")
		}
	default:
		fmt.Println("Invalid Choice")
	}

}
