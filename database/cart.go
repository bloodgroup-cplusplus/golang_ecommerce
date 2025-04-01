package database

import "errors"

var (
	ErrCantFindProduct = errors.New("Can't Find the product")
	ErrCantDecodeProducts = errors.New("Cant Find the Product")
	ErrUserIdIsNotValid = errors.New("This User is not Valid")
	ErrCantUpdateUser = errors.New("Cannot add this product to the cart")
	ErrCantRemoveItemCart =errors.New("Cannot remove this item from the cart")
	ErrCantGetItem = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func IntantBuyer() {

}