# golang-zerotier-service-api
A Golang wrapper for the Zerotier Service API 

This package was created for usage in the [GameKube](https://github.com/BumbleB-NL/gamekube) project. 

# Calls

The following API calls are currently implemented:

- Join network
- Leave network
- Get status

# Installation
Install the package with:

`go get github.com/PimSanders/golang-zerotier-service-api` 

# Usage
Create a new client:

```golang
zts := golangzerotierserviceapi.NewClient("https://api.zerotier.com/api/v1", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", false)
```

Join a network:

```golang
resp, err := zts.JoinNetwork("856127940c6a45a4")

if err != nil {
	log.Panic(err)
}

fmt.Println(resp.Result)
```
Leave a network:

```golang
resp, err := zts.LeaveNetwork("856127940c6a45a4")

if err != nil {
	log.Panic(err)
}

fmt.Println(resp.Status)
``` 
