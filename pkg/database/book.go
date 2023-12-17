package database

import (
	"fmt"
	"github.com/ddeshi/library/model"
	"github.com/jinzhu/gorm"
)

func AddBook(book model.Book) error {
	fmt.Println(book)
	//var result model.Book
	exist, err := IsBookExist(book)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	if !exist && err == nil {
		err := db.Create(&book).Error
		if err != nil {
			fmt.Println("save user error %v", err)
		}
	} else {
		return fmt.Errorf("该书籍已存在")
	}

	return nil
}

func DeleteBook(book model.Book) error {
	err := db.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(book model.Book) error {
	condition := model.Book{
		BookName: book.BookName,
		Author:   book.Author,
	}
	// 执行更新操作
	err := db.Model(&model.Book{}).Where(&condition).Updates(book).Error
	if err != nil {
		return err
	}

	return nil
}

func SearchBook(book model.Book) ([]model.Book, error) {
	var (
		result []model.Book
		err    error
	)
	condition := model.Book{
		BookName: book.BookName,
	}

	// 执行查询操作
	err = db.Where(&condition).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil

}

func IsBookExist(book model.Book) (bool, error) {
	var result model.Book
	err := db.Where("bookName = ? AND author = ?", book.BookName, book.Author).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 书不存在
			return false, nil
		} else {
			// 其他错误
			return false, err
		}
	}
	// 书存在
	fmt.Println("该书籍已存在")
	return true, nil
}

func IsBookExistByName(book model.Book) (bool, error) {
	var result model.Book
	err := db.Where("bookName = ?", book.BookName).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 书不存在
			return false, nil
		} else {
			// 其他错误
			return false, err
		}
	}
	// 书存在
	//fmt.Println("该书籍已存在")
	return true, nil
}
