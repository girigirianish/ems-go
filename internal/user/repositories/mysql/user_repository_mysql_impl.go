package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	genericerrors "github.com/girigirianish/ems-go/internal/generic_errors"
	"github.com/girigirianish/ems-go/internal/user/domain/entities"
	"github.com/girigirianish/ems-go/internal/user/domain/repositories"
	"github.com/labstack/gommon/log"
)

type mysqlRepository struct {
	conn *sql.DB
}

// Init will create an object that represent the category's Repository interface
func Init(db *sql.DB) repositories.UserRepository {
	return &mysqlRepository{
		conn: db,
	}
}

func (m *mysqlRepository) CreateUser(ctx context.Context, user *entities.UserEntity) (err error) {
	query := `INSERT users SET email=?,password=?,role_id=?,updated_at=?,created_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("Error while preparing statement ", err)
		return
	}
	user.UpdatedAt = time.Now().Unix()
	user.CreatedAt = time.Now().Unix()
	user.RoleID = 2
	res, err := stmt.ExecContext(ctx, user.Email, user.Password, user.RoleID, user.UpdatedAt, user.CreatedAt)
	if err != nil {
		log.Error("Error while executing statement ", err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Error("Got Error from LastInsertId method: ", err)
		return
	}
	user.ID = lastID
	return
}

func (m *mysqlRepository) GetUser(ctx context.Context, email string, password string) (*entities.UserEntity, error) {
	query := `SELECT id,email,password,role_id, updated_at,created_at FROM users WHERE email=? AND password=?`
	list, err := m.fetchUser(ctx, query, email, password)
	if err != nil {
		return nil, err
	}
	var user entities.UserEntity
	if len(list) > 0 {
		user = list[0]
	} else {
		return &user, genericerrors.ErrNotFound
	}

	return &user, nil
}

func (m *mysqlRepository) GetUserEducationDetail(ctx context.Context, id int64) (*[]entities.UserEducationDetail, error) {
	query := `SELECT id,user_id,certificate_degree_name,major,institute_university_name,starting_date,end_date,percentage,cgpa,updated_at,created_at FROM user_education_details WHERE user_id=?`
	usersDetail, err := m.fetchUserEduDetail(ctx, query, id)
	if err != nil {
		return nil, err
	}
	return &usersDetail, nil
}

func (m *mysqlRepository) GetUserExperienceDetail(ctx context.Context, id int64) (*[]entities.UserExperienceDetail, error) {
	query := `SELECT id,user_id,is_current_job,start_date,end_date,company_name,job_location_city,job_location_state,job_location_country,updated_at,created_at FROM user_experience_details  WHERE user_id=?`
	usersDetail, err := m.fetchUserExpDetail(ctx, query, id)
	if err != nil {
		return nil, err
	}
	return &usersDetail, nil
}

func (m *mysqlRepository) GetAllUsersDetail(ctx context.Context) (*[]entities.UserDetailEntity, error) {
	query := `SELECT id,user_id,name,email,date_of_birth,updated_at,created_at FROM user_details`
	usersDetail, err := m.fetchUserDetail(ctx, query)
	if err != nil {
		return nil, err
	}
	return &usersDetail, nil
}

func (m *mysqlRepository) fetchUser(ctx context.Context, query string, args ...interface{}) (result []entities.UserEntity, err error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]entities.UserEntity, 0)
	for rows.Next() {
		t := entities.UserEntity{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Password,
			&t.RoleID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRepository) fetchUserDetail(ctx context.Context, query string, args ...interface{}) (result []entities.UserDetailEntity, err error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]entities.UserDetailEntity, 0)
	for rows.Next() {
		t := entities.UserDetailEntity{}
		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Name,
			&t.Email,
			&t.DateOfBirth,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRepository) fetchUserExpDetail(ctx context.Context, query string, args ...interface{}) (result []entities.UserExperienceDetail, err error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]entities.UserExperienceDetail, 0)
	for rows.Next() {
		t := entities.UserExperienceDetail{}
		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.IsCurrentJob,
			&t.StartDate,
			&t.EndDate,
			&t.CompanyName,
			&t.JobLocationCity,
			&t.JobLocationState,
			&t.JobLocationCountry,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRepository) fetchUserEduDetail(ctx context.Context, query string, args ...interface{}) (result []entities.UserEducationDetail, err error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]entities.UserEducationDetail, 0)
	for rows.Next() {
		t := entities.UserEducationDetail{}
		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.CertificateDegreeName,
			&t.Major,
			&t.InstituteUniversityName,
			&t.StartingDate,
			&t.EndDate,
			&t.Percentage,
			&t.Cgpa,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRepository) CreateUserDetail(ctx context.Context, userDetail *entities.UserDetailEntity) (err error) {
	query := `INSERT user_details SET email=?,name=?,date_of_birth=?,user_id=?,updated_at=?,created_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("Error while preparing statement ", err)
		return err
	}
	userDetail.UpdatedAt = time.Now().Unix()
	userDetail.CreatedAt = time.Now().Unix()
	res, err := stmt.ExecContext(ctx, userDetail.Email, userDetail.Name, userDetail.DateOfBirth, userDetail.UserID, userDetail.UpdatedAt, userDetail.CreatedAt)
	if err != nil {
		log.Error("Error while executing statement ", err)
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Error("Got Error from LastInsertId method: ", err)
		return err
	}
	userDetail.ID = lastID
	return nil
}

func (m *mysqlRepository) CreateUserEducationDetails(ctx context.Context, userEduDetails []entities.UserEducationDetail) error {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, userDetail := range userEduDetails {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

		valueArgs = append(valueArgs, userDetail.UserID)
		valueArgs = append(valueArgs, userDetail.CertificateDegreeName)
		valueArgs = append(valueArgs, userDetail.Major)
		valueArgs = append(valueArgs, userDetail.InstituteUniversityName)
		valueArgs = append(valueArgs, userDetail.StartingDate)
		valueArgs = append(valueArgs, userDetail.EndDate)
		valueArgs = append(valueArgs, userDetail.Percentage)
		valueArgs = append(valueArgs, userDetail.Cgpa)
		valueArgs = append(valueArgs, userDetail.UpdatedAt)
		valueArgs = append(valueArgs, userDetail.CreatedAt)
	}
	smt := fmt.Sprintf(" INSERT INTO user_education_details (user_id,certificate_degree_name,major,institute_university_name,starting_date,end_date,percentage,cgpa,updated_at,created_at) VALUES %s", strings.Join(valueStrings, ","))
	fmt.Println("smttt:", smt)
	tx, beginErr := m.conn.Begin()
	if beginErr != nil {
		log.Error("Error while executing statement ", beginErr)
		return beginErr
	}
	_, err := tx.Exec(smt, valueArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (m *mysqlRepository) CreateUserExperienceDetails(ctx context.Context, userEduDetails []entities.UserExperienceDetail) error {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, userDetail := range userEduDetails {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

		valueArgs = append(valueArgs, userDetail.UserID)
		valueArgs = append(valueArgs, userDetail.IsCurrentJob)
		valueArgs = append(valueArgs, userDetail.StartDate)
		valueArgs = append(valueArgs, userDetail.EndDate)
		valueArgs = append(valueArgs, userDetail.CompanyName)
		valueArgs = append(valueArgs, userDetail.JobLocationCity)
		valueArgs = append(valueArgs, userDetail.JobLocationState)
		valueArgs = append(valueArgs, userDetail.JobLocationCountry)
		valueArgs = append(valueArgs, userDetail.UpdatedAt)
		valueArgs = append(valueArgs, userDetail.CreatedAt)
	}
	smt := fmt.Sprintf(" INSERT INTO user_experience_details (user_id,is_current_job,start_date,end_date,company_name,job_location_city,job_location_state,job_location_country,updated_at,created_at) VALUES %s", strings.Join(valueStrings, ","))
	fmt.Println("smttt:", smt)
	tx, beginErr := m.conn.Begin()
	if beginErr != nil {
		log.Error("Error while executing statement ", beginErr)
		return beginErr
	}
	_, err := tx.Exec(smt, valueArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
