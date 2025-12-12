package main

import "fmt"

func main() {

	var username, password string
	var inputusername, inputpassword string 
	var panjangKarakter int
	var adaAngka, valid bool
	var char byte
	var menu int

	for {
		 
		fmt.Println("1. Register")
		fmt.Println("2. login")
		fmt.Println("3. keluar")
		fmt.Print("pilih: ")
		fmt.Scan(&menu)

		if menu == 3 {
			fmt.Printf("\nanda iseng doang?")
			return
		}

		if menu == 1 {
			percobaanreg := 0 

			for percobaanreg
		} 
	}

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
		return
	}

	// login 
	fmt.Println("===LOGIN===")
	fmt.Println("")

	fmt.Print("username: ")
	fmt.Scan(&inputusername)

	fmt.Print("password: ")
	fmt.Scan(&inputpassword)

	if inputusername == username && inputpassword == password {
		fmt.Println("akses di berikan")
		fmt.Println("selamat datang", username)
		fmt.Println("")
		fmt.Println("==============================")
	} else {
		fmt.Println("akses di tolak!!")
		fmt.Println("cek kembai username atau password ")
		fmt.Println("")
	}
}
