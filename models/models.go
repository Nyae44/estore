package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	First_Name      string             `bson:"first_name" json:"first_name" validate:"required, min=2, max=30"`
	Last_Name       string             `bson:"last_name" json:"last_name" validate:"required, min=2, max=30"`
	Email           string             `bson:"email" json:"email" validate:"required,email"`
	Password        string             `bson:"password" json:"password" validate:"required, min=6"`
	Phone           string             `bson:"phone" json:"phone" validate:"required,number"`
	Token           string             `bson:"token" json:"token"`
	Refresh_Token   string             `bson:"refresh_token" json:"refresh_token"`
	Created_At      time.Time          `bson:"created_at" json:"created_at"`
	Updated_At      time.Time          `bson:"updated_at" json:"updated_at"`
	User_ID         string             `bson:"user_id" json:"user_id"`
	UserCart        []ProductUser      `bson:"user_cart" json:"user_cart"`
	Address_Details []Address
	Order_Status    []Order
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id,omitempty"`
	Product_Name string             `bson:"product_name" json:"product_name"`
	Price        uint64             `bson:"price" json:"price"`
	Rating       uint8              `bson:"rating" json:"rating"`
	Image        string             `bson:"image" json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id,omitempty"`
	Product_Name string             `bson:"product_name" json:"product_name"`
	Price        int64              `bson:"price" json:"price"`
	Rating       uint8              `bson:"rating" json:"rating"`
	Image        string             `bson:"image" json:"image"`
}

type Address struct {
	Address_ID  primitive.ObjectID `bson:"_id,omitempty"`
	House       string             `bson:"house" json:"house"`
	Street      string             `bson:"street" json:"street"`
	City        string             `bson:"city" json:"city"`
	Postal_Code string             `bson:"postal_code" json:"postal_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id,omitempty"`
	Order_Cart     []ProductUser      `bson:"order_cart" json:"order_cart"`
	Ordered_At     time.Time          `bson:"ordered_at" json:"ordered_at"`
	Price          uint64             `bson:"price" json:"price"`
	Discount       int                `bson:"discount" json:"discount"`
	Payment_Method Payment
}

type Payment struct {
	Mpesa            bool
	Cash_On_Delivery bool
}
