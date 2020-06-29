package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
)

type studentRepo struct {
	DB *sql.DB
}

func NewStudentRepo(db *sql.DB) models.StudentRepository {
	return &studentRepo{db}
}

func (s studentRepo) Fetch() ([]*models.Student, error) {
	var students []*models.Student
	rows, err := s.DB.Query(`SELECT * FROM student`)
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

func (s studentRepo) GetByID(id string) (models.Student, error) {
	var student models.Student
	err := s.DB.QueryRow("SELECT id, first_name, last_name, email FROM student WHERE id = ?", id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (s studentRepo) Store(student models.Student) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO student VALUES (?, ?, ?, ?)")
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

func (s studentRepo) Update(student models.Student) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE student
									SET first_name = ?, last_name = ?, email = ?
									WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(student.FirstName, student.LastName, student.Email, student.ID)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (s studentRepo) Delete(id string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM student WHERE id = ?")
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
