package main

import (
	"fmt"
	"strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := strings.ToUpper("Hello pinches prros")
	f("To upper %s\n", upper)
	f("To lower %s\n", strings.ToLower(upper))
	f("%s\n", strings.Title("tHiS a tExT TiTlE!!!!!"))
	// EqualFold returns true or false if s1 and s2 are unicode equals
	f("EqualFold: %v\n", strings.EqualFold("Orlando", "oRlaNdO")) // true
	f("EqualFold: %v\n", strings.EqualFold("Orlando", "ORLanDO")) // true
	// HasPrefix checks if the string starts with the second string argument
	f("Prefix: %v\n", strings.HasPrefix("Orlando", "or"))
	f("Prefix: %v\n", strings.HasPrefix("Orlando", "an"))
	// HasSuffix checks if the string ends with the second string argument
	f("Suffix: %v\n", strings.HasSuffix("Orlando", "do"))
	f("Suffix: %v\n", strings.HasSuffix("Orlando", "Do"))
	// Index returns the index where string containts the second argument, -1 otherwise
	f("Index: %v\n", strings.Index("Orlando", "an"))
	f("Index: %v\n", strings.Index("Orlando", "do"))
	// Count returns the number of occurrences given a substring
	f("Count: %v\n", strings.Count("Orlando", "o"))
	f("Count: %v\n", strings.Count("Orlando", "ñ"))
	// Repate returns a new string where will be concact the current string n times
	f("Repeat: %v\n", strings.Repeat("abcd ", 5))
	// TrimSpace returns the string after removing \t \n or whitespaces
	f("TrimSpace: %s\n", strings.TrimSpace("\t\t spacees prro \n"))
	// TrimLeft returns a new string after removing the \t \n or whitespaces from the left side
	// until the first character is encountered
	f("TrimLeft: %s\n", strings.TrimLeft("\t\t spacees prro \n", "\t\n"))
	//// TrimLeft returns a new string after removing the \t \n or whitespaces from the right side
	// until the last character is encountered
	f("TrimRight: %s\n", strings.TrimRight("\tspacees prro \n\n\n", "\t\n"))
	// Compares: returns 1, -1, 0 trying to compare the bytes of the s1, s2
	f("Compare: %v\n", strings.Compare("Orlando", "ORLANDO"))
	f("Compare: %v\n", strings.Compare("Orlando", "OrLaNdo"))
	f("Compare: %v\n", strings.Compare("Orlando", "orlanDo"))
	// Fields returns a string slice of the current string split by whitespace
	f("Fields: %v\n", strings.Fields("This an example string!"))
	f("Fields: %v\n", strings.Fields("Another text with\tspaces\n"))
	// Replacement returns a new string where every second argument is replace by the third argument.
	// The last argument represents how many replacement must occur, -1 means unlimited replacement
	f("%s\n", strings.Replace("abcd efg", "", "_", -1))
	f("%s\n", strings.Replace("abcd efg", "", "_", 4))
	f("%s\n", strings.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	// Joins returns an string given a string slice and the separator to join each string
	f("Join %s\n", strings.Join(lines, "+++"))
	// SplitAfter splits a string by whitespace when the second argument is encountered
	f("SplitAfter: %s\n", strings.SplitAfter("123++456++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	// TrimFunc trims a string base on the logic of the f function
	f("TrimFunc: %v\n", strings.TrimFunc("123 abc ABC \t", trimFunction))

}
