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

func (t teacherUseCase) GetByID(id string) (*models.Teacher, error) {
	res, err := t.teacherRepo.GetByID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t teacherUseCase) Store(teacher *models.Teacher) error {
	if err := t.teacherRepo.Store(teacher); err != nil {
		return err
	}
	return nil
}

func (t teacherUseCase) Update(id string, teacher *models.Teacher) error {
	if err := t.teacherRepo.Update(id, teacher); err != nil {
		return err
	}
	return nil
}

func (t teacherUseCase) Delete(id string) error {
	if err := t.teacherRepo.Delete(id); err != nil {
		return err
	}
	return nil
}
