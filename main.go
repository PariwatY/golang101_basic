package main

import (
	"gobasic/customer"
)

func main() {

	x := customer.Person{}
	x.SetName("Pek")
	println(x.GetName())

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run(":8081")

}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func cal(f func(int, int) int) {
	sum := f(50, 30)
	println(sum)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func hello(name string) string {
	return "hello" + name
}
