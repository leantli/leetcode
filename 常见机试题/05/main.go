package main

import (
	"log"
	"sort"
)

/**
题目描述:
双十一众多商品进行打折销售，小明想购买自己心仪的一些物品，但由于受购买资金限制，所以他决定从众多心仪商品中购买三件，而且想尽可能的花完资金，现在请你设计一个程序帮助小明计算尽可能花费的最大资金数额。
输入描述:
输入第一行为一维整型数组M，数组长度小于100，数组元素记录单个商品的价格，单个商品价格小于1000.
输入第二行为购买资金的额度R，R小于100000。
输出描述:
输出为满足上述条件的最大花费额度
注意:如果不存在满足上述条件的商品，请返回-1.
补充说明:
输入格式是正确的，无需考虑格式错误的情况
示例1
输入:
23,26,36,27
78
输出:
76
说明:
金额23、26和27相加得到76，而且最接近且小于输入金额78

示例2
输入:
23,30,40
26
输出:
-1
说明:
因为输入的商品，无法组合出来满足三件之和小于26.故返-1
**/

func main() {
	log.Println(calculate([]int{23, 26, 36, 27}, 78))
	log.Println(calculate([]int{23, 30, 40}, 26))
	log.Println(calculate([]int{100, 200, 300, 400, 500}, 800))
}

// 根据商品价值和预算，输出购买三件物品能花的最大钱
func calculate(prices []int, budget int) int {
	sort.Ints(prices)
	res := -1
	for i := len(prices) - 1; i >= 2; i-- {
		if prices[i] > budget {
			continue
		}
		for j := i - 1; j >= 1; j-- {
			if prices[j]+prices[i] > budget {
				continue
			}
			for k := j - 1; k >= 0; k-- {
				sum := prices[j] + prices[i] + prices[k]
				if sum > budget {
					continue
				}
				if sum > res {
					res = sum
				}
			}
		}
	}
	return res
}
