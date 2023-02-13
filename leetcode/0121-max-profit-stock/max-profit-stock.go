package leetcode

import "math"

func maxProfit(prices []int) int {
	min := math.MaxUint32
	maxProfit := 0

	for _, price := range prices {
		if price > min {
			if price-min > maxProfit {
				maxProfit = price - min
			}
		} else {
			min = price
		}
	}

	return maxProfit
}
