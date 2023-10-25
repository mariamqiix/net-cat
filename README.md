# \# **net-cat**

## **how to run:**

> ```
> run the project:
> go run .
> or
> go run . [port]
> exp: go run . 8989
>
> run the treminal for the user:
> - nc [IP] [port]
> exp: nc localhost 8989
> ```

### **about this project:**

in project this project you'll be able to send msgs in the terminal to others.  
it's a simple and versatile networking utility implemented in the Go programming language. It allows you to read from and write to network connections, making it a powerful tool for debugging, testing, and other networking tasks. This README file provides an overview of how to use Net-Cat Golang and its features.

## **about the code:**

- only 10 users is allowed.
- use **Ctrl+c** to exit the chat.
- use **"--ChangeName"** to changeyour name.

### **Allowed Packages**

- io
- log
- os
- fmt
- net
- sync
- time
- bufio
- errors
- strings
- reflect

### **Usage (general-tut)**

```
$ go run .
Listening on the port :8989
$ go run . 2525
Listening on the port :2525
$ go run . 2525 localhost
[USAGE]: ./TCPChat $port
$
```

- answer the client with a linux logo and ask for their name, when connection is received.

```
$ nc $IP $port
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:
```

- Accept connection with non-empty name

The client :

```
$ nc $IP $port
```

Server:

```
$ go run . 2525
Listening on the port :2525
```

Client1 (Yenlik):

```
$ nc localhost 2525
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Yenlik
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?
[2020-01-20 16:04:10][Yenlik]:
Lee has joined our chat...
[2020-01-20 16:04:15][Yenlik]:
[2020-01-20 16:04:32][Lee]:Hi everyone!
[2020-01-20 16:04:32][Yenlik]:
[2020-01-20 16:04:35][Lee]:How are you?
[2020-01-20 16:04:35][Yenlik]:great, and you?
[2020-01-20 16:04:41][Yenlik]:
[2020-01-20 16:04:44][Lee]:good!
[2020-01-20 16:04:44][Yenlik]:
[2020-01-20 16:04:50][Lee]:alright, see ya!
[2020-01-20 16:04:50][Yenlik]:bye-bye!
[2020-01-20 16:04:57][Yenlik]:
Lee has left our chat...
[2020-01-20 16:04:59][Yenlik]:
```

Client2 (Lee):

```
$ nc localhost 2525
Yenliks-MacBook-Air:simpleTCPChat ybokina$ nc localhost 2525
Yenliks-MacBook-Air:simpleTCPChat ybokina$ nc localhost 2525
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Lee
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?
[2020-01-20 16:04:15][Lee]:Hi everyone!
[2020-01-20 16:04:32][Lee]:How are you?
[2020-01-20 16:04:35][Lee]:
[2020-01-20 16:04:41][Yenlik]:great, and you?
[2020-01-20 16:04:41][Lee]:good!
[2020-01-20 16:04:44][Lee]:alright, see ya!
[2020-01-20 16:04:50][Lee]:
[2020-01-20 16:04:57][Yenlik]:bye-bye!
[2020-01-20 16:04:57][Lee]:^C
```

## **authers:**

- rhelal
- maabbas
- fatabbasutc8
