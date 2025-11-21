package main

import (
	"fmt"
	"math"
)

func FooBar()  {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}
	for i := len(nums) - 1; i >= 0; i-- {
		isNumberPrime := isPrime(nums[i])
		if isNumberPrime {
			continue
		}

		if nums[i] % 3 == 0 && nums[i] % 5 == 0 {
			fmt.Print("FooBar, ")
		} else if nums[i] % 3 == 0 {
			fmt.Print("Foo, ")
		} else if nums[i] % 5 == 0 {
			fmt.Print("Bar, ")
		} else {
			fmt.Printf("%d, ", nums[i])
		}
	}
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	if num == 2 || num == 3 {
		return true
	}

	if num % 2 == 0 || num % 3 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(num)))

	for i := 5; i < limit; i += 6 {
		if num % i == 0 || num % (i + 2) == 0 {
			return false
		}
	}

	return true
}