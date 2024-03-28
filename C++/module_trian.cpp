#include <math.h>

bool verification(float a, float b, float c) {
    if(a <= 0 || b <= 0 || c <= 0 || a + b <= c || a + c <= b || b + c <= a){
        return false;
    }
    return true;
}

float perimeter(float a, float b, float c) {
    float p = a + b + c;

    return p;
}

float area(float a, float b, float c, float p) {
    p = p / 2;

    float s = sqrt(p*(p-a)*(p-b)*(p-c));

    return s;
}

bool equilateral(float a, float b, float c) {
    if(a == b || b == c || c == a){
        return true;
    }

    return false;
}
