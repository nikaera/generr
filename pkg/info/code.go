package info

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	errPrefix     = "Err"
	errCodePrefix = "ErrCode"
)

type Code struct {
	Key  string
	Name string
}

func (c *Code) ErrorNum() int {
	num, _ := strconv.Atoi(c.Key)
	return num
}

func (c *Code) NameWithCode() string {
	return strings.Replace(c.Name, errPrefix, errCodePrefix, 1)
}

func (c *Code) GenCode(key string, name string) Code {
	var k string
	if c.Key == "" {
		k = key
	} else {
		k = c.Key + key
	}

	var n string
	if c.Name == "" {
		n = errPrefix
	} else {
		n = c.Name
	}
	return Code{
		Key:  k,
		Name: strcase.ToCamel(fmt.Sprintf("%s-%s", n, name)),
	}
}
