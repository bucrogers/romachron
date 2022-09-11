<h1 align="center">romachron</h1>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#run-from-docker">Run from Docker</a>
    <li><a href="#run-natively-without-docker">Run natively (without Docker)</a></li>
  </ol>
</details>

## Introduction
Roman digital clock, using Golang and React

## Run from Docker
* Prequisite: Docker installed
* Pulls images from dockerhub by default. To build from local changes, specify --build
* For Windows (not using WSL / Bash), subsitute powershell script run.ps1 for run.sh below
### Fastest launch
```
./run.sh
```
### Build from source, with layer-caching
```
./run.sh --build
```
### Stop
```
./run.sh stop
```
## Run natively (without Docker)
All instructions below assume brew package manager (available for macos and linux (including Windows WSL))

### Prerequisites
* Tested on go 1.19, node 16.17.0. The steps below will install these or newer versions.
* brew package manager assumed (available for MacOS, Linux / WSL)

```
brew install go
brew install nvm
nvm install --lts
```

### First-time clone prerequisites
```
cd frontend
npm install
```

## Startup server and frontend
#### server (leave terminal window open)
```
cd server
go run main.go
```

#### frontend (leave terminal window open)
```
cd frontend
npm start
```



