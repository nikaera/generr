package info

import (
	"fmt"
	"strconv"
)

type Error struct {
	Name string            `yaml:"name"`
	Msg  string            `yaml:"message"`
	Key  *string           `yaml:"key"`
	Desc *string           `yaml:"description"`
	Url  *string           `yaml:"url"`
	Meta map[string]string `yaml:"metadata"`
}

type ErrorWithCode struct {
	Code  Code
	Error Error
}

func ParseErrors(ecs []*ErrorWithCode, code Code, infos []Info) []*ErrorWithCode {
	k := func(i int, isGroup bool) string {
		n := i + 1
		if isGroup {
			return strconv.Itoa(n)
		} else {
			return fmt.Sprintf("%03d", n)
		}
	}

	for i, v := range infos {
		c := code.GenCode(k(i, true), v.Name)
		if v.Errors != nil {
			for j, e := range *v.Errors {
				cc := c.GenCode(k(j, false), e.Name)
				ecs = append(ecs, &ErrorWithCode{
					Error: e,
					Code:  cc,
				})
			}
		}

		if v.Groups != nil {
			ecs = ParseErrors(ecs, c, *v.Groups)
		}
	}

	return ecs
}
