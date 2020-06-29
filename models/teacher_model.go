package models

type Teacher struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type TeacherUseCase interface {
	Fetch() ([]*Teacher, error)
	GetByID(string) (Teacher, error)
	Store(Teacher) error
	Update(Teacher) error
	Delete(string) error
}

type TeacherRepository interface {
	Fetch() ([]*Teacher, error)
	GetByID(string) (Teacher, error)
	Store(Teacher) error
	Update(Teacher) error
	Delete(string) error
}
