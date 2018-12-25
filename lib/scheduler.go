package lib

import (
	"log"

	"github.com/robfig/cron"
)

func Schedule(schedule string, fn func()) {
	c := cron.New()
	c.AddFunc(schedule, fn)

	log.Println("Scheduled purchasing started...")
	c.Start()
}
