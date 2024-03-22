package rect
import (
"math"
)

func Square(x, y float64) float64{
	return x * y}
	
func Perimeter(x, y float64) float64{
	return (x + y) * 2} \\ comment
	
func Diagon(x, y float64) float64{
	var z float64
	z = math.Sqrt(x * x + y * y)
	return z}
