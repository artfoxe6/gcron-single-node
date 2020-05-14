package main

import "gcron"

func main() {
	gcron.LoadConfig()
	scheduler := gcron.NewScheduler()
	scheduler.Start()
}