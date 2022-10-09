package main

import (
	"fmt"
	"os"
	"strconv"
)

func kaprekar(s string) string {
	sdown := s
	sup := s
	var res int
	if s == "6174" {
		return "done"
	}

	// sort s into descending order
	for i := 0; i < len(sdown); i++ {
		for j := i + 1; j < len(sdown); j++ {
			if sdown[i] < sdown[j] {
				sdown = sdown[:i] + string(sdown[j]) + sdown[i+1:j] + string(sdown[i]) + sdown[j+1:]
			}
		}
	}
	// sort s into ascending order
	for i := 0; i < len(sup); i++ {
		for j := i + 1; j < len(sup); j++ {
			if sup[i] > sup[j] {
				sup = sup[:i] + string(sup[j]) + sup[i+1:j] + string(sup[i]) + sup[j+1:]
			}
		}
	}
	res = Atoi(sdown) - Atoi(sup)
	fmt.Println(res)
	return kaprekar(strconv.Itoa(res))
}

func Atoi(s string) int {
	sign := 1
	res := 0
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, e := range s {
		if e >= '0' && e <= '9' {
			res = res*10 + int(e-48)
		}
	}
	return res * sign
}

func main() {
	if len(os.Args) == 2 && len(os.Args[1]) == 4 {
		// check if theres atleast two different digits
		for i := 0; i < len(os.Args[1]); i++ {
			for j := i + 1; j < len(os.Args[1]); j++ {
				if os.Args[1][i] == os.Args[1][j] {
					fmt.Println("Input must have at least two different digits")
					return
				}
			}
		}
		fmt.Println(os.Args[1])
		fmt.Println(kaprekar(os.Args[1]))
	} else {
		fmt.Println("Please enter a four digit number")
	}
}
