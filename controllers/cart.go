package controllers

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProduct  = errors.New("can't find  product")
	ErrUserIdIsNotValid   = errors.New("user id isn't valid")
	ErrCantUpdateUser     = errors.New("can't update user")
	ErrCantRemoveItemCart = errors.New("can't remove item from cart")
	ErrCantGetItem        = errors.New("can't get item from cart")
	ErrCantBuyItem        = errors.New("can't update purchase")
)

func AddToCart() {

}

func RemoveFromCart() {

}
func GetItemFromCart() {

}

func BuyFromCart() {

}

func InstantBuy() {

}
