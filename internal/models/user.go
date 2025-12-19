package models

// type CreateUserRequest struct {
// 	Name string `json:"name"`
// 	Dob  string `json:"dob"` // YYYY-MM-DD
// }

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	Dob  string `json:"dob"  validate:"required,datetime=2006-01-02"`
}

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age"`
}
