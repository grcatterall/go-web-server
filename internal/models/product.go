package models

type Product struct {
	ID          string
	Name        string
	Price       float32
	Description string
}

func (p Product) GetName() string {
	return p.Name
}
