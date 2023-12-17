package Service

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/ddeshi/library/model"
	"github.com/gin-gonic/gin"
	"regexp"

	"net/http"
)

func BindAndValid(c *gin.Context, form interface{}) int {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError
	}
	if !check {
		//MarkErrors(valid.Errors)
		return http.StatusBadRequest
	}

	return http.StatusOK
}

const (
	emailRegex = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	phoneRegex = `^1[3456789]\d{9}$`
)

func EmailRegexp(email string) error {
	// 编译正则表达式
	emailRegexp := regexp.MustCompile(emailRegex)
	// 匹配字符串
	if !emailRegexp.MatchString(email) {
		return errors.New("邮箱正则表达式匹配失败")
	}
	return nil
}

func PhoneRegexp(phone string) error {
	// 编译正则表达式
	phoneRegexp := regexp.MustCompile(phoneRegex)
	// 匹配字符串
	if !phoneRegexp.MatchString(phone) {
		return errors.New("手机号正则表达式匹配失败")
	}
	return nil
}

func Msg(res *model.HTTPRespResult, c *gin.Context) {
	c.JSON(res.Code, gin.H{
		"msg": res.Msg,
	})
}
