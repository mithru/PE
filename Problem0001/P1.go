package main

import "fmt"

var num = 1000;

func main() {
	var sum = 0;
	for i := 0; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum);
}