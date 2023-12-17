package User

import (
	"fmt"
	"github.com/ddeshi/library/model"
	"github.com/ddeshi/library/pkg/Service"
	"github.com/ddeshi/library/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Login user login
func Login(c *gin.Context) {
	var (
		reqBody model.User
		res     = &model.HTTPRespResult{}
	)
	//defer c.Set(model.GinResponseKey, res)

	res.Code = Service.BindAndValid(c, &reqBody)

	fmt.Println(reqBody)

	loginUser, exist := database.GetUser(reqBody)
	if !exist {
		fmt.Printf("账号密码错误或该用户未注册")
		c.JSON(400, gin.H{
			"msg": "账号密码错误或该用户未注册",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":      "登录成功",
		"userInfo": loginUser,
	})
}

// Register new user register
func Register(c *gin.Context) {
	var (
		reqBody model.User
		res     = &model.HTTPRespResult{}
	)

	defer c.Set(model.GinResponseKey, res)

	res.Code = Service.BindAndValid(c, &reqBody)
	if res.Code != 200 {
		res.Err = fmt.Errorf("failed to unmarshal request")
		res.Msg = "failed to unmarshal request"
	}

	err := database.CheckUserInfo(reqBody)
	if err != nil {
		logrus.Errorf("errerr")
		c.JSON(400, gin.H{
			"msg": "用户已存在",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "新用户注册成功",
	})

	fmt.Println(reqBody)
}
