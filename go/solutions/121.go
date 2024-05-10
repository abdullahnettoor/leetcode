package solutions

/* Question *******************************
? 121. Best Time to Buy and Sell Stock
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

Constraints:
1 <= prices.length <= 10^5
0 <= prices[i] <= 10^4
******************************************/

func Output121() any {
	return maxProfit([]int{7, 1, 5, 3, 6, 4})
}

// * Solution 1 -- Using Two Pointers -- Sliding Window -- O(n)
func maxProfit(prices []int) int {
	var lowest, highest = 10000, 0
	days := len(prices)
	var profit, k = 0, 0
	for i := range prices {
		if lowest > prices[i] {
			lowest = prices[i]
		}
		if highest < prices[i] {
			highest = prices[i]
		}
		if k == days-1 {
			k = i + 1
		}
	}

	return profit
}

// * Solution 2 -- Using Nested Loop -- Brute Force -- O(n^2)
// func maxProfit(prices []int) int {
//     var profit int
//     for i := 0; i < len(prices)-1; i++{
//         for j := i+1; j < len(prices); j++ {
//             profit = max(profit, prices[j] - prices[i])
//         }
//     }

//     return profit
// }