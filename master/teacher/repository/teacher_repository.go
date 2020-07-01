package repositoriest

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
	"github.com/vivaldy22/cleanEnigmaSchool/tools/queries"
)

type teacherRepo struct {
	db *sql.DB
}

func NewTeacherRepo(db *sql.DB) models.TeacherRepository {
	return &teacherRepo{db}
}

func (t teacherRepo) Fetch() ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	rows, err := t.db.Query(queries.SELECT_ALL_TEACHER)
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
	err := t.db.QueryRow(queries.SELECT_TEACHER_ID, id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email)
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

	stmt, err := tx.Prepare(queries.INSERT_TEACHER)
	if err != nil {
		return err
	}

	teacher.ID = uuid.New().String()
	_, err = stmt.Exec(teacher.ID, teacher.FirstName, teacher.LastName, teacher.Email)
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

	stmt, err := tx.Prepare(queries.UPDATE_TEACHER)
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

	stmt, err := tx.Prepare(queries.DELETE_TEACHER_ID)
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
