package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getMax(numbers ...int) int {
	res := 0
	for _, v := range numbers {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	var input string
	fmt.Print("Masukkan beberapa angka pisahkan antar angka dengan koma (,):")
	fmt.Scan(&input)
	tmp := strings.Split(input, ",")
	inputArray := make([]int, 0, len(tmp))
	for _, v := range tmp {
		v, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
			continue
		}
		inputArray = append(inputArray, v)
	}
	fmt.Println(
		getMax(inputArray...),
	)
}
