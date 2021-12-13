/*
Pada Golang menggunakan metode pass-by-value,
apabila membuat variabel baru yang diisi dengan variabel lain,
maka nilai variabel baru tersebut akan berisi salinan value
dari variabel sumber
**
Pointer adalah kemampuan membuat reference ke lokasi data
di memori yang sama, tanpa menduplikasi data yg sudah ada
***
Pointer dapat membuat pass-by-reference
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Kebumen", "Jawa Tengah", "Indonesia"}

	// mendefinisikan pass-by-reference dengan pointer
	// menggunakan operator &
	address2 := &address1
	address2.City = "Bandung" // value address1 juga akan berubah

	address3 := &address1
	address3.Country = "Korea"

	// Operator *
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// hanya value address2 yg berubah

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // semua variabel yg reference ke address1 akan berubah

	// Membuat pointer dengan function new()
	// function new membuat data kosong
	var address4 *Address = new(Address)
	address4.City = "Bogor"

	var alamat = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "",
	}

	// pointer di function
	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)
}
