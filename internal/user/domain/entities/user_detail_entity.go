package entities

// UserDetailEntity ...
type UserDetailEntity struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id,omitempty"`
	Email       string `json:"email" validate:"required"`
	Name        string `json:"name" validate:"required"`
	DateOfBirth string `json:"date_of_brith" validate:"required"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
}
