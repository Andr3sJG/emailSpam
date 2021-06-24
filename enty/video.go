package enty

type Video struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `json:"title" gorm:"type:varchar(100)"`
	URL   string `json:"url" gorm:"type:varchar(300);UNIQUE"`
	Email string `json:"email" gorm:"type:varchar(300)"`
}
