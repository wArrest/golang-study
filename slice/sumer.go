package slice

import "fmt"

func Sum(numbers []int) (res int) {
	for _, v := range numbers {
		res += v
	}
	return
}

func SumAll(numbersToSum ...[]int) (res []int) {
	for _, v := range numbersToSum {
		fmt.Println(len(v))
		if len(v) > 0 {
			res = append(res, Sum(v[0:]))
		}
	}
	return
}
