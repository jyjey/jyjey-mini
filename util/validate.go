package util

import (
	"fmt"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

//CheckStruct 校验Strust
func CheckStruct(s interface{}) (string, error) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
			return err.Translate(trans), err
		}
	}
	return "", nil
}
