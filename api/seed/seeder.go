package seed

import (
	"log"

	"github.com/dallaspadilla/silverline/api/models"
	"github.com/jinzhu/gorm"
)

var pings = []models.Ping{
	models.Ping{
		PacketLoss:   "0",
		MinRoundTrip: "21",
		MaxRoundTrip: "26",
		AvgRoundTrip: "23",
	},
	models.Ping{
		PacketLoss:   "1",
		MinRoundTrip: "22",
		MaxRoundTrip: "27",
		AvgRoundTrip: "24",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Ping{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Ping{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range pings {
		err = db.Debug().Model(&models.Ping{}).Create(&pings[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
