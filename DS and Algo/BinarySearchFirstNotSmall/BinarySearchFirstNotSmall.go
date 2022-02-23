package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func firstNotSmaller(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	index := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] >= target {
			index = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return index
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	res := firstNotSmaller(arr, target)
	fmt.Println(res)
}
