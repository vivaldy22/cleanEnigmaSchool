package models

type Subject struct {
	ID          string `json:"id"`
	SubjectName string `json:"subject"`
}

type SubjectUseCase interface {
	Fetch() ([]*Subject, error)
	GetByID(string) (Subject, error)
	Store(Subject) error
	Update(Subject) error
	Delete(string) error
}

type SubjectRepository interface {
	Fetch() ([]*Subject, error)
	GetByID(string) (Subject, error)
	Store(Subject) error
	Update(Subject) error
	Delete(string) error
}
