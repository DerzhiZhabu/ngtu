package rect
import (
"math"
)

func square(x, y float64) float64{
	return x * y}
	
func perimeter(x, y float64) float64{
	return (x + y) * 2} \\ returning perim
	
func diagon(x, y float64) float64{
	var z float64
	z = math.Sqrt(x * x + y * y)
	return z}
