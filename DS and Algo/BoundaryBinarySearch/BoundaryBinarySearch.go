package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BoundaryBinarySearch(arr []bool) int {
	if len(arr) == 0 {
		return -1
	}
	start := 0
	end := len(arr) - 1
	index := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == true {
			index = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return index
}

func arrayAtob(arr []string) []bool {
	res := []bool{}
	for _, x := range arr {
		v, _ := strconv.ParseBool(x)
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
	arr := arrayAtob(splitWords(scanner.Text()))
	res := BoundaryBinarySearch(arr)
	fmt.Println(res)
}
