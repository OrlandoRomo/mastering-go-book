#include <stdio.h>
#include <stdlib.h>
#include "quickSort.h"

int main()
{
    printf("Sorting array\n");
    // Create a random array with 100 elements
    GoInt64 data[1000000];
    for (GoInt64 i= 0; i < 1000000;i++){
        data[i] = rand()%1000;
    }

    // GoInt data[6] = {77, 12, 5, 99, 28, 23};
    // For GoSlices we create an array of type GoInt and then pass it the declaration of GoSlice.
    // The second paramater is the length of the Slice
    GoSlice nums = {data, 1000000};
    SortArray(nums);
    for (GoInt64 i = 0; i < 1000000; i++){
        // Casting every element into GoInt type
        printf("%lld, ", ((GoInt64 *)nums.data)[i]);
    }
    printf("\n");
    return 0;
}