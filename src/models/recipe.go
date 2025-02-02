package models

type Recipe struct {
	ID                  int    `json:"id" gorm:"AUTO_INCREMENT"`
	UserId              string `json:"userId"`
	ExtractionEquipment string `json:"extractionEquipment"`
	CoffeeType          string `json:"coffeeType"`
	WaterTemperature    int    `json:"waterTemperature"`
	BeanDose            int    `json:"beanDose"`
}
