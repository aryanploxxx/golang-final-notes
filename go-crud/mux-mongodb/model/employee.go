package model

type Employee struct {
	EmployeeID string `json:"employee_id,omniempty" bson:"employee_id"`
	Name       string `json:"name,omniempty" bson:"name"`
	Department string `json:"department,omniempty" bson:"department"`
}

// bson, because mongodb does not understand json
