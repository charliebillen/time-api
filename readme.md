# Time API

A simple service that returns a JSON representation of the hours, minutes and seconds of the current UTC time:
```
% go run cmd/webserver/main.go
% curl localhost:8000/time
{"Hour":16,"Minute":12,"Second":40}
``` 
Just for playing TDDing Go services.