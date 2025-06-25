package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func CombinedNumber(nums []int) string {
	fmt.Printf("origin_list: %v\n", nums)

	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})
	res := strings.Join(strNums,"")
	fmt.Printf("result_list:%v", res)
	fmt.Printf("\n")
	if len(strNums) > 0 && strNums[0] == "0" {
		return "0"
	}

	return joinStrings(strNums)
}

func joinStrings(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}
