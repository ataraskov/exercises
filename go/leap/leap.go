// leap is a package for working with leap years.
package leap

// IsLeapYear reports whether a given year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
