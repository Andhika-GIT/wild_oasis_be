package main

import (
	"context"
	"log"

	"github.com/Andhika-GIT/wild_oasis_be/internal/infrastructure/config"
)

func main() {
	app := config.Bootstrap()

	log.Println("Seeding data")

	err := app.CabinService.SeedCabins(context.Background())

	if err != nil {
		log.Fatalf("error while seeding : %v", err)
	}

	log.Println("Seeding completed")
}
