package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func squareRoot(n int) int {
	if n == 0 {
		return 0
	}
	target := n
	res := 1
	for 1 <= n {
		sq := n / 2
		//consider the ceil value if mod not 0.
		if n%2 != 0 {
			sq += 1
		}
		if sq*sq <= target {
			res = sq
			return res
		} else {
			n = n / 2
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	res := squareRoot(n)
	fmt.Println(res)
}
