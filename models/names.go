// models/names.go
package models

type NameList struct {
	NameF string  `json:"fName"`
	NameL string  `json:"lName"`
	Age   float32 `json:"Age"`
}

var Names = []NameList{
	{NameF: "Quandale", NameL: "Dingle the third", Age: 59},
	{NameF: "James", NameL: "Howlet", Age: 134},
	{NameF: "Bihan", NameL: "Tundra", Age: 38},
}
