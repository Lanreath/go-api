package model

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AddCategory struct {
	Name string `json:"name"`
}

func (a AddCategory) Validation() error {
	if a.Name == "" {
		return ErrNameInvalid
	}
	return nil
}

func CategoryOne(id int) (Category, error) {
	for _, category := range categories {
		if category.ID == id {
			return category, nil
		}
	}
	return Category{}, ErrNotFound
}

type UpdateCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u UpdateCategory) Validation() error {
	switch {
	case u.ID == 0:
		return ErrIDInvalid
	case u.Name == "":
		return ErrNameInvalid
	}
	return nil
}

func CategoriesAll(q string) ([]Category, error) {
	if q != "" {
		var categoriesFiltered []Category
		for _, category := range categories {
			if category.Name == q {
				categoriesFiltered = append(categoriesFiltered, category)
			}
		}
		return categoriesFiltered, nil
	}
	return categories, nil
}

func (c Category) InsertCategory() (int, error) {
	categoryMaxID++
	c.ID = categoryMaxID
	categories = append(categories, c)
	return c.ID, nil
}

func (c Category) UpdateCategory() error {
	for i, category := range categories {
		if category.ID == c.ID {
			categories[i].Name = c.Name
			return nil
		}
	}
	return ErrNotFound
}

func DeleteCategory(id int) error {
	for i, category := range categories {
		if category.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

var categoryMaxID = 3

var categories = []Category{
	{ID: 1, Name: "Pasta"},
	{ID: 2, Name: "Dessert"},
	{ID: 3, Name: "Meat"},
}
