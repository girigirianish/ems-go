package entities

// UserEntity ...
type UserEntity struct {
	ID        int64  `json:"id"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	RoleID    int64  `json:"role_id,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
