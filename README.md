# Shortener
Go+redis+kafka
## Install
Install app
`make instal`

Build images
`make build`

Start shared services. Run PostgreSQL 
`make start-all`

Start app golang
`make start`

Stop app
`make stop`

## Add local package

`cd pkg/utils` => `go mod init`

in main `go.mod` file add our module
```
require (
    ...
	utils v1.0.0 // indirect
)

replace utils v1.0.0 => ./pkg/utils
```

run command `go mod tidy` and get package in main project `go get utils@v1.0.0 `

