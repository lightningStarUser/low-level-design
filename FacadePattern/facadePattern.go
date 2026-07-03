package main

import (
	"fmt"
)

type Payment struct{}

func (p *Payment) Charge(accountID string, amount float64) bool{
	fmt.Printf("Charging amount %f for the account %s\n", amount, accountID)
	return true
}

type Warehouse struct{}

func (w *Warehouse) CheckStock(itemID string) bool{
	fmt.Printf("Checking stock for item %s\n", itemID)
	return true
}

func (w *Warehouse) ReserveItem(itemID string){
	fmt.Printf("Reserving item: %s\n", itemID)
}

type Shipping struct{}

func (s *Shipping) ShipItem(itemID string, address string) string{
	fmt.Printf("Shipping item %s to address %s\n", itemID, address)
	return "TRACKING_ID"
}

type OrderFacade struct{
	payment *Payment
	warehouse *Warehouse
	shipping *Shipping
}

func (o *OrderFacade) PlaceOrder(accountID string, itemID string, amount float64, address string) (string, error){
	if !o.warehouse.CheckStock(itemID){
		return "", fmt.Errorf("item is out of stock")
	}
	o.warehouse.ReserveItem(itemID)
	if !o.payment.Charge(accountID, amount) {
		return "", fmt.Errorf("payment failed")
	}
	trackID := o.shipping.ShipItem(itemID, address)
	return trackID, nil
}

func main() {
	order := &OrderFacade{
		payment: &Payment{},
		warehouse: &Warehouse{},
		shipping: &Shipping{},
	}
	trackID, err := order.PlaceOrder("lightningUser", "AB_123", 678.34, "Bengaluru, Karnataka, India")
	if err != nil {
		fmt.Println("Error placing order:", err)
	} else {
		fmt.Println("Order placed successfully. Tracking ID:", trackID)
	}
}