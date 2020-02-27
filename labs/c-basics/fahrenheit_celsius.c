#include <stdio.h>

#define   LOWER  0       /* lower limit of table */
#define   UPPER  300     /* upper limit */
#define   STEP   20      /* step size */

/* print Fahrenheit-Celsius table */

int main(int argc, char **argv)
{
    int fahr;

    if(argc < 2){
        for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP)
	    printf("Fahrenheit: %3d, Celcius: %6.1f\n", fahr, (5.0/9.0)*(fahr-32));
    }

    if(argc == 2){
        fahr = atoi(argv[1]);
        printf("Fahrenheit: %3d, Celcius: %6.1f\n", fahr, (5.0/9.0)*(fahr-32));
    }
    else{
        int start, end, increment, f;
        
        
         start = atoi(argv[1]);
         end = atoi(argv[2]);
         increment = atoi(argv[3]);

         for(f=start; f<= end; f= f+increment){
             printf("Fahrenheit: %3d, Celcius: %6.1f\n", f, (5.0/9.0)*(f-32));
         }

    }

    return 0;
}
