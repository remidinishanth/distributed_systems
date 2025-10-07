/************************************************************
tcpechoserver.c

This is a concurrent echo server (it echos all input back to 
the client), using tcp protocols.

Copyright (C) 1995 by Fred Sullivan      All Rights Reserved
Changes by Richard Rabbat, Feb. 1999

source: https://web.mit.edu/2.993/www/lectures/tcpechoserver.c
************************************************************/

#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>

#define PORT 5501
#define QUEUELENGTH 5

void fatalerror(char* executable, char* message) {
  fprintf(stderr, "%s: %s\n", executable, message);
  perror(executable);
  fflush(stdout);
  exit(1);
}

void print_address (int tempaddr)
{
  int addr1, addr2, addr3, addr4;

  /* calculate address in terms of num1.num2.num3.num4 */
  addr1 = tempaddr/256/256/256;
  tempaddr -= addr1*256*256*256;
  addr2 = tempaddr/256/256;
  tempaddr -= addr2*256*256;
  addr3 = tempaddr/256;
  addr4 = tempaddr - addr3*256;
  /* print it out */
  fprintf(stderr, "Accepted connect request from %d.%d.%d.%d\n", addr1, addr2, addr3, addr4);
}
main(int argc, char *argv[]) {
  int sockfd, newsockfd;
  struct sockaddr_in clientaddr, serveraddr;
  int clientaddrlength;
  int pid;
  char* executable, c;

  /* Remember the program name for error messages. */
  executable = argv[0];

  /* Open a TCP socket. */
  sockfd = socket(AF_INET, SOCK_STREAM, 0);
  if (sockfd < 0)
    fatalerror(executable, "can't open socket");

  /* Bind the address to the socket. */
  bzero(&serveraddr, sizeof(serveraddr));
  serveraddr.sin_family = AF_INET;
  serveraddr.sin_addr.s_addr = htonl(INADDR_ANY);
  serveraddr.sin_port = htons(PORT);
  if (bind(sockfd, (struct sockaddr *) &serveraddr,
           sizeof(serveraddr)) != 0)
    fatalerror(executable, "can't bind to socket");
  if (listen(sockfd, QUEUELENGTH) < 0)
    fatalerror(executable, "can't listen");

  while (1) {
    /* Wait for a connection. */
    clientaddrlength = sizeof(clientaddr);
    newsockfd = accept(sockfd, 
                       (struct sockaddr *) &clientaddr,
                       &clientaddrlength);

    print_address(clientaddr.sin_addr.s_addr);

    if (newsockfd < 0)
      fatalerror(executable, "accept failure");

    /* Fork a child to handle the connection. */
    pid = fork();
    if (pid < 0)
      fatalerror(executable, "fork error");
    else if (pid == 0) {
      /* I'm the child. */
      close(sockfd);

      while (read(newsockfd, &c, 1) == 1)
      /* Echo the character. */
      if (write(newsockfd, &c, 1) != 1)
        fatalerror(executable, "can't write to socket");

      exit(EXIT_SUCCESS);
    }
    else
      /* I'm the parent. */
      close(newsockfd);
  }
}
