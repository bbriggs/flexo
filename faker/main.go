package main

import (
	"github.com/bxcodec/faker/v3"

	"flexo/model"
	"flexo/util"
)

func main() {

	db := util.DBconnect(
		util.LookupEnv("FLEXO_DB_USER", "flexo"),
		util.LookupEnv("FLEXO_DB_PASS", "flexo"),
		util.LookupEnv("FLEXO_DB_HOST", "localhost"),
		util.LookupEnv("FLEXO_DB_NAME", "flexo"),
		util.LookupEnv("FLEXO_DB_SSLMODE", false),
	)

	categoryCount := 15
	teamCount := 100

	fakeCategories(db, categoryCount)
	fakeTeams(db, teamCount)
}

func fakeCategories(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		category := model.Category{}
		err := faker.FakeData(&category)
		res := db.Create(&category)
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}

func fakeTeams(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		team := model.Team{}
		err := faker.FakeData(&team)
		res := db.Create(&team)
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}
