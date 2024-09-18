package model

type User struct {
	BaseEntity `bson:",inline" json:",inline"`
	Username   string `bson:"username" json:"username"`
	Password   string `bson:"password" json:"-"`
	Email      string `bson:"email" json:"email"`
}
