package model

type User struct {
	BaseEntity `bson:",inline" json:",inline"`
	Username   string `bson:"name" json:"name"`
	Password   string `bson:"password" json:"-"`
	Email      string `bson:"email" json:"email"`
}
