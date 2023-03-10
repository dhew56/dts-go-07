package main

import (
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
}
