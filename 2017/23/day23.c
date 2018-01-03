#include <stdio.h>

int main(int argc, char**argv) 
{
    int a = 1;
    int b = 0;
    int c = 0;
    int d = 0;
    int e = 0;
    int f = 0;
    int g = 0;
    int h = 0;

    b = 67;
    c = b;

    if (a != 0) {
        b = (b*100) + 100000;
        c = b + 17000;
    }

    while (1) {
        printf("a: %d, b: %d, c: %d, d: %d, e: %d, f: %d, g: %d, h: %d\n", a, b, c, d, e, f, g, h);
        fflush(stdout);

        f = 1;

        if ((b & 1) == 0) {
            f = 0;
        } else {
            for (d = 3; d < (b>>1); d++) {
                if ((b%d) == 0) {
                    f = 0;
                }
            }
        }

        if (f == 0) {
            h++;
        }

        if (b == c) {
            break;
        }

        b += 17;
    }

    printf("a: %d, b: %d, c: %d, d: %d, e: %d, f: %d, g: %d, h: %d\n", a, b, c, d, e, f, g, h);
}
