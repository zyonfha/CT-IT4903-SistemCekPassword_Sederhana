package main

import "fmt"

func main() {

	var username, password string
	var inputusername, inputpassword string
	var panjangKarakter int
	var adaAngka, valid bool
	var char byte

	var menu int

	programawal := true

	for programawal {

		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&menu)

		if menu == 1 {
			percobaanreg := 0
			registersukses := false

			for percobaanreg < 3 && !registersukses {
				valid = true

				fmt.Println("Register")
				fmt.Print("Buat Username: ")
				fmt.Scan(&username)

				fmt.Print("Buat Password: ")
				fmt.Scan(&password)

				panjangKarakter = len(password)
				adaAngka = false

				for i := 0; i < panjangKarakter; i++ {
					char = password[i]
					if char >= '0' && char <= '9' {
						adaAngka = true
					}
				}

				if panjangKarakter < 8 {
					fmt.Println("Password minimal 8 karakter")
					valid = false
				} else if !adaAngka {
					fmt.Println("Pasworod minimal ada 1 angka")
					valid = false
				}

				if valid {
					registersukses = true
					fmt.Println("Register Berhasil")

					adahurufbesar := false
					adahurufkecil := false
					adadigit := false
					adasimbol := false

					for i := 0; i < len(password); i++ {
						c := password[i]
						if c >= 'A' && c <= 'Z' {
							adahurufbesar = true
						} else if c >= 'a' && c <= 'z' {
							adahurufkecil = true
						} else if c >= '0' && c <= '9' {
							adadigit = true
						} else {
							adasimbol = true
						}
					}

					level := "Lemah"
					if len(password) >= 12 && adahurufbesar && adahurufkecil && adadigit && adasimbol {
						level = "Sangat Kuat"
					} else if len(password) >= 8 && adadigit && (adahurufbesar || adasimbol) {
						level = "Kuat"
					} else if len(password) >= 8 && adadigit {
						level = "Sedang"
					}

					fmt.Println("Level Password:", level)

					percobaanlog := 0
					loginsukses := false
					for percobaanlog < 3 && !loginsukses {

						fmt.Println("===LOGIN===")
						fmt.Println("")
						fmt.Print("username: ")
						fmt.Scan(&inputusername)

						fmt.Print("password: ")
						fmt.Scan(&inputpassword)

						if inputusername == username && inputpassword == password {
							loginsukses = true
							fmt.Println("akses di berikan")
							fmt.Println("selamat datang", username)
							fmt.Println("")
							programawal = false

						} else {
							percobaanlog++
							fmt.Println("akses di tolak!!")
							fmt.Println("Sisa percobaan:", 3-percobaanlog)
						}
					}

					if !loginsukses {
						fmt.Println("")
						fmt.Println("Login gagal 3 kali.")
						fmt.Println("anda pelupa ya ?")
						programawal = false
					}

				} else {
					percobaanreg++
					fmt.Println("Register gagal ")
				}
			}

			if percobaanreg == 3 {
				fmt.Println("register gagal 3 kali. Program dihentikan.")
				programawal = false
			}
		}

		if menu == 2 {
			if username == "" {
				fmt.Println("Belum ada akun, register dulu.")
			} else {

				percobaanlog := 0
				loginsukses := false

				for percobaanlog < 3 && !loginsukses {

					fmt.Print("username: ")
					fmt.Scan(&inputusername)

					fmt.Print("password: ")
					fmt.Scan(&inputpassword)

					if inputusername == username && inputpassword == password {
						loginsukses = true
						fmt.Println("akses di berikan")
						fmt.Println("selamat datang", username)
						fmt.Println("")
						programawal = false
					} else {
						percobaanlog++
						fmt.Println("Login gagal, sisa:", 3-percobaanlog)
					}
				}

				if !loginsukses {
					fmt.Println("")
					fmt.Println("Login gagal 3 kali.")
					fmt.Println("anda pelupa ya ?")
					programawal = false
				}
			}

		}

		if menu == 3 {
			fmt.Println("Keluar program.")
			programawal = false

		}
	}
}
