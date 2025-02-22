package models

type Recipe struct {
	ID                  int    `json:"id" gorm:"AUTO_INCREMENT"`
	UserId              string `json:"userId"`
	Title               string `json:"title"`
	GrindSizeId         int    `json:"grindSizeId"`
	ExtractionEquipment string `json:"extractionEquipment"`
	CoffeeType          string `json:"coffeeType"`
	WaterTemperature    int    `json:"waterTemperature"`
	BeanDose            int    `json:"beanDose"`
}
