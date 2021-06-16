#include <stdio.h>
#include "usedByC.h"

int main(int argc, char **argv){
    GoInt x = 12;
    GoInt y = 23;

    printf("About to call a go function\n");
    PrintMessage();

    GoInt8 p = Multiply(x,y);
    printf("Product: %d\n", p);
    printf("It worked!\n");
    return 0;
}