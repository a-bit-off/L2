package main

import (
<<<<<<< HEAD
	"fmt"
	"pattern"
=======
	"pattern"
	"pattern/02_builder"
>>>>>>> 38619e63b0cc8c2b3bb7e4a980b341bbbdcc409c
)

func main() {
<<<<<<< HEAD
	fmt.Println("Facade:")
	pattern.RunFacade()
	fmt.Println()

	fmt.Println("Builder:")
	pattern.RunBuilder()
	fmt.Println()

	fmt.Println("Visitor:")
	pattern.RunVisitor()
	fmt.Println()

=======
	//pattern.RunFacade()
	//pattern.RunBuilder()
<<<<<<< HEAD
	pattern.RunBuilderV2()
>>>>>>> 0eedc6e1923ecfaf078a3cc47485a376dc32e59d
=======
	builder.RunBuilderV2()
	pattern.RunStrategy()
>>>>>>> 38619e63b0cc8c2b3bb7e4a980b341bbbdcc409c
}
