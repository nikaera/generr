package info

type Info struct {
	Name   string   `yaml:"name"`
	Groups *[]Info  `yaml:"groups"`
	Errors *[]Error `yaml:"errors"`
}
