#include <stdio.h>

int main(int argc, char**argv) 
{
    int b = (67*100) + 100000;
    int c = b + 17000;
    int h = 0;

    for (; b <= c; b += 17) {
        if ((b & 1) == 0) {
            h++;
            continue;
        }

        for (int d = 3; d < (b>>1); d++) {
            if ((b%d) == 0) {
                h++;
                break;
            }
        }
    }

    printf("Factored: %d\n", h);
}
