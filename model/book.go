package model

type Book struct {
	BookID   int    `json:"bookID" gorm:"column:bookID;primary_key"`
	BookName string `json:"bookName" gorm:"column:bookName"`
	Author   string `json:"author" gorm:"column:author"`
	// 当前数量
	Amount int `json:"amount" gorm:"column:amount"`
	// 总数量
	TotalAmount int `json:"totalAmount" gorm:"column:total_amount"`
	// 图书位置
	Position string `json:"position" gorm:"column:positon"`
	// 可借阅状态，1可借阅， 0不可
	Status int `json:"status" gorm:"column:status"`
	// 借阅次数
	BorrowedTimes int `json:"borrowedTimes" gorm:"column:borrowed_times"`
}

func (b *Book) TableName() string {
	return "book"
}
