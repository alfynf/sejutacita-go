package models

type User struct {
	Name     string `bson:"name" json:"name" form:"name"`
	Username string `bson:"username" json:"username" form:"username"`
	Email    string `bson:"email" json:"email" form:"email"`
	Password string `bson:"password" json:"password" form:"password"`
	Role     string `bson:"role" json:"role" form:"role"`
	Token    string `bson:"token" json:"token" form:"token"`
}

type GetUser struct {
	Name     string `bson:"name" json:"name" form:"name"`
	Username string `bson:"username" json:"username" form:"username"`
	Email    string `bson:"email" json:"email" form:"email"`
	Role     string `bson:"role" json:"role" form:"role"`
}
