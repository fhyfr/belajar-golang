package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	/* package flag */
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

	/* package strconv / string conversion */
	boolean, err := strconv.ParseBool("t")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("245000", 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// convert string to int with decimal base (automatic)
	// alternate of ParseInt
	value, err := strconv.Atoi("10231313")

	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err.Error())
	}
}