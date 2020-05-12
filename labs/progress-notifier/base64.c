#include <stdio.h>
#include "logger.h"
#include <inttypes.h>
#include <string.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>
#include <fcntl.h>
#include <signal.h>


//Encoding algorithm from: https://en.wikibooks.org/wiki/Algorithm_Implementation/Miscellaneous/Base64


#define WHITESPACE 64
#define EQUALS     65
#define INVALID    66

static const unsigned char d[] = {
    66,66,66,66,66,66,66,66,66,66,64,66,66,66,66,66,66,66,66,66,66,66,66,66,66,
    66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,62,66,66,66,63,52,53,
    54,55,56,57,58,59,60,61,66,66,66,65,66,66,66, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
    10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,66,66,66,66,66,66,26,27,28,
    29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,66,66,
    66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,
    66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,
    66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,
    66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,
    66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,66,
    66,66,66,66,66,66
};

off_t size;
size_t progress;

// OPTION #1

int base64decode (char *in, size_t inLen, unsigned char *out, size_t *outLen) { 
    char *end = in + inLen;
    char iter = 0;
    uint32_t buf = 0;
    size_t len = 0;
    
    while (in < end) {
        sleep(1);
        unsigned char c = d[*in++];
        
        switch (c) {
        case WHITESPACE: continue;   /* skip whitespace */
        case INVALID:    return 1;   /* invalid input, return error */
        case EQUALS:                 /* pad character, end of data */
            in = end;
            continue;
        default:
            buf = buf << 6 | c;
            iter++; // increment the number of iteration
            /* If the buffer is full, split it into bytes */
            if (iter == 4) {
                if ((len += 3) > *outLen) return 1; /* buffer overflow */
                *(out++) = (buf >> 16) & 255;
                *(out++) = (buf >> 8) & 255;
                *(out++) = buf & 255;
                buf = 0; iter = 0;

            }   
        }
    }
   
    if (iter == 3) {
        if ((len += 2) > *outLen) return 1; /* buffer overflow */
        *(out++) = (buf >> 10) & 255;
        *(out++) = (buf >> 2) & 255;
    }
    else if (iter == 2) {
        if (++len > *outLen) return 1; /* buffer overflow */
        *(out++) = (buf >> 4) & 255;
    }

    *outLen = len; /* modify to reflect the actual output size */
    return 0;
}

// OPTION #2

int base64encode(const void* data_buf, size_t dataLength, char* result, size_t resultSize)
{
   const char base64chars[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
   const uint8_t *data = (const uint8_t *)data_buf;
   size_t resultIndex = 0;
   size_t x;
   uint32_t n = 0;
   progress = 0;
   int padCount = dataLength % 3;
   uint8_t n0, n1, n2, n3;
   for (x = 0; x < dataLength; x += 3) 
   {
       progress++;
       sleep(1);
      n = ((uint32_t)data[x]) << 16; 
      
      if((x+1) < dataLength)
         n += ((uint32_t)data[x+1]) << 8;
      
      if((x+2) < dataLength)
         n += data[x+2];
      n0 = (uint8_t)(n >> 18) & 63;
      n1 = (uint8_t)(n >> 12) & 63;
      n2 = (uint8_t)(n >> 6) & 63;
      n3 = (uint8_t)n & 63;

      if(resultIndex >= resultSize) return 1;   /* indicate failure: buffer too small */
      result[resultIndex++] = base64chars[n0];
      if(resultIndex >= resultSize) return 1;   /* indicate failure: buffer too small */
      result[resultIndex++] = base64chars[n1];
      if((x+1) < dataLength)
      {
         if(resultIndex >= resultSize) return 1; 
         result[resultIndex++] = base64chars[n2];
      }
      if((x+2) < dataLength)
      {
         if(resultIndex >= resultSize) return 1; 
         result[resultIndex++] = base64chars[n3];
      }
   }
   if (padCount > 0) 
   { 
      for (; padCount < 3; padCount++) 
      { 
         if(resultIndex >= resultSize) return 1;   
         result[resultIndex++] = '=';
      } 
   }
   if(resultIndex >= resultSize) return 1; 
   result[resultIndex] = 0;
   return 0;  
}

static void
sigHandler(int sig)
{
    infof("Processing File: %lf\n", 100.0f*progress/size);
}


int main(char argc, char** argv){
    int fd = openFile(argv[2]);
    if(fd == -1){
        errorf("Invalid file\n");
        return -1;
    }
    size = getSize(fd);

    char* fileBuffer = (char*)malloc(sizeof(char)*size);
    read(fd,fileBuffer,size);
    if (signal(SIGINT, sigHandler) == SIG_ERR)
        infof("signal");
    
    if(!strcmp(argv[1],"--encode")){
        char* encodedBuffer = (char*)malloc(sizeof(char)*60000000);
        base64encode(fileBuffer, size, encodedBuffer, 60000000);
    }
    else if(!strcmp(argv[1],"--decode")){
        char* decodedBuffer = (char*)malloc(sizeof(char)*size);
        size_t decSize = sizeof(char)*size;
        base64decode(fileBuffer, size, decodedBuffer, &decSize);
    }
    else{
        errorf("Invalid\n");
    }

    return 0;
}

off_t getSize(int fd){
    off_t curOffset = lseek(fd, (size_t)0, SEEK_CUR);
    off_t size = lseek(fd, (size_t)0, SEEK_END); 
    lseek(fd, curOffset, SEEK_SET);
    return size;
}

int openFile(char* filename){
    int fd = open(filename, O_RDONLY);
    return fd;
}

