package entity

var Yaml = &yamlData{}

type yamlData struct {
	Access  [][]string
	Develop struct {
		UserId int64 `yaml:"user_id,omitempty"`
	}
}
