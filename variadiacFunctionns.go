package main

import "fmt"

// alloww you to create functions that accepts a variablww number of arguments
//this FUNCTIONS ARE defined by prefixing the type of the last parameter with an ellipsis, this is how an ellipses looks like like spread operator ...
func main() { 
// func functionName(param1 type1, param2 type2, param3 ...type3) rturnType{
// 	function body
// }
// ...type3 mans no limit for the number  of parameter for type3
fmt.Println("Sum 1, 2, 3:", sum(1, 2, 3))
}

//variafsic function to calculate the sume of integers
func sum(nums ...int) int {
total := 0
for _, v := range nums {
	total += v
}
return total
}