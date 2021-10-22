package main

import "fmt"

func primeNumber(bilangan int) bool {

		for i:=2; i <= bilangan;i++{
		if bilangan % i == 0{
			return false
		
		}else{
			return true
		}
	}

	return true

}

func main(){
	fmt.Println(primeNumber(1000000007)) // true

	fmt.Println(primeNumber(1500450271)) // true

	fmt.Println(primeNumber(1000000000)) // false

	fmt.Println(primeNumber(10000000019)) // true

	fmt.Println(primeNumber(10000000033)) // true

	fmt.Println(primeNumber(2))

	fmt.Println(primeNumber(7))

	fmt.Println(primeNumber(10))

	fmt.Println(primeNumber(10))

	fmt.Println(primeNumber(10))


}