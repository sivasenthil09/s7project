package models

type User struct {
	Name string `json:"name" bson:"name"`
	Password string`json:"password" bson:"password"`
	ConfirmPassword string `json:"confirm" bson:"confirm"`
	PhoneNumber string `json:"phonenumber" bson:"phonenumber"`
	Email string `json:"email" bson:"email"`
	Address Address `json:"address" bson:"address"`
	
}
type Address struct {
	Area string `json:"area" bson:"area"`
	City string `json:"city" bson:"city"`
	Pincode string `json:"pincode" bson:"pincode"`
}