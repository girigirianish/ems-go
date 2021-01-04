package entities

// UserExperienceDetail ...
type UserExperienceDetail struct {
	ID                 int64  `json:"id"`
	UserID             int64  `json:"user_id,omitempty"`
	IsCurrentJob       string `json:"is_current_job"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	CompanyName        string `json:"company_name"`
	JobLocationCity    string `json:"job_location_city"`
	JobLocationState   string `json:"job_location_state"`
	JobLocationCountry string `json:"job_location_country"`
	UpdatedAt          int64  `json:"updated_at,omitempty"`
	CreatedAt          int64  `json:"created_at,omitempty"`
}
