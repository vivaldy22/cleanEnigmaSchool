package models

type Subject struct {
	ID          string `json:"id"`
	SubjectName string `json:"subjectName"`
}

type SubjectUseCase interface {
	Fetch() ([]*Subject, error)
	GetByID(string) (*Subject, error)
	Store(*Subject) error
	Update(string, *Subject) error
	Delete(string) error
}

type SubjectRepository interface {
	Fetch() ([]*Subject, error)
	GetByID(string) (*Subject, error)
	Store(*Subject) error
	Update(string, *Subject) error
	Delete(string) error
}
