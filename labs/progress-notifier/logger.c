#include <stdio.h>
#include "logger.h"
#include <stdarg.h>
#include <string.h>
#include <syslog.h>

char logTp = 0; // 0: default/stdout 1:syslogger 

void textcolor(int attr, int fg, int bg)
{	char command[13];
	sprintf(command, "%c[%d;%d;%dm", 0x1B, attr, fg + 30, bg + 40);
	printf("%s", command);
}

int initLogger(char *logType){
    if(strcmp(logType,"stdout")){
        logTp = 1;
        return logTp;
    }
    if(strcmp(logType,"syslog")){
        logTp = 0;
        return logTp;
    }
    return -1;
}
int infof(const char *format, ...){
    va_list arg;
    va_start(arg, format);
    textcolor(BRIGHT, BLUE, BLACK);
    if(logTp){
        vsyslog(1, format, arg);
        return 1;
    }
    int output =  vfprintf (stdout, format, arg);
    textcolor(RESET, WHITE, BLACK);	
    va_end(arg);
    return output;
}
int warnf(const char *format, ...){
    va_list arg;
    va_start(arg, format);
    textcolor(BRIGHT, YELLOW, BLACK);
    if(logTp){
        vsyslog(1, format, arg);
        return 1;
    }
    int output = vfprintf (stdout, format, arg);
    textcolor(RESET, WHITE, BLACK);	
    va_end(arg);
    return output;
}
int errorf(const char *format, ...){
    va_list arg;
    va_start(arg, format);
    textcolor(BRIGHT, RED, BLACK);
    if(logTp){
        vsyslog(1, format, arg);
        return 1;
    }
    int output = vfprintf (stdout, format, arg);
    textcolor(RESET, WHITE, BLACK);	
    va_end(arg);
    return output;
}
int panicf(const char *format, ...){
    va_list arg;
    va_start(arg, format);
    textcolor(BRIGHT, RED, BLACK);
    if(logTp){
        vsyslog(1, format, arg);
        return 1;
    }
	int output = vfprintf (stdout, format, arg);
    textcolor(RESET, WHITE, BLACK);	
    va_end(arg);
    return output;
}