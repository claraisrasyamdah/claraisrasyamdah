package main

import "fmt"

func sumDigit(bilangan int){
	var n1, hasil int
	hasil = 0 
	for bilangan != 0 {
		n1 = bilangan % 10 
		hasil = hasil + n1
		bilangan = bilangan / 10 
	}
	fmt.Println(hasil)
}

func main(){
	var x int
	fmt.Scan(&x)
	sumDigit(x)
}