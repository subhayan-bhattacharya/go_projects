package main

import "fmt"

type Item struct {
    ID       string
    Name     string
    Price    float64
    Quantity int
}

type ShoppingCart struct {
    Items []*Item  // Notice: slice of pointers!
    Owner string
}

func (s *ShoppingCart) AddItem(item *Item) {
	s.Items = append(s.Items, item)
}

func (s *ShoppingCart) UpdateQuantity(itemId string, newQuantity int) error {
	for _, item := range s.Items {
		if item != nil {
			if item.ID == itemId {
				item.Quantity = newQuantity
				return nil
			}
		}
	}
	return fmt.Errorf("The item with item id %s does not exist", itemId)
}

func (s *ShoppingCart) GetTotal() float64 {
	var total float64
	for _, item := range s.Items {
		if item != nil {
			total += item.Price * float64(item.Quantity)
		}
	}
	return total
}

func (s *ShoppingCart) RemoveItem(itemID string) error {
	for index, item := range s.Items {
		if item != nil {
			if item.ID == itemID {
				s.Items = append(s.Items[:index], s.Items[index + 1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("The item id %s is not found", itemID)
}

func main() {
	item1 := Item{
		ID: "123",
		Name: "Pen",
		Price: 12,
		Quantity: 2,
	}
	item2 := Item{
		ID: "456",
		Name: "keyboard",
		Price: 120,
		Quantity: 2,
	}
	item3 := Item{
		ID: "789",
		Name: "Mouse",
		Price: 40,
		Quantity: 3,
	}

	shoppingCart := ShoppingCart{
		Owner: "Subhayan",
	}

	shoppingCart.AddItem(&item1)
	shoppingCart.AddItem(&item2)
	shoppingCart.AddItem(&item3)
	fmt.Printf("The total of the shopping cart is %f\n", shoppingCart.GetTotal())
	shoppingCart.RemoveItem("123")
	fmt.Printf("The total of the shopping cart after removing is %f\n", shoppingCart.GetTotal())
}