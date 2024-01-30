package solutions

/* Question **********************************
? 2839. Check if Strings Can be Made Equal With Operations I
You are given two strings s1 and s2, both of length 4, consisting of lowercase English letters.
You can apply the following operation on any of the two strings any number of times:
Choose any two indices i and j such that j - i = 2, then swap the two characters at those indices in the string.
Return true if you can make the strings s1 and s2 equal, and false otherwise.

Example 1:
Input: s1 = "abcd", s2 = "cdab"
Output: true
Explanation: We can do the following operations on s1:
- Choose the indices i = 0, j = 2. The resulting string is s1 = "cbad".
- Choose the indices i = 1, j = 3. The resulting string is s1 = "cdab" = s2.

Example 2:
Input: s1 = "abcd", s2 = "dacb"
Output: false
Explanation: It is not possible to make the two strings equal.

Constraints:
s1.length == s2.length == 4
s1 and s2 consist only of lowercase English letters.
*********************************************/

func Output2839() any {
	return canBeEqual("abcd", "dacb")
}

// * Solution 1
func canBeEqual(s1 string, s2 string) bool {
	// Check if the strings are already equal
	if s1 == s2 {
		return true
	}

	// Check if swapping indices 0 and 2 in s1 makes it equal to s2
	if s1[0] == s2[2] && s1[2] == s2[0] && s1[1] == s2[1] && s1[3] == s2[3] {
		return true
	}

	// Check if swapping indices 1 and 3 in s1 makes it equal to s2
	if s1[1] == s2[3] && s1[3] == s2[1] && s1[0] == s2[0] && s1[2] == s2[2] {
		return true
	}

	// Check if performing both swaps makes it equal
	return s1[0] == s2[2] && s1[2] == s2[0] && s1[1] == s2[3] && s1[3] == s2[1]
}

// * Solution 2
// func swap(s string, i, j int) string {
// 	b := []byte(s)
// 	b[i], b[j] = b[j], b[i]

// 	return string(b)
// }

// func canBeEqual(s1 string, s2 string) bool {
// 	return s1 == s2 ||
// 		swap(s1, 0, 2) == s2 ||
// 		swap(s1, 1, 3) == s2 ||
// 		swap(swap(s1, 0, 2), 1, 3) == s2
// }
