package main

import (
	"fmt"
	"log"
	"strconv"
)

/**
题目描述:
存在一种虚拟IPV4地址，由4小节组成，每节的范围为0~255，以#号间隔，虚拟IPV4地址可以转换为一个32位的
整数，例如:128#0#255#255，转换为32位整数的结果为2147549183 (0x8000FFFF)1#0#0#0，
转换为32位整数的结果为16777216 (0x01000000)现以字符串形式给出一个虚拟IPv4地址，
限制第1小节的范围为1~128，即每一节范围分别为(1~128)#(0~255)#(0~255)#(0~255)，
要求每个IPv4地址只能对应到唯一的整数上。如果是非法IPv4，返回invalid IP
**/

func main() {
	fmt.Println(work([]int{100, 101, 1, 5}))
	fmt.Println(work([]int{1, 201, 5}))
}

// 将4个ip段转成32位的整型，如果ip段不符合规范则返回"invalid IP"
func work(segs []int) string {
	if len(segs) < 4 {
		return "invalid IP"
	}
	var res string
	if segs[0] < 1 || segs[0] > 255 {
		return "invalid IP"
	}
	res += fmt.Sprintf("%b", segs[0])
	for i := 1; i < 4; i++ {
		if segs[i] < 0 || segs[0] > 255 {
			return "invalid IP"
		}
		res += fmt.Sprintf("%08b", segs[i])
	}
	fmt.Println(res)
	num, err := strconv.ParseInt(res, 2, 64)
	if err != nil {
		log.Panicln(err)
	}
	return strconv.Itoa(int(num))
}
