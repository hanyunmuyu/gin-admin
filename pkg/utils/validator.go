package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslation "github.com/go-playground/validator/v10/Translations/zh"
	"reflect"
	"regexp"
	"sync"
)

var (
	uni   *ut.UniversalTranslator
	Trans ut.Translator
)

func mobile(fl validator.FieldLevel) bool {
	mobile, ok := fl.Field().Interface().(string)
	if ok {
		reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
		rgx := regexp.MustCompile(reg)
		return rgx.MatchString(mobile)
	}
	return false
}

func init() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	Trans, _ = uni.GetTranslator("zh")
}

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}
func (v *DefaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		Trans.Add("mobile", "{0}必须是一个有效的邮箱", true)
		_ = zhTranslation.RegisterDefaultTranslations(v.validate, Trans)
		_ = v.validate.RegisterValidation("mobile", mobile)
		v.validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
