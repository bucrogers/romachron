# romachron
Roman digital clock, using golang and React

# Run from Docker
## Linux / MacOS
```
./run.sh
```
```
./run.sh stop
```
## Windows
```
./run.ps1
```
```
./run.ps1 stop
```

# Run natively (without Docker)
All instructions below assume brew package manager (available for macos and linux (including Windows WSL))

## Host prerequisites
Tested on go 1.19, node 16.17.0. The steps below will install these or new versions.

`brew install go`

`brew install nvm`

`nvm install --lts`

## First-time clone prerequisites
`cd frontend`

`npm install`

## Startup server and frontend
### server (leave terminal window open)
`cd server`

`go run main.go`

### frontend (leave terminal window open)
`cd ../frontend`

`npm start`



