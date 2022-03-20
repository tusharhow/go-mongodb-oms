package models


type Employee struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}