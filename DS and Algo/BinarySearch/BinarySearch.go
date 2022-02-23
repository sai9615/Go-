package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
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
	res := BinarySearch(arr, target)
	fmt.Println(res)
}
