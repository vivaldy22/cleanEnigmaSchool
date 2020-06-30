package usecase

import (
	"github.com/vivaldy22/cleanEnigmaSchool/models"
)

type subjectUseCase struct {
	subjectRepo models.SubjectRepository
}

func NewSubjectUseCase(s models.SubjectRepository) models.SubjectUseCase {
	return &subjectUseCase{subjectRepo: s}
}

func (s subjectUseCase) Fetch() ([]*models.Subject, error) {
	res, err := s.subjectRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s subjectUseCase) GetByID(id string) (*models.Subject, error) {
	res, err := s.subjectRepo.GetByID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s subjectUseCase) Store(subject *models.Subject) error {
	if err := s.subjectRepo.Store(subject); err != nil {
		return err
	}
	return nil
}

func (s subjectUseCase) Update(id string, subject *models.Subject) error {
	if err := s.subjectRepo.Update(id, subject); err != nil {
		return err
	}
	return nil
}

func (s subjectUseCase) Delete(id string) error {
	if err := s.subjectRepo.Delete(id); err != nil {
		return err
	}
	return nil
}
