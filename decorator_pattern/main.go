package main

import (
	decoratorp "decorator_pattern/decorator_p"
	"fmt"
)

/*

Problems faced during the implementation:

> Ran into stack overflow error (usually happens when there endless recursion)
> Here, when we call getIngredients() with Milk{} struct. instead milk.decorator.getIngredients().. i called milk.getIngredients()
> Should be aware of that.
-----------------------------------------------------------------------------------------------------------------

What is the Decorator Pattern?
The Decorator Pattern allows you to dynamically extend objects' behavior by "wrapping" them with additional functionality.
It relies on composition rather than inheritance, which aligns perfectly with Go\'s design philosophy.
Using the Decorator Pattern, you can add responsibilities to objects in a non-intrusive and scalable way.

-----------------------------------------------------------------------------------------------------------------
Key Features of the Decorator Pattern
Dynamic Behavior: Extends functionality at runtime without altering the original object.
Composition over Inheritance: Leverages Go's ability to compose behaviors via interfaces and embedded structs.
Modular Design: Enhances maintainability and reusability by avoiding tightly coupled code.

*/

func main() {
	myCoffee := decoratorp.NewCoffee(5, "coffee")
	fmt.Printf("Price: %d | Ingredients: [%v] \n", myCoffee.GetPrice(), myCoffee.GetIngredients())
	myCoffee = decoratorp.NewMilk(myCoffee)
	fmt.Printf("Price: %d | Ingredients: [%v] \n", myCoffee.GetPrice(), myCoffee.GetIngredients())
	myCoffee = decoratorp.NewSugar(myCoffee)
	fmt.Printf("Price: %d | Ingredients: [%v]", myCoffee.GetPrice(), myCoffee.GetIngredients())

}
