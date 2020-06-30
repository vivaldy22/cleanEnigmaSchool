package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vivaldy22/cleanEnigmaSchool/models"
)

type subjectRepo struct {
	db *sql.DB
}

func NewSubjectRepo(db *sql.DB) models.SubjectRepository {
	return &subjectRepo{db}
}

func (s subjectRepo) Fetch() ([]*models.Subject, error) {
	var subjects []*models.Subject
	rows, err := s.db.Query(`SELECT * FROM subject`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Subject)
		err := rows.Scan(&each.ID, &each.SubjectName)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, each)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return subjects, nil
}

func (s subjectRepo) GetByID(id string) (*models.Subject, error) {
	var subject = new(models.Subject)
	err := s.db.QueryRow("SELECT * FROM subject WHERE id = ?", id).Scan(&subject.ID, &subject.SubjectName)
	if err != nil {
		return subject, err
	}
	return subject, nil
}

func (s subjectRepo) Store(subject *models.Subject) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO subject VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(uuid.New(), subject.SubjectName)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (s subjectRepo) Update(id string, subject *models.Subject) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE subject
									SET subject_name = ?
									WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(subject.SubjectName, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (s subjectRepo) Delete(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM subject WHERE id = ?")
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
