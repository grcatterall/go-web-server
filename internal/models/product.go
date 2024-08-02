package models

type Product struct {
	ID          string
	Name        string
	Price       float32
	Description string
	Slug 		string
}

func (p Product) GetName() string {
	return p.Name
}

func (p Product) GetSlug() string {
	return p.Slug
}

func (p Product) GetDescription() string {
	return p.Description
}

func (p Product) GetPrice() float32 {
	return p.Price
}
