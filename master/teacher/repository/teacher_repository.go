package repositoriest

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
)

type teacherRepo struct {
	db *sql.DB
}

func NewTeacherRepo(db *sql.DB) models.TeacherRepository {
	return &teacherRepo{db}
}

func (t teacherRepo) Fetch() ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	rows, err := t.db.Query(`SELECT * FROM teacher`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Teacher)
		err := rows.Scan(&each.ID, &each.FirstName, &each.LastName, &each.Email)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, each)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teachers, nil
}

func (t teacherRepo) GetByID(id string) (*models.Teacher, error) {
	var teacher = new(models.Teacher)
	err := t.db.QueryRow("SELECT * FROM teacher WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (t teacherRepo) Store(teacher *models.Teacher) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO teacher VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(uuid.New(), teacher.FirstName, teacher.LastName, teacher.Email)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (t teacherRepo) Update(id string, teacher *models.Teacher) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE teacher
									SET first_name = ?, last_name = ?, email = ?
									WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(teacher.FirstName, teacher.LastName, teacher.Email, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (t teacherRepo) Delete(id string) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM teacher WHERE id = ?")
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
