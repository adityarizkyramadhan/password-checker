# Proyek 1: Membuat Program Validasi Password

Tujuan proyek ini untuk membantu saya khususnya untuk check sebuah password input dari user pada API yang saya buat

## Deskripsi Proyek

Proyek ini bertujuan untuk membuat sebuah program menggunakan bahasa pemrograman Go untuk memeriksa karakter dan panjang dari sebuah password yang diinputkan oleh pengguna. Program akan memberikan umpan balik apakah password yang diinputkan memenuhi kriteria yang ditentukan dengan boolean.

## Berikut adalah langkah-langkah umum yang akan dilakukan oleh program:

- Menerima input dari pengguna berupa password.
- Memeriksa apakah password memenuhi configurasi yang ditentukan.
- Jika password memenuhi semua kriteria yang ditentukan, program akan memberikan pesan bahwa password valid.

## Perbaikan yang dapat dilakukan

Proyek ini dapat diperluas dengan menambahkan fitur-fitur seperti penggunaan regex untuk pemeriksaan pola karakter tertentu, memperbolehkan pengguna untuk menentukan persyaratan password mereka sendiri, atau menyimpan log aktivitas validasi password. Melakukan check pada password apakah password tersebut mudah ditebak atau tidak.


## Cara Menjalankan Program

```go
package main

import (
	"fmt"

	passwordchecker "github.com/adityarizkyramadhan/password-checker"
)

func main() {
	passwordChecker := passwordchecker.NewPasswordChecker(passwordchecker.PasswordCheckerConfig{
		MinLength:      8,
		MaxLength:      20,
		AllowSymbol:    true,
		AllowSpace:     false,
		AllowNumber:    true,
		AllowLowerCase: true,
		AllowUpperCase: true,
		MustUnique:     true,
	})

	fmt.Println(passwordChecker.Check("12345678Adf)"))
}
```
