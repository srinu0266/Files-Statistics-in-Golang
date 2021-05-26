# Abstract
This Application is to fetch files from directory given from cli and get file info and post to http endpoint specified in cli.

# Quickstart

```bash
cd stage1
```

```bash
go mod download
```

##  compile and build exe
```bash
go build -o stage1.exe
```
## Run the application
```bash
.\stage1.exe -httpendpoint=http://127.0.0.1:8080/files -goroutines=3
```

## For Cli args help
```bash
.\stage1.exe --help
```




