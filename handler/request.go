package handler

import "fmt"

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("location", "boolean")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("location", "boolean")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) ValidateUpdate() error {

	if r.Link != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Role != "" || r.Salary <= 0 {
		return nil
	}

	return fmt.Errorf("some one param must be provid")
}
