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

// * Solution -- Using Two Pointers -- Greedy -- Time O(n) - Space O(1)
func maxProfit(prices []int) int {
	l, r := 0, 1
	var profit int
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit = max(profit, prices[r]-prices[l])
		} else {
			l = r
		}
		r++
	}
	return profit
}

// * Solution -- Using Single Variable -- Greedy -- Time O(n) - Space O(1)
// func maxProfit(prices []int) int {
//     profit := 0
//     min := prices[0]
//     for _, price := range prices {
//         if min > price {
//             min = price
//         } else {
//             profit = max(profit, price - min)
//         }
//     }
//     return profit
// }


// * Solution -- Iterative - 2 Loops -- Brute Force -- Time O(n^2) - Space O(1)
// func maxProfit(prices []int) int {
//     var profit int
//     for i := 0; i < len(prices)-1; i++{
//         for j := i+1; j < len(prices); j++ {
//             profit = max(profit, prices[j] - prices[i])
//         }
//     }
//     return profit
// }
