package models

type Student struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type StudentUseCase interface {
	Fetch() ([]*Student, error)
	GetByID(string) (*Student, error)
	Store(*Student) error
	Update(string, *Student) error
	Delete(string) error
}

type StudentRepository interface {
	Fetch() ([]*Student, error)
	GetByID(string) (*Student, error)
	Store(*Student) error
	Update(string, *Student) error
	Delete(string) error
}
