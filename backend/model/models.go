package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("example.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}

	err = db.AutoMigrate(&Student{}, &Course{}, &Selection{}, &Teacher{})
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}

	CreateStudentsExample()
	CreateTeachersExample()
	CreateCoursesExample()
	CreateSelectionsExample()
}
