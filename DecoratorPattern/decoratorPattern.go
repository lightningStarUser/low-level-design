package main

import (
	"fmt"
)

type Pizza interface{
	getPrice() float64
}

type Margherita struct{}

func (m *Margherita) getPrice() float64{
	fmt.Println("Margherita pizza is ready")
	return 200
}

type CheeseTopping struct{
	pizza Pizza
}

func (c *CheeseTopping) getPrice() float64{
	fmt.Println("Added cheese topping")
	return c.pizza.getPrice() + 40
}

type JalapenoTopping struct{
	pizza Pizza
}

func (j *JalapenoTopping) getPrice() float64{
	fmt.Println("Added jalapeno topping")
	return j.pizza.getPrice() + 60
}

func main(){
	margherita := &Margherita{}
	cheese := &CheeseTopping{pizza: margherita}
	jalapeno := &JalapenoTopping{pizza: cheese}
	fmt.Println("Total price of pizza is", jalapeno.getPrice())
}