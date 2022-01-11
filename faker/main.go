package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"

	"github.com/SECCDC/flexo/model"
	"github.com/SECCDC/flexo/util"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	db := util.DBconnect(
		util.LookupEnv("FLEXO_DB_USER", "flexo"),
		util.LookupEnv("FLEXO_DB_PASS", "flexo"),
		util.LookupEnv("FLEXO_DB_HOST", "localhost:5432"),
		util.LookupEnv("FLEXO_DB_NAME", "flexo"),
		util.LookupEnv("FLEXO_DB_SSLMODE", "disable"),
	)

	categoryCount := 15
	teamCount := 100
	targetCount := 50

	fakeCategories(db, categoryCount)
	fakeTeams(db, teamCount)
	fakeTargets(db, targetCount)
}

func fakeCategories(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		category := model.Category{}

		err := faker.FakeData(&category)
		if err != nil {
			fmt.Println(err)
		}
		category.Multiplier = randRange(1, 15)

		res := db.Create(&category)
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}

func fakeTargets(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		target := model.Target{}

		err := faker.FakeData(&target)
		if err != nil {
			fmt.Println(err)
		}

		res := db.Create(&target)
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}

func fakeTeams(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		team := model.Team{}

		err := faker.FakeData(&team)
		if err != nil {
			fmt.Println(err)
		}

		res := db.Create(&team)
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
}

/* #nosec */
func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
