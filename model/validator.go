package model

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

// CustomValidator 自定义校验器
type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// Validate 实现Echo 的Validate接口
func (c *CustomValidator) Validate(i interface{}) error {
	c.lazyInit()
	return c.validate.Struct(i)
}

func (c *CustomValidator) lazyInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}
