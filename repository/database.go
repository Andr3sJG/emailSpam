package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"spam.com/enty"
)

type VideoRepository interface {
	Save(video enty.Video)
	FindAll() []enty.Video

}

type database struct {
	connection *gorm.DB
}

func freak(err error) {
	if err != nil {
		panic(err)
	}
}

func NewVideoRepository() VideoRepository {

	//dsn := "root:red@tcp(127.0.0.1:3306)/Test"
	dsn := "n9e2tq4wo6uhxslh:jmm9tpr4pjtkyazi@tcp(ao9moanwus0rjiex.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/apc4exudm9mlh3ul"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	freak(err)

	db.AutoMigrate(&enty.Video{})

	return &database{
		connection: db,
	}

}


func (db *database) FindAll() []enty.Video{
	var videos []enty.Video
	db.connection.Find(&videos)
	return videos
}

func (db *database) Save(video enty.Video){
	db.connection.Create(&video)
}
