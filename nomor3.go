package main

import "fmt"



func ArrayMerge(arrayA, arrayB []string) []string {

	// your code here
	var nampung []string

	//cek
	cekArray := true

	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(arrayB); j++ {
			if arrayA[i] == arrayB[j] {
				cekArray = false
			}
		}
		if cekArray == true {
			nampung = append(nampung, arrayA[i])
		}
	}

	gabung := append(nampung, arrayB[:]...)
	return gabung
}


func main() {

   // Test cases

   fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

   // ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

 

   fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

   // ["sergei", "jin", "steve", "bryan"]

 

   fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

   // ["alisa", "yoshimitsu", "devil jin", "law"]

 

   fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

   // ["devil jin", "sergei"]

 

   fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

   // ["hwoarang"]

 

   fmt.Println(ArrayMerge([]string{}, []string{}))

   // []

}