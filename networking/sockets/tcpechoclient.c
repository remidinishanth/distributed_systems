/************************************************************
tcpechoclient.c

This is a client for the tcp echo server.  It sends anything
read from standard input to the server, reads the responses,
and sends them to standard output.

Copyright (C) 1995 by Fred Sullivan      All Rights Reserved
Changes 1999, Richard Rabbat

Source: https://web.mit.edu/2.993/www/lectures/tcpechoclient.c
************************************************************/

#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>

#define PORT 5501

void fatalerror(char* executable, char* message) {
  fprintf(stderr, "%s: %s\n", executable, message);
  perror(executable);
  fflush(stdout);
  exit(1);
}

main(int argc, char *argv[]) {
  int sockfd, c;
  struct sockaddr_in serveraddr;
  char *executable, *server, outbuffer, inbuffer;
  struct hostent *hp;
  int serveraddrlength;

  /* Remember the program name for error messages. */
  executable = argv[0];

  /* find server */
  if (argc == 2)
    server = argv[1];
  else
    fatalerror (executable, "need IP address of server");

  /* look up specified server hostname */
  bzero(&serveraddr, sizeof(serveraddr));
  serveraddr.sin_addr.s_addr = inet_addr(server);
  if ((hp = gethostbyaddr((void*)&(serveraddr.sin_addr), 4, AF_INET)) == NULL)
  {
    fprintf (stderr, "%s: %s: no such host?/n", executable, server);
    exit (1);
  }

  /* put host's address and address type in socket structure */
  bcopy (hp->h_addr, & serveraddr.sin_addr, hp->h_length);
  serveraddr.sin_family = hp->h_addrtype;
  serveraddr.sin_port = htons(PORT);

  /* Open a socket. */
  if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0)
    fatalerror(executable, "can't open socket");

  /* Connect to the server. */
  if (connect(sockfd, (struct sockaddr *) &serveraddr,
              sizeof(serveraddr)) < 0)
    fatalerror(executable, "can't connect to server");

  fprintf (stderr, "type stuff to get it echoed back\n");

  /* Copy input to the server. */
  while ((c = getchar()) != EOF) {
    /* Write a character to the socket. */
    outbuffer = c;
    if (write(sockfd, &outbuffer, 1) != 1)
      fatalerror(executable, "can't write to socket");
    /* Read the response and print it. */
    if (read(sockfd, &inbuffer, 1) != 1)
      fatalerror(executable, "can't read from socket");
    putchar(inbuffer);
  }

  close(sockfd);

  exit(EXIT_SUCCESS);
}
