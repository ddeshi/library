package Book

import (
	"fmt"
	"github.com/ddeshi/library/model"
	"github.com/ddeshi/library/pkg/Service"
	"github.com/ddeshi/library/pkg/database"
	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {
	var (
		reqBody model.Book
		res     = &model.HTTPRespResult{}
	)

	res.Code = Service.BindAndValid(c, &reqBody)
	if res.Code != 200 {
		res.Err = fmt.Errorf("failed to unmarshal request")
		res.Msg = "failed to unmarshal request"
	}
	if err := database.AddBook(reqBody); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"msg": "该书籍已存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":      "该书籍添加成功",
		"bookInfo": reqBody,
	})
	fmt.Println(reqBody)
}

func DeleteBook(c *gin.Context) {
	var (
		reqBody model.Book
		res     = &model.HTTPRespResult{}
		err     error
	)

	res.Code = Service.BindAndValid(c, &reqBody)
	if res.Code != 200 {
		res.Err = fmt.Errorf("failed to unmarshal request")
		res.Msg = "failed to unmarshal request"
	}

	isExist, err := database.IsBookExist(reqBody)
	if err != nil {
		res.Code = 400
		res.Err = err
		res.Msg = "database search error"
		Service.Msg(res, c)
		return
	}
	if !isExist {
		fmt.Println("the boook is not exist")
		res.Code = 400
		res.Err = fmt.Errorf("the boook is not exist")
		res.Msg = "the boook is not exist"
		Service.Msg(res, c)
		return
	}

	err = database.DeleteBook(reqBody)
	if err != nil {
		res.Code = 400
		res.Err = err
		res.Msg = "delete book error"
		Service.Msg(res, c)
		return
	}
	c.JSON(200, gin.H{
		"msg":      "该书籍删除成功",
		"bookInfo": reqBody,
	})
}

func UpdateBook(c *gin.Context) {
	var (
		reqBody model.Book
		res     = &model.HTTPRespResult{}
		err     error
	)

	res.Code = Service.BindAndValid(c, &reqBody)
	if res.Code != 200 {
		res.Err = fmt.Errorf("failed to unmarshal request")
		res.Msg = "failed to unmarshal request"
	}

	isExist, err := database.IsBookExist(reqBody)
	if err != nil {
		res.Code = 400
		res.Err = err
		res.Msg = "database search error"
		Service.Msg(res, c)
		return
	}
	if !isExist {
		fmt.Println("the boook is not exist")
		res.Code = 400
		res.Err = fmt.Errorf("the boook is not exist")
		res.Msg = "the boook is not exist"
		Service.Msg(res, c)
		return
	}

	err = database.UpdateBook(reqBody)
	if err != nil {
		res.Code = 400
		res.Err = fmt.Errorf("update boook information error")
		res.Msg = "update boook information error"
		Service.Msg(res, c)
		return
	}

	c.JSON(200, gin.H{
		"msg":      "该书籍信息更新成功",
		"bookInfo": reqBody,
	})

}

func SearchBook(c *gin.Context) {
	var (
		reqBody model.Book
		res     = &model.HTTPRespResult{}
		err     error
	)

	res.Code = Service.BindAndValid(c, &reqBody)
	if res.Code != 200 {
		res.Err = fmt.Errorf("failed to unmarshal request")
		res.Msg = "failed to unmarshal request"
	}

	isExist, err := database.IsBookExistByName(reqBody)
	if err != nil {
		fmt.Println(err)
		res.Code = 400
		res.Err = err
		res.Msg = "database search error"
		Service.Msg(res, c)
		return
	}
	if !isExist {
		fmt.Println("the boook is not exist")
		res.Code = 400
		res.Err = fmt.Errorf("the boook is not exist")
		res.Msg = "the boook is not exist"
		Service.Msg(res, c)
		return
	}
	//var books []model.Book
	books, err := database.SearchBook(reqBody)
	if err != nil {
		res.Code = 400
		res.Err = err
		res.Msg = "database search error"
		Service.Msg(res, c)
		return
	}

	c.JSON(200, gin.H{
		"msg":      "该书籍查询成功",
		"bookInfo": books,
	})

}
