package database

import (
	"fmt"
	"github.com/ddeshi/library/model"
	"github.com/jinzhu/gorm"
)

func CheckUserInfo(user model.User) error {
	fmt.Println(user)
	var result model.User
	var err error
	err = db.Where("readerName = ? AND password = ? AND email = ?", user.Name, user.Phone, user.Email).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//save userInfo
			err := db.Create(&user).Error
			if err != nil {
				fmt.Println("save user error %v", err)
			}
		} else {
			return err
		}
	}

	return nil
}

func GetUser(user model.User) (model.User, bool) {
	var result model.User
	err := db.Where("phone = ?", user.Phone).First(&result).Error
	if err != nil {
		fmt.Printf("get userInfo error: %s\n", err)
		return result, false
	}
	if result.Password != user.Password {
		fmt.Printf("user password error: \n")
		return result, false
	}
	return result, true
}
