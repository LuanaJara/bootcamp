package main

import 
	"fmt"

	func main() {

		var a [10]int
     fmt.Println("comienzo:", a)

	 a[9] = 100
     fmt.Println("dsd:", a)
     fmt.Println("hst:", a[9])
 
	 fmt.Println("cnt:", len(a))

	 b := [10]int{1, 2, 3, 4, 5,6,7,8,9,10}
     fmt.Println("imp:", b)
	 
	 var twoD [5][8]int
     for i := 0; i < 5; i++ {
        for j := 0; j < 8; j++ {
            twoD[i][j] = i + j
        }
     }
     fmt.Println("intrv: ", twoD)
}
	