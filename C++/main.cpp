#include <iostream>
#include <locale.h>

#include "module_trian.h"
#include "module_rect.h"

using namespace std;

int main() {
    setlocale(LC_ALL, "Russian");

    int choice;

    cout << "Rectangle - 1" << endl;
    cout << "Triangle - 2" << endl;

    cin >> choice;

    if (choice == 1) {
        float a, b;
        float s, p, d;

        cin >> a >> b;

        s = square(a, b);
        p = perimeter(a, b);
        d = diagon(a, b);

        cout << "Rectangle perimeter: " << p << endl;
        cout << "Rectangle area: " << s << endl;
        cout << "Rectangle diagonal length: " << d;

    }
    else if (choice == 2) {
        float a, b, c;
        float perimeter_tri, area_tri;
        bool equi;

        cin >> a >> b >> c;

        if (verification(a, b, c)) {
            perimeter_tri = perimeter(a, b, c);
            area_tri = area(a, b, c, perimeter_tri);
            equi = equilateral(a, b, c);

            cout << "Triangle perimeter: " << perimeter_tri << endl;
            cout << "Triangle area: " << area_tri << endl;
            cout << "Is it an equilateral triangle? " << (equi ? "Yes" : "No") << endl;

        }
        else {
            cout << "The triangle does not exist!" << endl;
        }
    }

    else {
        cout << "eror" << endl;
    }

    return 0;
}