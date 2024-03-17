package main

import (
"fmt"
"laba1/rect"
"laba1/trian"
)

func main(){
	var rect_side1, rect_side2 float64
	var trian_side1, trian_side2, trian_side3 float64
	
	fmt.Println("Choose figure: 1 - rectangle; 2 - triangle")
	
	var choose int
	fmt.Scan(&choose)
	
	if choose == 1{
		fmt.Println("Type 2 sides of rectangle:")
		fmt.Scan(&rect_side1, &rect_side2)
		
		var squ float64 = rect.Square(rect_side1, rect_side2)
		var perim float64 = rect.Perimeter(rect_side1, rect_side2)
		var diag float64 = rect.Diagon(rect_side1, rect_side2)
		
		fmt.Println("Area -", squ, "Perimete -", perim, "Diagonal -", diag)
	} else if choose == 2{
		fmt.Println("Type 3 sides of triangle:")
		fmt.Scan(&trian_side1, &trian_side2, &trian_side3)
		
		var perim float64 = trian.Perimeter(trian_side1, trian_side2, trian_side3)
		var squ float64 = trian.Area(trian_side1, trian_side2, trian_side3, perim)
		var rb bool = trian.Equilateral(trian_side1, trian_side2, trian_side3)
		
		if trian.Verification(trian_side1, trian_side2, trian_side3){
			fmt.Println("Area -", squ, "Perimiter -", perim, "Equilateral -", rb)
		}else{
			fmt.Println("This is not triangle")}
	}else{
		fmt.Println("What is this")}
}
