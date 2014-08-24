package main

import (
"math"
"fmt"
)

func main() {
	fmt.Println(largestPrimeFactor(600851475143));
}

func largestPrimeFactor(num int64) int64{
	n := float64 (num)
	for i := math.Floor(math.Sqrt(n)); i >= 2 ; i -= 1.0 {
		//fmt.Println(i)
		if math.Mod(n, i) == 0 {
			if isPrime(int(i)) {
				return int64(i)
			}
		}
	}	
	return num;
}

func isPrime(num int) bool {
	// 1 and and anything lower aren't primes
	if num <= 1 {
		return false
	}
	
	// math.Mod requires a float
	n := float64(num)
	
	// get rid of all even numbers except 2
	if math.Mod(n, 2) == 0 {
		return num == 2
	}
	
	// solving only until the rounded sqrt of n
	for i := 3.0; i <= math.Floor(math.Sqrt(n)); i += 2.0 {
	
		if math.Mod(n, i) == 0 {
			return false
		}
	}
	return true
}