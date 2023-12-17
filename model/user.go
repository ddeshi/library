package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Email    string `json:"email" gorm:"column:email"`
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
	IsAdmin  bool   `json:"isAdmin" gorm:"column:isAdmin"`
}

func (r *User) TableName() string {
	return "user"
}
