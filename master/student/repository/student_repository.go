package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
	"github.com/vivaldy22/cleanEnigmaSchool/tools/queries"
)

type studentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) models.StudentRepository {
	return &studentRepo{db}
}

func (s studentRepo) Fetch() ([]*models.Student, error) {
	var students []*models.Student
	rows, err := s.db.Query(queries.SELECT_ALL_STUDENT)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Student)
		err := rows.Scan(&each.ID, &each.FirstName, &each.LastName, &each.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, each)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}

func (s studentRepo) GetByID(id string) (*models.Student, error) {
	var student = new(models.Student)
	err := s.db.QueryRow(queries.SELECT_STUDENT_ID, id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (s studentRepo) Store(student *models.Student) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.INSERT_STUDENT)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(uuid.New(), student.FirstName, student.LastName, student.Email)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (s studentRepo) Update(id string, student *models.Student) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.UPDATE_STUDENT)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(student.FirstName, student.LastName, student.Email, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (s studentRepo) Delete(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(queries.DELETE_STUDENT_ID)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}
