package main

import "fmt"

func main(){

	for i := 0; i <= 5; i++ {
		fmt.Printf("Nilai i = %d\n", i) 
		if i == 5 {
			for j := 0; j <= 10; j++ {
				if j == 5 {
					cyrillic := "САШАРВО"
					for index, value := range cyrillic {
						fmt.Printf("character %U, %c, starts at byte position %d\n", value, value, index)
					}
				} else {
					fmt.Printf("Nilai j = %d\n", j)
				}
			}
		} 
	}
	
}