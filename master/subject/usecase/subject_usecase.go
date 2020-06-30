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

func (s subjectUseCase) GetByID(id string) (models.Subject, error) {
	res, err := s.subjectRepo.GetByID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s subjectUseCase) Store(subject models.Subject) error {
	return s.subjectRepo.Store(subject)
}

func (s subjectUseCase) Update(subject models.Subject) error {
	return s.subjectRepo.Update(subject)
}

func (s subjectUseCase) Delete(id string) error {
	return s.subjectRepo.Delete(id)
}
