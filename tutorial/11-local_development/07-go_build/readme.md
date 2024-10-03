# Go Build

`go build` compiles go code into an executable program

## Build an executable

Ensure you are in your hellogo repo, then run:

```bash
go build
```

Build and run binary file

```bash
go build && ./binaryName
```

Binary name is usually module name from `go.mod` file. 
Example `module github.com/pawelNu/hellogo` binary name will be `hellogo`.

For Windows build

```bash
GOOS=windows GOARCH=amd64 go build -o main.exe
```

For complex project

```bash
GOOS=windows GOARCH=amd64 go build -o output.exe
```

Run the new program:

```bash
./hellogo
```
