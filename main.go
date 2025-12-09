package main

import "fmt"

func main() {
	var username, password string
	var panjangKarakter int
	var adaAngka, valid bool
	var char byte

	// register dan input username
	fmt.Println("Register")
	fmt.Print("Buat Username: ")
	fmt.Scan(&username)

	//input password
	fmt.Print("Buat Password: ")
	fmt.Scan(&password)

	panjangKarakter = len(password)
	adaAngka = false

	//cek setiap karakter
	for i := 0; i < panjangKarakter; i++ {
		char = password[i]

		if char >= '0' && char <= '9' {
			adaAngka = true
		}
	}
	valid = true

	if panjangKarakter < 8 {
		fmt.Println("Password minimal 8 karakter")
		valid = false
	} else if !adaAngka {
		fmt.Println("Pasworod minimal ada 1 angka")
		valid = false
	}

	if valid {
		fmt.Println("Register Berhasil")
	} else {
		fmt.Println("Register gagal")
	}

}
