package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Enter the numbers a, b and c respectively :")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Println(a, b, c)

	sum := add(a, b, c)
	fmt.Println("Sum of a, b and c = : ", sum)

	average(a, b, c)

	min_value, max_value := min_max(a, b, c)
	fmt.Println("Minimum of the three is : ", min_value)
	fmt.Println("Maximum of the three is : ", max_value)

	radius := b
	area_of_circle := calculate_circle_area(float64(radius))
	fmt.Println("Area of Circle = ", area_of_circle)
}
