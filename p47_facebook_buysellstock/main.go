// Given a array of numbers representing the stock prices of a company in chronological order, write a function that calculates the maximum profit you could have made from buying and selling that stock once. You must buy before you can sell it.
// For example, given [9, 11, 8, 5, 7, 10], you should return 5, since you could buy the stock at 5 dollars and sell it at 10 dollars.

package main

import "fmt"

func main() {
	fmt.Println(maxProfitBuySellOnce([]int{5, 3, 1}))
	fmt.Println(maxProfitBuySellOnce([]int{5}))
	fmt.Println(maxProfitBuySellOnce([]int{5, 10}))
	fmt.Println(maxProfitBuySellOnce([]int{9, 11, 8, 5, 7, 10}))
}

func maxProfitBuySellOnce(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	tempMin := prices[0]
	tempMaxProfit := 0
	for i := 1; i < len(prices); i++ {
		profit_if_sell_now := prices[i] - tempMin
		if profit_if_sell_now > tempMaxProfit {
			tempMaxProfit = profit_if_sell_now
		}
		if tempMin > prices[i] {
			tempMin = prices[i]
		}
	}

	return tempMaxProfit
}
