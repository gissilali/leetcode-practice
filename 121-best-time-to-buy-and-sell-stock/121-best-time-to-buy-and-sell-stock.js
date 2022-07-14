/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  let maxProfit = 0;
  let buyDay = 0;
  let sellDay = 1;
  for (let i = 0; i < prices.length; i++) {
    while (sellDay < prices.length) {
      if (prices[buyDay] < prices[sellDay]) {
        let profit = prices[sellDay] - prices[buyDay];
        maxProfit = Math.max(profit, maxProfit);
      } else {
        buyDay = sellDay;
      }
      sellDay = sellDay + 1;
    }
  }
  return maxProfit;
};