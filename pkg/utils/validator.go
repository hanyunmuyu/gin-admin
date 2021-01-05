package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslation "github.com/go-playground/validator/v10/Translations/zh"
	"reflect"
	"sync"
)

var (
	uni   *ut.UniversalTranslator
	Trans ut.Translator
)

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
		_ = zhTranslation.RegisterDefaultTranslations(v.validate, Trans)
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
