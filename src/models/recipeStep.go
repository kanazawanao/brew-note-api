package models

type RecipeStep struct {
	ID          int    `json:"id" gorm:"AUTO_INCREMENT"`
	RecipeId    int    `json:"recipeId"`
	StepNumber  int    `json:"stepNumber"`
	Seconds     int    `json:"seconds"`
	AmountWater int    `json:"amountWater"`
	Memo        string `json:"memo"`
}
