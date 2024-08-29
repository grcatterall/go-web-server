package models

type Category struct {
	ID string
	Name string
	Description string
	Slug string
}

func (c Category) GetName() string {
	return c.Name
}

func (c Category) GetDescription() string {
	return c.Description
}

func (c Category) GetSlug() string {
	return c.Slug
}