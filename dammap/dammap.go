package dammap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/zyun-i/dam-data/model"
)

func Run() {
	// json
	jsonFromFile, err := ioutil.ReadFile("./data/dam-map-sorted.json")
	if err != nil {
		log.Fatal(err)
	}

	var jsonData []*model.Dammap
	err = json.Unmarshal(jsonFromFile, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", jsonData[0])

	dsn := "host=localhost user=postgres password=pass dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.CreateInBatches(jsonData, 100)

	if err != nil {
		log.Fatal("Could not connect database")
	}
}
