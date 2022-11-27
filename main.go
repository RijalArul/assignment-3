package main

import (
	"fmt"
	"go-cron/config"
	"go-cron/models"

	"math/rand"

	"time"

	"github.com/jasonlvhit/gocron"
)

const (
	Aman   = "Aman"
	Siaga  = "Siaga"
	Bahaya = "Bahaya"
)

func task() {
	rand.Seed(time.Now().UnixNano())
	db := config.GetDB()
	element := models.Element{}
	min, max := 1, 100
	valueWater, valueWind := rand.Intn(max-min+1)+min, rand.Intn(max-min+1)+min
	fmt.Println(valueWater, valueWind)

	if db.Model(&element).Where("id = ?", 1).Updates(&element).RowsAffected == 0 {
		newElement := db.Create(&element)
		fmt.Println("New Elements", newElement)
	}

	switch valueWater > 0 && valueWind > 0 {
	case valueWater < 5:
		element = models.Element{
			Water:       valueWater,
			StatusWater: Aman,
		}
		db.Model(&element).Where("id = ?", 1).Updates(&element)
		fmt.Println(element)
	case valueWater >= 6 && valueWater <= 8:
		element = models.Element{
			Water:       valueWater,
			StatusWater: Siaga,
		}
		db.Model(&element).Where("id = ?", 1).Updates(&element)
		fmt.Println(element)
	case valueWater > 8:
		element = models.Element{
			Water:       valueWater,
			StatusWater: Bahaya,
		}
		db.Model(&element).Where("id = ?", 1).Updates(&element)
	case valueWind < 6:
		element = models.Element{
			Wind:       valueWind,
			StatusWind: Aman,
		}
		db.Model(&element).Where("id = ?", 1).Updates(&element)
		fmt.Println(element)
	case valueWind >= 7 && valueWind <= 15:
		element = models.Element{
			Wind:       valueWind,
			StatusWind: Siaga,
		}
		db.Model(&element).Where("id = ?", 1).Updates(&element)
		fmt.Println(element)
	case valueWind > 15:
		element = models.Element{
			Wind:       valueWind,
			StatusWind: Bahaya,
		}
		db.Model(&element).Where("id = ?", 1).Updates(&element)
		fmt.Println(element)
	default:
		defaultElement := db.Where("id = ?", 1).First(&element)
		fmt.Println(defaultElement)
	}

}

func main() {
	config.StartDB()
	s := gocron.NewScheduler()
	s.Every(15).Seconds().Do(task)
	<-s.Start()
}
