package main

import (
	"dts-go-07/photo"
	"dts-go-07/user"
	"fmt"
)

func main() {

	//CHALLENGE 01-Dhewangga Arie-Kelas 07//

	//menampilkan nilai i : 21
	var i int = 21
	fmt.Printf("%d \n", i)

	//menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	//menampilkan tanda %
	fmt.Println("%")

	//menampilkan nilai boolean j : true
	var j bool = true
	fmt.Printf("%t \n", j)

	//menampilkan unicode russia : Я (ya)
	cyrillic := 'Я'
	fmt.Printf("%c \n", cyrillic)

	//menampilkan nilai base 10 : 21
	fmt.Printf("%d \n", i)

	//menampilkan nilai base 8 :25
	fmt.Printf("%o \n", i)

	//menampilkan nilai base 16 : f
	output := 15
	fmt.Printf("%x \n", output)

	//menampilkan nilai base 16 : F 13
	fmt.Printf("%X 13 \n", output)

	//menampilkan unicode karakter Я : U+042F
	fmt.Printf("%U \n", cyrillic)

	//menampilkan float : 123.456000
	var k float64 = 123.456
	fmt.Printf("%f \n", k)

	//menampilkan float scientific : 1.234560e+02
	fmt.Printf("%e \n", k)

	// MINI QUIZ
	// budi adalah user bertipe admin
	// dia memiliki foto sebanyak 10 foto
	// buatlah program yang akan menghasilkan output
	// 1. jika user adalah admin dan foto > 10 => akan print "ok"
	// 2. jika user adalah admin dan foto >= 10 => akan print "tidak ok"
	// 3. selain diatas => akan print "oke gak ya?"
	// kriteria:
	// type user disimpan di package user
	// batas foto disimpan di package photo
	// program akan dijalankan di main.go

	pengguna := 0
	if pengguna == user.ADMIN {
		fmt.Println("pengguna adalah ADMIN")
	} else if pengguna == user.NORMAL {
		fmt.Println("pengguna adalah NORMAL")
	} else if pengguna == user.VIP {
		fmt.Println("pengguna adalah VIP")
	}

	budiType := 0 // admin 0
	budiPhoto := 10

	// give white space
	fmt.Println()

	// budiType == user.ADMIN => true
	// budiPhoto > photo.BATAS_FOTO => false
	// budiPhoto >= photo.BATAS_FOTO => true

	if budiType == user.ADMIN && budiPhoto > photo.BATAS_FOTO {
		fmt.Println("ok")
	} else if budiType == user.ADMIN && budiPhoto >= photo.BATAS_FOTO {
		fmt.Println("tidak ok")
	} else {
		fmt.Println("oke gak ya?")
	}
}
