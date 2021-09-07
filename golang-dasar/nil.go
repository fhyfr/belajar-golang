/*
Secara default saat menginisiasi variabel,
maka di Golang akan otomatis dibuatkan default value,
seperti int -> 0, bool -> false, dsb
**
Nil atau null (data kosong) pada Golang bisa didefinisikan dengan
kata kunci "nil"
***
Nil hanya bisa digunakan pada tipe data interface, function, map,
slice, pointer, dan channel
*/

package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = newMap("") // argumen string kosong

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}
