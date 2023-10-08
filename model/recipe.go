package model

import (
	"errors"
	"time"
)

type Recipe struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Ingredients  string `json:"ingredients"`
	Steps        string `json:"steps"`
	CreationDate string `json:"creationDate"`
	CategoryID   int    `json:"categoryID"`
	UserID       int    `json:"userID"`
}

var (
	ErrIngredientsInvalid = errors.New("ingredients are invalid")
	ErrStepsInvalid       = errors.New("steps are invalid")
	ErrUserIDInvalid      = errors.New("userID is invalid")
	ErrCategoryIDInvalid  = errors.New("categoryID is invalid")
)

type AddRecipe struct {
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Steps       string `json:"steps"`
	CategoryID  int    `json:"categoryID"`
	UserID      int    `json:"userID"`
}

func (a AddRecipe) Validation() error {
	UserExists := false
	CategoryExists := false
	for _, user := range users {
		if user.ID == a.UserID {
			UserExists = true
		}
	}
	for _, category := range categories {
		if category.ID == a.CategoryID {
			CategoryExists = true
		}
	}
	switch {
	case a.Name == "":
		return ErrNameInvalid
	case a.Ingredients == "":
		return ErrIngredientsInvalid
	case a.Steps == "":
		return ErrStepsInvalid
	case !UserExists:
		return ErrUserIDInvalid
	case !CategoryExists:
		return ErrCategoryIDInvalid
	}
	return nil
}

func RecipeOne(id int) (Recipe, error) {
	for _, recipe := range recipes {
		if recipe.ID == id {
			return recipe, nil
		}
	}
	return Recipe{}, ErrNotFound
}

type UpdateRecipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Steps       string `json:"steps"`
	CategoryID  int    `json:"categoryID"`
}

func (u UpdateRecipe) Validation() error {
	CategoryExists := false
	for _, category := range categories {
		if category.ID == u.CategoryID {
			CategoryExists = true
		}
	}
	switch {
	case u.ID == 0:
		return ErrIDInvalid
	case u.Name == "":
		return ErrNameInvalid
	case u.Ingredients == "":
		return ErrIngredientsInvalid
	case u.Steps == "":
		return ErrStepsInvalid
	case !CategoryExists:
		return ErrCategoryIDInvalid
	}
	return nil
}

func RecipesAll(q string) ([]Recipe, error) {
	if q != "" {
		var recipesFiltered []Recipe
		for _, recipe := range recipes {
			if recipe.Name == q {
				recipesFiltered = append(recipesFiltered, recipe)
			}
		}
		return recipesFiltered, nil
	}
	return recipes, nil
}

func (r Recipe) InsertRecipe() (int, error) {
	recipeMaxID++
	r.ID = recipeMaxID
	r.CreationDate = time.Now().Format("2006-01-02")
	recipes = append(recipes, r)
	return r.ID, nil
}

func DeleteRecipe(id int) error {
	for i, recipe := range recipes {
		if recipe.ID == id {
			recipes = append(recipes[:i], recipes[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (r Recipe) UpdateRecipe() error {
	for i, recipe := range recipes {
		if recipe.ID == r.ID {
			recipes[i].Name = r.Name
			recipes[i].Ingredients = r.Ingredients
			recipes[i].Steps = r.Steps
			recipes[i].CategoryID = r.CategoryID
			return nil
		}
	}
	return ErrNotFound
}

var recipeMaxID = 4

var recipes = []Recipe{
	{ID: 1, Name: "Carbonara", Ingredients: "pasta, bacon, eggs, parmesan", Steps: "1. Cook pasta, 2. Cook bacon, 3. Mix eggs and parmesan, 4. Mix all together", CreationDate: "2019-01-01", CategoryID: 1, UserID: 1},
	{ID: 2, Name: "Bolognese", Ingredients: "pasta, minced meat, tomato sauce, onion, garlic", Steps: "1. Cook pasta, 2. Cook minced meat, 3. Cook onion and garlic, 4. Mix all together", CreationDate: "2019-01-02", CategoryID: 1, UserID: 2},
	{ID: 3, Name: "Tiramisu", Ingredients: "eggs, sugar, mascarpone, coffee, ladyfingers", Steps: "1. Mix eggs and sugar, 2. Mix mascarpone and coffee, 3. Mix all together", CreationDate: "2019-01-03", CategoryID: 2, UserID: 3},
	{ID: 4, Name: "Panna cotta", Ingredients: "cream, sugar, gelatin, vanilla", Steps: "1. Mix cream and sugar, 2. Add gelatin and vanilla, 3. Mix all together", CreationDate: "2019-01-04", CategoryID: 2, UserID: 1},
}
