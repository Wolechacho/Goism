// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear calculates if the year given is a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
