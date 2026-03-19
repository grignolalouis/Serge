package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"serge/backend/pkg/database"
)

func main() {
	// Load .env file (ignore if not found)
	_ = godotenv.Load()

	// Get DSN from environment or use default
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=serge port=5432 sslmode=disable"
	}

	// Connect to database
	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
	fmt.Println("🌱 Starting database seeding...")

	// Uncomment to clear all data before seeding
	// if err := ClearDatabase(db); err != nil {
	//     log.Fatalf("❌ Failed to clear database: %v", err)
	// }

	// Seed ground truth
	if err := SeedGroundTruth(db); err != nil {
		log.Fatalf("❌ Failed to seed ground truth: %v", err)
	}
	fmt.Println("✅ Ground truth injected successfully")

	// Seed fake data
	if err := SeedFakeData(db); err != nil {
		log.Fatalf("❌ Failed to seed fake data: %v", err)
	}
	fmt.Println("✅ Fake data injected successfully")

	fmt.Println("🎉 Database seeding completed!")
}
