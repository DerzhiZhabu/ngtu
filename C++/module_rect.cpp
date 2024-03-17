#include <math.h>
#include <iostream>

float square(float x, float y){
	return x * y;}

float perimeter(float x, float y){
	return (x + y) * 2;}

float diagon(float x, float y){
	return sqrt(x * x + y * y);}

using namespace std;

int main(){
	float x = 3;
	float y = 4;
	cout << square(x, y) << " " << perimeter(x, y) << " " << diagon(x, y)<< endl;
}

