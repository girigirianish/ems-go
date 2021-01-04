package models

import "github.com/girigirianish/ems-go/internal/user/domain/entities"

// UserDetailsReqResponse godoc
type UserDetailsReqResponse struct {
	PersonalDetails   *entities.UserDetailEntity      `json:"personal_details"`
	EducationDetails  []entities.UserEducationDetail  `json:"education_details"`
	ExperienceDetails []entities.UserExperienceDetail `json:"experience_details"`
}
