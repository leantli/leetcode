package main

import "math"

// https://leetcode.cn/problems/poor-pigs/
// 458. 可怜的小猪

// 微软小鼠试毒题翻版
// 这里我们先确定能进行几轮实验，k = minutesToTest/minutesToDie
// 这里再让我们来考虑，一轮实验情况下，一只猪，能测几个桶？
// 显然，一只猪能测出两个桶是否有毒, 1 2 两个桶，猪吃桶1，死了则1有毒，没死则2有毒
// 那么一轮实验情况下，两只猪能测几个桶？其实最多能测4个桶
// 桶 1 2 3 4，其二进制分别为 01 10 11 100，这里不管100
// 猪a代表第一位进制为1，猪2代表第二位进制为1,则猪a吃桶1,3，猪b吃桶2,3
// 此时桶1有毒，则猪a死猪b活，桶2有毒，则猪ab都死，桶3有毒，则猪b死，猪ab都没死，则桶4有毒
// 以此类推，x只猪在一轮实验中能测出 2^x 个桶有毒
// 如果此题只限制一轮实验，则此时可以 2^x=buckets <--> x = log_2_(buckets)解决
// 但显然，这里还有可以多做几轮实验，那么这个几轮实验怎么体现？
// 我们知道，1只猪，一轮实验后有两种状态，要么一轮实验后死，要么一轮实验后活，因此这里是 2^x
// 而1只猪，两轮实验后有三种状态，要么一轮实验后死，要么两轮实验后死，要么两轮实验后活，以此类推猪的状态为 k+1, k 为实验次数
// 因此此题为 (k+1)^x = buckets <--> x = log_k+1_(buckets) <==> x = (lg(buckets)/lg(k+1))+1 向上取整
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	time := minutesToTest / minutesToDie
	base := time + 1
	// math.Ceil 对将结果向上取整
	return int(math.Ceil(math.Log(float64(buckets))/math.Log(float64(base)) - 1e-5))
}

// -1e-5是因为 125，1，4这个用例的时候由于数据精度问题，125求5的对数会算成3.000000000004，向上取整得到4，导致错误
// 这个还是有点离谱的
