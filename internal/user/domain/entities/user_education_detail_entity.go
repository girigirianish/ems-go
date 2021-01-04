package entities

// UserEducationDetail ...
type UserEducationDetail struct {
	ID                      int64  `json:"id"`
	UserID                  int64  `json:"user_id,omitempty"`
	CertificateDegreeName   string `json:"certificate_degree_name"`
	Major                   string `json:"major"`
	InstituteUniversityName string `json:"institute_university_name"`
	StartingDate            string `json:"starting_date"`
	EndDate                 string `json:"end_date"`
	Percentage              int64  `json:"percentage"`
	Cgpa                    int64  `json:"cgpa"`
	UpdatedAt               int64  `json:"updated_at,omitempty"`
	CreatedAt               int64  `json:"created_at,omitempty"`
}
