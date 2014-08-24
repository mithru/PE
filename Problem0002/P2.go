package main

import "fmt"

var num1 = 1;
var num2 = 1;
var sum = 0;
var max = 4000000;

var sum_of_even_values = 0;

func main() {
	for num1 < max { 
	if num1 % 2 == 0 {
		sum_of_even_values += num1;
	}
	sum = num1 + num2;
	num2 = num1;
	num1 = sum;
	fmt.Println(sum);	
	}
	fmt.Println("Answer :", sum_of_even_values);
}
