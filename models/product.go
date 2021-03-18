package models

type Product struct {
	Id 			uint
	Title 		string `gorm:"unique"`
	Type 		string
	Description string
	Filename    string
	Height 		float64
	Width 		float64
	Price 		float64
	Rating 		float64
}

type CartItem struct {
	Product Product
	Qty 	int
}

type Cart struct {
	CartItems 	[]CartItem
	TotalPrice 	float64
	NumItems 	int
}

func (cart *Cart) AddItem(item CartItem){

	itemFound := CartItem{}
	for _, cartItem1 := range cart.CartItems {
		if cartItem1.Product.Id == item.Product.Id {
			itemFound = cartItem1
			break
		}
	}

	if itemFound.Product.Id == 0 {

		cart.CartItems = append(cart.CartItems, item)
		cart.TotalPrice = cart.TotalPrice + (item.Product.Price * float64(item.Qty))
		cart.NumItems += 1
	} else {

		//I'll ignore this item if it already exist into the cart.
	}
}