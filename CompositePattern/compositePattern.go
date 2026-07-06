package main

import "fmt"

type CartItem interface {
	getPrice() float64
	display()
}

type Product struct {
	name  string
	price float64
}

func (p *Product) getPrice() float64 {
	return p.price
}

func (p *Product) display() {
	fmt.Printf("Product Name: %s \n", p.name)
}

type ProductBundle struct {
	name     string
	products []CartItem
}

func (pb *ProductBundle) addItem(cartItem CartItem) {
	pb.products = append(pb.products, cartItem)
}

func (pb *ProductBundle) getPrice() float64 {
	totalPrice := 0.0
	for _, val := range pb.products {
		totalPrice += val.getPrice()
	}
	return totalPrice
}

func (pb *ProductBundle) display() {
	fmt.Printf("Product Bundle Name: %s \n", pb.name)
	for _, val := range pb.products {
		val.display()
	}
}

func main() {
	// defining some products
	mouse := &Product{
		name:  "Mouse",
		price: 899,
	}
	keyboard := &Product{
		name:  "Keyboard",
		price: 2399,
	}
	monitor := &Product{
		name:  "Monitor",
		price: 3799,
	}
	cable := &Product{
		name:  "Cable",
		price: 299,
	}
	cable.display()
	fmt.Printf("%f\n", cable.getPrice())
	// create a bundle

	computerBundle := &ProductBundle{
		name:     "Computer Bundle",
		products: []CartItem{},
	}
	computerBundle.addItem(mouse)
	computerBundle.addItem(keyboard)
	computerBundle.addItem(monitor)
	totalPrice := computerBundle.getPrice()
	computerBundle.display()
	fmt.Printf("Total price: %f\n", totalPrice)
}