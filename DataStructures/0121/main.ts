function maxProfit(prices: number[]): number {
    let minPrice = 1e9
    let price = 1e9
    let profit = 0
    let tempProfit = 0
    for (price of prices){
        tempProfit = price - minPrice
        if (tempProfit < 0){
            minPrice = price
        } else if (tempProfit > 0 && tempProfit > profit){
            profit = tempProfit
        }
    }
    return profit
};