## GO servers 
Testing implementations of Servers
- Websocket
- Http Server

Additional implementations 
- Configuration file reading from different env file based on parsed args. 
- Basic State machine checking incorrect state change requests, invokes process if state change is successful.

#### Basic Websocket implementation
A Basic websocket implementation with a client reading and receiving a response from the server.

Start server
```
go run socket_server.go  
```
Send a Client request 
```
go run socket_client.go  
```

#### Http Server
A http server with a state machine implementation. 
Allowing requests to change state and get the current state.
```
go run start_server.go local
```
`local` arg is for which config file the system should read from.

Get Current State
```
curl -X GET http://localhost:8080/current
```

Change state
```
curl -X GET http://localhost:8080/state -H 'Content-Type: application/json' -d '{"RequestType":"first"}'
```

Available States
```
first
optionalsecond
third
optionalfourth
fifth
```