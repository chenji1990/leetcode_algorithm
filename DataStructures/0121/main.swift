class Solution {
    func maxProfit(_ prices: [Int]) -> Int {
        var minPrice = Int.max
        var profit = 0
        var tempProfit = 0

        for price in prices{
            tempProfit = price - minPrice
            if tempProfit < 0{
                minPrice = price
            } else if tempProfit > 0 && tempProfit > profit{
                profit = tempProfit
            }
        }

        return profit
    }
}