# Time API

![Go](https://github.com/charliebillen/time-api/workflows/Go/badge.svg)

A simple service that returns a JSON representation of the hours, minutes and seconds of the current UTC time:
```
% docker build -t time .
% docker run --rm -p 8000:8000 time
% curl localhost:8000/time
{"Hour":16,"Minute":12,"Second":40}
``` 
Just for playing TDDing Go services.

## Features
- [x] Simple DI example
- [x] Multi-stage Docker build example
- [x] CI process
