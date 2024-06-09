#include <stdio.h>
#include "libsum.h"
#include <time.h>

int main(void)
{
    int i,ret;
    int testCount = 10000000;

    clock_t start = clock();
    for(i=0;i<testCount;i++){
        ret = someCFunc(3);
    }

    clock_t end = clock();
    double time_elapsed = (end - start) * 1000000000.0 / CLOCKS_PER_SEC;
    printf("C调用Go函数开销: %f", time_elapsed/testCount);
    return 0;
}
