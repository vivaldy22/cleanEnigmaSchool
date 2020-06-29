package usecase

import (
	"github.com/vivaldy22/cleanEnigmaSchool/models"
)

type teacherUseCase struct {
	teacherRepo models.TeacherRepository
}

func NewTeacherUseCase(t models.TeacherRepository) models.TeacherUseCase {
	return &teacherUseCase{teacherRepo: t}
}

func (t teacherUseCase) Fetch() ([]*models.Teacher, error) {
	res, err := t.teacherRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t teacherUseCase) GetByID(id string) (models.Teacher, error) {
	res, err := t.teacherRepo.GetByID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t teacherUseCase) Store(teacher models.Teacher) error {
	return t.teacherRepo.Store(teacher)
}

func (t teacherUseCase) Update(teacher models.Teacher) error {
	return t.teacherRepo.Update(teacher)
}

func (t teacherUseCase) Delete(id string) error {
	return t.teacherRepo.Delete(id)
}
