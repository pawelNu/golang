# Golang

Learning Golang

- [Golang](#golang)
  - [Setup](#setup)
  - [Tutorials](#tutorials)
    - [Go Programming - Golang Course with Bonus Projects](#go-programming---golang-course-with-bonus-projects)
    - [Golang Tutorial for Beginners | Full Go Course](#golang-tutorial-for-beginners--full-go-course)
    - [The Best Resources to Learn Golang (If I Could Start Over)](#the-best-resources-to-learn-golang-if-i-could-start-over)
    - [Golang course for beginners](#golang-course-for-beginners)
    - [Learn Go Programming by Building 11 Projects – Full Course](#learn-go-programming-by-building-11-projects--full-course)
    - [What is the BEST Way to Study Golang? (and get hired...)](#what-is-the-best-way-to-study-golang-and-get-hired)
  - [REST API Client](#rest-api-client)

## Setup

Install Go: https://go.dev/doc/install -> download for Linux

1. Remove any previous Go installation by deleting the `/usr/local/go` folder (if it exists).

    ```
    sudo rm -rf /usr/local/go
    ```

2. Extract the archive you just downloaded into `/usr/local`, creating a fresh Go tree in `/usr/local/go`:

    ```
    sudo tar -C /usr/local -xzf /home/$(whoami)/Downloads/go1.23.0.linux-amd64.tar.gz
    ```

    Do not untar the archive into an existing `/usr/local/go` tree. This is known to produce broken Go installations.

3. Add /usr/local/go/bin to the PATH environment variable.

    You can do this by adding the following line to your `$HOME/.profile` or `/etc/profile` (for a system-wide installation): `export PATH=$PATH:/usr/local/go/bin`

    ```
    nano ~/.profile
    ```

    Ctrl + X, then y, then Enter

    ```
    source ~/.profile
    ```

4. Verify that you've installed Go by opening a command prompt and typing the following command:

    ```
    go version
    ```

    go version go1.23.0 linux/amd64

5. Whole process:

    ```
    linux:~/Desktop/Magazyn/projects/golang$ sudo rm -rf /usr/local/go
    [sudo] password for user:
    linux:~/Desktop/Magazyn/projects/golang$ sudo tar -C /usr/local -xzf /home/$(whoami)/Downloads/go1.23.0.linux-amd64.tar.gz
    linux:~/Desktop/Magazyn/projects/golang$ nano ~/.profile
    linux:~/Desktop/Magazyn/projects/golang$ source ~/.profile
    linux:~/Desktop/Magazyn/projects/golang$ go version
    go version go1.23.0 linux/amd64
    linux:~/Desktop/Magazyn/projects/golang$
    ```

Sometimes it requires to restart IDE or system to find `go/bin`.

## Tutorials

### Go Programming - Golang Course with Bonus Projects

https://www.youtube.com/watch?v=un6ZyFkqFKo&ab_channel=freeCodeCamp.org

GitHub: https://github.com/bootdotdev/fcc-learn-golang-assets

TODO 

### Golang Tutorial for Beginners | Full Go Course

TODO todo course https://www.youtube.com/watch?v=yyUHQIec83I&ab_channel=TechWorldwithNana

Repo: https://gitlab.com/nanuchi/go-full-course-youtube

### The Best Resources to Learn Golang (If I Could Start Over)

TODO watch https://www.youtube.com/watch?v=qT14b1pxtiI&ab_channel=Melkey

### Golang course for beginners

TODO watch https://www.youtube.com/watch?v=h9IAzCxDmtU&ab_channel=Visat

No repo

### Learn Go Programming by Building 11 Projects – Full Course

TODO watch https://www.youtube.com/watch?v=jFfo23yIWac&ab_channel=freeCodeCamp.org

### What is the BEST Way to Study Golang? (and get hired...)

TODO watch https://www.youtube.com/watch?v=DCU9GkggAPg&ab_channel=Melkey

## REST API Client

https://install.advancedrestclient.com/home

REST Client in VS Code
