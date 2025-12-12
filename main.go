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
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&menu)

		if menu == 3 {
			fmt.Println("Keluar program.")
			return
		}

		if menu == 1 {
			percobaanreg := 0

			for percobaanreg < 3 {

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

					adahurufbesar := false //besar
					adahurufkecil := false //kecil
					adadigit := false //digit
					adasimbol := false // simbol

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
					for percobaanlog < 3 {
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
							return 
						} else {
							percobaanlog++
							fmt.Println("akses di tolak!!")
							fmt.Println("Sisa percobaan:", 3-percobaanlog)
						}
					}

					if percobaanlog == 3 {
						fmt.Println("")
						fmt.Println("Login gagal 3 kali.")
						fmt.Println("anda pelupa ya ?")
						return 
					}
					// ===== END TAMBAHAN =====

					break
				} else {
					percobaanreg++
					fmt.Println("Register gagal, ",)
					if percobaanreg == 3 {
						fmt.Println("Register gagal 3 kali. Program dihentikan.")
						return
					}
				}
			}
		}
		if menu == 2 {
			if username == "" {
				fmt.Println("Belum ada akun, register dulu.")
				continue
			}

			percobaanlog := 0
			for percobaanlog < 3 {

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
					return
				} else {
					percobaanlog++
					fmt.Println("akses di tolak!!")
					fmt.Println("Sisa percobaan:", 3-percobaanlog)
				}
			}

			if percobaanlog == 3 {
				fmt.Println("")
				fmt.Println("Login gagal 3 kali.")
				fmt.Println("anda pelupa ya ?")
				return
			}
		}
	}
}
