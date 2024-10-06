# My project

-   [My project](#my-project)
    -   [Create go.mod file](#create-gomod-file)
    -   [Build and run application](#build-and-run-application)
    -   [Add packages](#add-packages)
    -   [Problems](#problems)
        -   [listen tcp :8000: bind: address already in use](#listen-tcp-8000-bind-address-already-in-use)

## Create go.mod file

```bash
go mod init github.com/pawelNu/rssagg
```

## Build and run application

```bash
go mod tidy
go build -o rssagg && ./rssagg
```

Or use script [build_and_run.sh](build_and_run.sh).

```bash
#!/bin/bash
# chmod +x build_and_run.sh

dir=$(dirname "$0")
cd "$dir" || exit

go mod tidy

go build -o rssagg && ./rssagg
```

## Add packages

```bash
go get github.com/joho/godotenv
go get github.com/go-chi/chi
go get github.com/go-chi/cors
go mod vendor # to copy library code to vendor folder
```

## Problems

### listen tcp :8000: bind: address already in use

Look what is running on port 8000:

```bash
lsof -i :8000
```

Results:

```
COMMAND   PID  USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
rssagg  40789 pawel    3u  IPv6  85160      0t0  TCP *:8000 (LISTEN)
```

Close the process:

```bash
kill -9 40789
```
