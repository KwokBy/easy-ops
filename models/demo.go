package models

type Demo struct {
	Name string `json:"name"`
}

func (d *Demo)TableName() string {
	return "t_demo"
}