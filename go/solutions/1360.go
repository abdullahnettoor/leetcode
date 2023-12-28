package solutions

import "time"

/* Question **********************************
? 1360. Number of Days Between Two Dates
Write a program to count the number of days between two dates.

The two dates are given as strings, their format is YYYY-MM-DD as shown in the examples.

Example 1:
Input: date1 = "2019-06-29", date2 = "2019-06-30"
Output: 1

Example 2:
Input: date1 = "2020-01-15", date2 = "2019-12-31"
Output: 15

Constraints:
The given dates are valid dates between the years 1971 and 2100.
*********************************************/

func Output1360() any {
	return daysBetweenDates("2019-06-29", "2019-06-30")
}

// * Solution 1
func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse("2006-01-02", date1)
	d2, _ := time.Parse("2006-01-02", date2)
	return int(d2.Sub(d1).Abs().Hours() / 24)
}
