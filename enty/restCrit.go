package enty

type Restaurant struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Location string `json:"location" gorm:"type:varchar(300);UNIQUE"`
	Cuisine  string `json:"cuisine" gorm:"type:varchar(300)"`
	Critcal  string `json:"critical" gorm:"type:varchar(1000)"`
}
