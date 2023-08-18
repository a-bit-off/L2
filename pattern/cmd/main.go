package main

import (
	"fmt"
	"pattern"
)

func main() {
	fmt.Println("Facade:")
	pattern.RunFacade()
	fmt.Println()

	fmt.Println("Builder:")
	pattern.RunBuilder()
	fmt.Println()

	fmt.Println("Visitor:")
	pattern.RunVisitor()
	fmt.Println()

}
