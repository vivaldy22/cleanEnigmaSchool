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

func (s studentUseCase) GetByID(id string) (models.Student, error) {
	res, err := s.studentRepo.GetByID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s studentUseCase) Store(student models.Student) error {
	return s.studentRepo.Store(student)
}

func (s studentUseCase) Update(student models.Student) error {
	return s.studentRepo.Update(student)
}

func (s studentUseCase) Delete(id string) error {
	return s.studentRepo.Delete(id)
}
