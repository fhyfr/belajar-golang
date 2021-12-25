package main

import (
	"fmt"
	"os"

	// jika package yg mengandung fungsi init tidak dipanggil bisa gunakan blank identifier
	_ "github.com/fhyfr/belajar-golang/golang-dasar/database"
)

func main()  {
	// result := database.GetDatabase()
	// fmt.Println(result)

	// package os : argument
	args := os.Args
	fmt.Println(args)

	// package os : hostname
	name, error := os.Hostname()

	if error == nil {
		fmt.Println("Hostname: ", name)
	} else {
		fmt.Println("error: ", error)
	}

	// package os : get os environment
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println("username: ", username, "password: ", password)
}