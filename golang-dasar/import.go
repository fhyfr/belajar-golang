package main

import (
	"belajar-golang/golang-dasar/helper"
)

func main () {
	helper.SayHello("Firman")
	helper.fullName("Helsa", "Nanda") 

	var apps = helper.Application
	var version = helper.version
}