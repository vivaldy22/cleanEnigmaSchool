package usecase

import (
	"github.com/vivaldy22/cleanEnigmaSchool/models"
)

type studentUseCase struct {
	studentRepo models.StudentRepository
}

func NewStudentUseCase(s models.StudentRepository) models.StudentUseCase {
	return &studentUseCase{studentRepo: s}
}

func (s studentUseCase) Fetch() ([]*models.Student, error) {
	res, err := s.studentRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s studentUseCase) GetByID(id string) (*models.Student, error) {
	res, err := s.studentRepo.GetByID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s studentUseCase) Store(student *models.Student) error {
	if err := s.studentRepo.Store(student); err != nil {
		return err
	}
	return nil
}

func (s studentUseCase) Update(id string, student *models.Student) error {
	if err := s.studentRepo.Update(id, student); err != nil {
		return err
	}
	return nil
}

func (s studentUseCase) Delete(id string) error {
	if err := s.studentRepo.Delete(id); err != nil {
		return err
	}
	return nil
}
