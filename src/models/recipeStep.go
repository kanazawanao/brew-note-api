package models

type RecipeStep struct {
	ID         int    `json:"id" gorm:"AUTO_INCREMENT"`
	RecipeId   int    `json:"recipeId"`
	StepNumber int    `json:"stepNumber"`
	Memo       string `json:"memo"`
	Seconds    int    `json:"seconds"`
}
