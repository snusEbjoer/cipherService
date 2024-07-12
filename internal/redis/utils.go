package redis

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func addr() string {
	log.Println(fmt.Sprintf("%s:%s", os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PORT")))
	return fmt.Sprintf("%s:%s", os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PORT"))
}

func database() int {
	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		log.Fatalf("Failed to get redis database number hint: check env")
	}
	return db
}
