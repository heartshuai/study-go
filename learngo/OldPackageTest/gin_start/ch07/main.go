package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

var trans ut.Translator

type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required,min=3,max=10"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
type SignUpForm struct {
	Age        uint8  `form:"age" json:"age" xml:"age" binding:"gte=1,lte=130"`
	Name       string `form:"name" json:"name" xml:"name" binding:"required,min=3"`
	Email      string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required"`
	RePassword string `form:"repassword" json:"repassword" xml:"repassword" binding:"required,eqfield=Password"` //跨字段
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境 后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		}

	}
	return
}

func main() {
	err := InitTrans("zh")
	if err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	router := gin.Default()
	router.POST("/loginJSON", func(c *gin.Context) {
		var loginForm LoginForm

		if err := c.ShouldBind(&loginForm); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": removeTopStruct(errs.Translate(trans)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"user":     loginForm.User,
				"password": loginForm.Password,
			})
		}
	})
	router.POST("/signUp", func(c *gin.Context) {
		var signUpForm SignUpForm
		if err := c.ShouldBind(&signUpForm); err != nil {
			var errs validator.ValidationErrors
			ok := errors.As(err, &errs)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": removeTopStruct(errs.Translate(trans)),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name":     signUpForm.Name,
				"email":    signUpForm.Email,
				"password": signUpForm.Password,
			})
		}
	})
	router.Run()
}
