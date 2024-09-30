package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	ConvertIp()
}

/*
*
将ip转换为uint32
不能使用strconv函数
不能使用split函数
程序要足够健壮，能够处理输入的非法字符串
*/
var (
	zero     = '0'
	nine     = '9'
	sepBytes = '.'
)

func ConvertIp() (uint32, error) {
	ipStr := "999.0.0.1"
	fmt.Printf("input ip string : %s\n", ipStr)
	var tempSum uint32 = 0
	var result uint32 = 0
	sepCount := 0
	for _, ch := range ipStr { // ch is a rune
		fmt.Printf("input ip ch : %v\n", ch)
		if ch == sepBytes {
			fmt.Printf("input sep")
			if tempSum > 256 || tempSum < 0 {
				fmt.Printf("invalid input ip 256 ")
				return 0, errors.New("invalid input ipStr")
			}
			result = uint32(result*2 ^ 3 + tempSum)
			sepCount++
			continue
		}
		if ch < zero || ch > nine {
			fmt.Println("invalid input str ch")
			return 0, errors.New("invalid input ipStr")
		}
		tempSum = tempSum*10 + uint32(ch-zero)

	}
	if sepCount != 3 {
		fmt.Println("invalid input str")
		return 0, errors.New("invalid input ipStr")
	}
	fmt.Printf("result : %d", result)
	return result, nil

}
