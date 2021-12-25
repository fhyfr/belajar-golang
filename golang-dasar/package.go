package main

import (
	"flag"
	"fmt"
)

func main() {
	// package flag
	var host *string = flag.String("host", "localhost", "put your database host")
	var user *string = flag.String("user", "firman", "put your database user")
	var password *string = flag.String("password", "firmanganteng123", "put your database password")
	var pool *int = flag.Int("pool", 2, "limit of database pooling")

	flag.Parse()

	fmt.Println("DB_HOST: ", *host)
	fmt.Println("DB_USER: ", *user)
	fmt.Println("DB_PASSWORD: ", *password)
	fmt.Println("DB_POOL: ", *pool)

	// package flag: usage
	// ex: go run golang-dasar/package.go -host=https://hostname.com -user=root -password=root123456789 -pool=2
}