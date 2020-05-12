#include <stdio.h>
#include "logger.h"
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/inotify.h>
#include <sys/types.h>

#define BUF_LEN sizeof(struct inotify_event) * 1024

int fd;
int wd;
int rd;
char* p;
struct inotify_event *event;


void displayEvent(struct inotify_event* event);

int main(char argc, char** argv)
{
    
    // Place your magic here
    if(argc < 2)
    {
        errorf("Invalid input, please type in a directory name\n");
        return -1;
    }

    fd = inotify_init1(O_NONBLOCK);
    wd = inotify_add_watch(fd, argv[1], IN_ALL_EVENTS);
    char* buffer = (char*)malloc(BUF_LEN);

    while(1)
    {
        rd = read(fd, buffer, BUF_LEN);
        p = buffer;
        event = (struct inotify_event*)p;
        for(p = buffer; p < buffer + rd; )
        {
            event = (struct inotify_event *) p;
            displayEvent(event);
            p += sizeof(struct inotify_event) + event -> len;
        }
    }
    close(fd);
    return 0;
}

void displayEvent(struct inotify_event* event)
{
    if(event->mask & IN_ACCESS)     infof("Directory successfully access\n");
    if(event->mask & IN_CREATE)     infof("Directory successfully created\n");
    if(event->mask & IN_DELETE)     errorf("Directory successfully deleted\n");
    if(event->mask & IN_OPEN)       infof("Open directory\n");
    if(event->mask & IN_MODIFY)     warnf("Directory successfully modificated\n");

}
