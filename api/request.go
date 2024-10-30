package api

import (
	"fmt"
)

type StudentRequest struct {
	Name 		string 	`json:"name"`
	CPF 		int 		`json:"cpf"`
	Email 	string 	`json:"email"`
	Age 		int 		`json:"age"`
	Active 	*bool 	`json:"active"` // bool as a pointer to True or False
}

func errParamRequired(param, typ string) error {
	 return fmt.Errorf("param '%s' of type  '%s' is required", param, typ)
}

func(s *StudentRequest) Validade() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.CPF == 0 {
		return errParamRequired("cpf", "int")
	}
	if s.Email == "" {
		return	errParamRequired("email", "string")
 }
	if s.Age == 0 {
		return errParamRequired("age", "int")
}
	if s.Active == nil {
		return errParamRequired("active", "bool")
}

	return nil 
}